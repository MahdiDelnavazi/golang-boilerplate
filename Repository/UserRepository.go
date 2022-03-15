package Repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	"golang-boilerplate/Entity"
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
	queryError := userRepository.database.Get(&user, `SELECT * FROM newuser($1 , 0)`, creatUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, nil
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
