package Repository

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang-boilerplate/Config"
	"golang-boilerplate/DTO/Request/User"
	"golang-boilerplate/Service"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	userRepository := NewUserRepository(Logger, TestDB)
	user := User.CreateUserRequest{UserName: "mimdl"}
	createUser, err := userRepository.CreateUser(user)
	require.NoError(t, err)
	require.NotNil(t, createUser)
	require.Equal(t, createUser.UserName, user.UserName)

}

var TestDB *sqlx.DB
var Logger *zap.SugaredLogger

func init() {
	config := Config.EnvironmentConfig{}
	Logger, _ = Service.NewLogger("polaris")
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	// =====================================================
	// Open Database Connection
	db, err := Service.DatabaseOpen(Service.DatabaseConfig{
		User:         config.DB.User,
		Password:     config.DB.Password,
		Host:         config.DB.Host,
		Name:         config.DB.Name,
		MaxIdleConns: config.DB.MaxIdleConns,
		MaxOpenConns: config.DB.MaxOpenConns,
		DisableTLS:   config.DB.DisableTLS,
	})
	if err != nil {
		fmt.Errorf("cannot connect to db: %s", err)
	}
	TestDB = db
}
