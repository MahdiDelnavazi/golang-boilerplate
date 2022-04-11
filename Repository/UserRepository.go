package Repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	"golang-boilerplate/Entity"
	"golang-boilerplate/Helper"
)

type UserRepository struct {
	logger   *zap.SugaredLogger
	database *sqlx.DB
}

func NewUserRepository(logger *zap.SugaredLogger, database *sqlx.DB) *UserRepository {
	return &UserRepository{
		logger:   logger,
		database: database,
	}
}

// CreateUser exec query for create new user in database
func (userRepository *UserRepository) CreateUser(creatUserRequest User.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}
	password, err := Helper.HashPassword(creatUserRequest.Password)
	if err != nil {
		return Entity.User{}, nil
	}
	queryError := userRepository.database.Get(&user, `SELECT * FROM newuser($1 , $2, 0)`, creatUserRequest.UserName, password)
	if queryError != nil {
		return Entity.User{}, nil
	}
	return user, queryError
}

// LoginUser for login users
func (userRepository *UserRepository) LoginUser(loginUserRequest User.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}
	queryError := userRepository.database.Get(&user, `SELECT * FROM loginuser($1)`, loginUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	if !Helper.CheckPasswordHash(loginUserRequest.Password, user.Password) {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	return user, queryError
}

// CheckUserName check username exist or not
func (userRepository *UserRepository) CheckUserName(creatUserRequest User.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}
	queryError := userRepository.database.Get(&user, `SELECT * FROM checkuserexist($1)`, creatUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, nil
	}
	return user, queryError
}

// to tikcket servesam ino call mikonam az user service
// CheckUserName check username exist or not
func (userRepository *UserRepository) GetUser(creatUserRequest User.GetUserRequest) (Entity.User, error) {
	user := Entity.User{}
	queryError := userRepository.database.Get(&user, `SELECT * FROM getuser($1)`, creatUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, queryError
	}
	return user, queryError
}
