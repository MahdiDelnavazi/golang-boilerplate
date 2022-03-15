package Service

import (
	"errors"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	User2 "golang-boilerplate/DTO/Response/User"
	"golang-boilerplate/Repository"
	"log"
)

type UserService struct {
	userRepository *Repository.UserRepository
	logger         *zap.SugaredLogger
}

func NewUserService(logger *zap.SugaredLogger, userRepository *Repository.UserRepository) *UserService {
	return &UserService{logger: logger, userRepository: userRepository}
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
