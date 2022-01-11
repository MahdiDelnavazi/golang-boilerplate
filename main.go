package main

import (
	"fmt"
	"golang-boilerplate/Router"
	"golang-boilerplate/Service"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

func main() {
	log, err := Service.NewLogger("Polaris Storage Service")
	if err != nil {
		fmt.Errorf("error at start %w", err)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("Startup", "Error", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// Set MaxProcss
	if _, maxProcessorError := maxprocs.Set(); maxProcessorError != nil {
		return fmt.Errorf("failed to set maxprocs: %w", maxProcessorError)
	}

	// Read Config from Environment Variables
	cfg := struct {
		Api struct {
			ApiHost         string        `env:"API_HOST" envDefault:"localhost:3000"`
			ReadTimeOut     time.Duration `env:"API_READ_TIMEOUT" envDefault:"5s"`
			WriteTimeOut    time.Duration `env:"API_WRITE_TIMEOUT" envDefault:"5s"`
			ShutdownTimeout time.Duration `env:"API_SHUT_DOWN_TIMEOUT" envDefault:"5s"`
		}
		DB struct {
			User         string `env:"DB_USER" envDefault:"postgres"`
			Password     string `env:"DB_PASSWORD" envDefault:"postgres"`
			Host         string `env:"DB_HOST" envDefault:"localhost"`
			Name         string `env:"DB_NAME" envDefault:"postgres"`
			MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"10"`
			MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"100"`
			DisableTLS   bool   `env:"DB_DISABLE_TLS" envDefault:"false"`
		}
	}{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	// App Starting
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20
	app.Static("/assets/", "./public")
	Router.Routes(app, log)

	// =====================================================
	// Open Database Connection
	db, err := Service.DatabaseOpen(Service.DatabaseConfig{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("opening Database: %w", err)
	}

	errorChannel := make(chan error)
	func() {
		log.Infow("Project Running On PORT")
		errorChannel <- app.Run(cfg.Api.ApiHost)
	}()

	defer func() {
		log.Infow("shutdown", "status", "here", "host", cfg.DB.Host)
		db.Close()
	}()

	return nil
}
