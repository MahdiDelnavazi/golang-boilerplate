package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"go.uber.org/automaxprocs/maxprocs"
	"golang-boilerplate/Config"
	"golang-boilerplate/Middleware/token"
	"golang-boilerplate/Router"
	"golang-boilerplate/Service"
	"log"
)

func main() {
	logger, loggerError := Service.NewLogger("Polaris Storage Service")
	if loggerError != nil {
		fmt.Errorf("error at start %w", loggerError)
	}
	defer logger.Sync()

	if _, maxProcessorError := maxprocs.Set(); maxProcessorError != nil {
		fmt.Errorf("failed to set maxprocs: %w", maxProcessorError)
	}

	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}
	fmt.Printf("%+v\n", config)

	// =====================================================
	// Open Database Connection
	database, err := Service.DatabaseOpen(Service.DatabaseConfig{
		User:         config.DB.User,
		Password:     config.DB.Password,
		Host:         config.DB.Host,
		Name:         config.DB.Name,
		MaxIdleConns: config.DB.MaxIdleConns,
		MaxOpenConns: config.DB.MaxOpenConns,
		DisableTLS:   config.DB.DisableTLS,
	})
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	logger.Infow("Project database, ", "database", database)

	defer func() {
		logger.Infow("shutdown", "status", "here", "host", config.DB.Host)
		database.Close()
	}()

	tokenMaker, err := token.NewPasetoMaker(config.Token.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token: %W", err)
	}

	// App Starting
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20
	app.Static("/assets/", "./public")
	Router.Routes(app, logger, database, tokenMaker)

	errorChannel := make(chan error)
	func() {
		logger.Infow("Project Running On PORT", config.Api.ApiHost)
		errorChannel <- app.Run(config.Api.ApiHost)
	}()
}
