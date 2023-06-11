package main

import (
	"context"
	"errors"
	"firstpass/internal/server/config"
	"firstpass/internal/server/handler"
	"firstpass/internal/server/httpserver"
	"firstpass/internal/server/repository"
	"firstpass/internal/server/service"
	"flag"
	"github.com/caarlos0/env/v6"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	// Create and read config
	cfg := config.Config{}
	// Parse config from flag
	flag.StringVar(&cfg.RunAddress, "a", ":8080", "server address and port")
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	flag.StringVar(&cfg.DatabaseURI, "d", "postgres://postgres:pass@localhost/postgres?sslmode=disable", "database Uniform Resource Identifier")
	//flag.StringVar(&cfg.DatabaseURI, "d", "postgres://postgres:pass@postgres/postgres?sslmode=disable", "database Uniform Resource Identifier")
	flag.Parse()
	// Parse config from env
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}
	log.Printf("using config: %#v", cfg)
	// Start server
	err = run(cfg)
	if err != nil {
		log.Fatalf("gophermart server failed: %v", err)
	}
}

func run(cfg config.Config) (err error) {

	defer func() {
		// handle panic
		if x := recover(); x != nil {
			log.Printf("runtime panic: %v\n", x)
			err = errors.New("runtime panic: " + x.(string))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		log.Printf("interrupt signal")
		cancel()
	}()

	// Connect to DB
	db, err := sqlx.Connect("pgx", cfg.DatabaseURI)
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}

	// Create table if not exists
	file, err := os.ReadFile(filepath.Join("migrations", "01_init_up.sql"))
	if err != nil {
		log.Fatalf("unable to create tables in db: %v", err)
	}
	schema := string(file)
	db.MustExec(schema)

	log.Println("connected to DB")

	// Create repositories
	userRepo := repository.NewUserRepository(db)
	noteRepo := repository.NewNoteRepository(db)
	cardRepo := repository.NewCardRepository(db)
	binRepo := repository.NewBinRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	// Create services
	userService := service.NewUserService(userRepo)
	noteService := service.NewNoteService(noteRepo)
	cardService := service.NewCardService(cardRepo)
	binService := service.NewBinService(binRepo)
	accountService := service.NewAccountService(accountRepo)
	// Create handler
	h := handler.New(userService, noteService, cardService, binService, accountService)

	// Create server
	srv := httpserver.New()

	// Start http server
	err = srv.Run(ctx, cfg.RunAddress, h.InitRoutes())
	if err != nil && err == http.ErrServerClosed {
		log.Printf("server closed: %v", err)
		return nil
	}
	return err
}
