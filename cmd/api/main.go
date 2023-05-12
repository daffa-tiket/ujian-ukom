package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"

	"projects/internal/database"
	"projects/internal/env"
	"projects/internal/version"

	"github.com/go-playground/validator/v10"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatalf("%s\n%s", err, trace)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
}

type application struct {
	config config
	db     *database.DB
	logger *log.Logger
	wg     sync.WaitGroup
	validator *validator.Validate
}

func run(logger *log.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = env.GetString("DB_DSN", "daffashafwan:@localhost:5432/appdb?sslmode=disable")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	validate := validator.New()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
		validator: validate,
	}

	return app.serveHTTP()
}
