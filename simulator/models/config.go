package models

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type EnvConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBSSLMode  string `env:"DB_SSLMODE"`
	HTTPHost   string `env:"HTTP_Host"`
	Port       string `env:"PORT"`
}

var EnvironmentConfig EnvConfig
var DBConnectionString string

func (cfg DBConfig) ConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
}

func (envcfg *EnvConfig) LoadConfig() error {

	// Load configuration from .env or environment variables
	_ = godotenv.Load()

	// Convert dbPort from string to int
	portInt, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	// Optionally override from environment variables
	envcfg = &EnvConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     portInt,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		HTTPHost:   os.Getenv("HTTP_HOST"),
		Port:       os.Getenv("PORT"),
	}

	EnvironmentConfig = *envcfg

	DBConfig := DBConfig{
		Host:     envcfg.DBHost,
		Port:     envcfg.DBPort,
		User:     envcfg.DBUser,
		Password: envcfg.DBPassword,
		Database: envcfg.DBName,
	}

	DBConnectionString = DBConfig.ConnString()

	return nil
}
