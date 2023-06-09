package cmd

import (
	"sync"

	"projects/internal/database"
	"log"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	BaseURL  string
	HttpPort int
	DB       struct {
		DSN         string
		AutoMigrate bool
	}
}

type Application struct {
	Config    Config
	DB        *database.DB
	Logger    *log.Logger
	wg        sync.WaitGroup
	Validator *validator.Validate
}