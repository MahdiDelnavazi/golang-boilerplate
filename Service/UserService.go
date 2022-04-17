package Service

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	User2 "golang-boilerplate/DTO/Response/User"
	"golang-boilerplate/Middleware/token"
	"golang-boilerplate/Repository"
	"log"
	"time"
)

type UserService struct {
	userRepository *Repository.UserRepository
	logger         *zap.SugaredLogger
	token          token.Maker
	redis          *redis.Client
}

func NewUserService(logger *zap.SugaredLogger, userRepository *Repository.UserRepository, token token.Maker, redis *redis.Client) *UserService {
	return &UserService{logger: logger, userRepository: userRepository, token: token, redis: redis}
}

func (userService UserService) Create(createUserRequest User.CreateUserRequest) (User2.CreateUserResponse, error) {

	// validate username len and not empty
	validationError := ValidationCheck(createUserRequest)
	log.Println(validationError)
	if validationError != nil {
		return User2.CreateUserResponse{}, validationError
	}

	// check if user exist return error .
	checkUserName, _ := userService.userRepository.CheckUserName(createUserRequest)
	if checkUserName.UserName != "" {
		return User2.CreateUserResponse{UserName: checkUserName.UserName}, errors.New("user exist")
	}

	userResponse, userRepositoryError := userService.userRepository.CreateUser(createUserRequest)
	if userRepositoryError != nil {
		return User2.CreateUserResponse{}, validationError
	}
	// we need a transformer
	return User2.CreateUserResponse{UserName: userResponse.UserName}, nil
}

func (userService UserService) LoginUser(loginUserRequest User.LoginUserRequest) (User2.LoginUserResponse, error) {
	// validate username len and not empty
	validationError := ValidationCheck(loginUserRequest)

	if validationError != nil {
		return User2.LoginUserResponse{}, validationError
	}

	user, getUserError := userService.userRepository.LoginUser(loginUserRequest)
	if getUserError != nil {
		return User2.LoginUserResponse{}, getUserError
	}

	//create new token for login
	accessToken, err := userService.token.CreateToken(loginUserRequest.UserName, time.Hour)

	if err != nil {
		return User2.LoginUserResponse{}, err
	}

	refreshToken, errRefreshToken := userService.token.CreateToken(loginUserRequest.UserName, time.Hour*120)
	if errRefreshToken != nil {
		return User2.LoginUserResponse{}, err
	}

	// we need a transformer
	return User2.LoginUserResponse{UserName: user.UserName, AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (userService UserService) LogoutUser(request User.LogoutRequest) (response string, err error) {
	payload, _ := userService.token.VerifyToken(request.Token)
	err = userService.redis.Set(payload.Username, request.Token, 0).Err()
	if err != nil {
		return "logout failed", err
	}

	return "logout successfully", err
}

func (userService UserService) GetUser(getUserRequest User.GetUserRequest) (User2.GetUserResponse, error) {
	// validate username len and not empty
	validationError := ValidationCheck(getUserRequest)

	if validationError != nil {
		return User2.GetUserResponse{}, validationError
	}

	user, getUserError := userService.userRepository.GetUser(getUserRequest)
	if getUserError != nil {
		return User2.GetUserResponse{}, getUserError
	}
	// we need a transformer
	return User2.GetUserResponse{UserId: user.Id, UserName: user.UserName}, nil
}
