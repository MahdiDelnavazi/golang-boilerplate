package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"golang-boilerplate/Config"
	"golang-boilerplate/Service"
	"os"
	"testing"
)

var TestQuery *sqlx.DB

func TestMain(m *testing.M) {
	config := Config.EnvironmentConfig{}
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
	TestQuery = db

	os.Exit(m.Run())

}
