package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/services"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	port int
	dsn  string
}

type Application struct {
	accountService  *services.AccountService
	merchantService *services.MerchantService
	paymentService  *services.PaymentService
	userService     *services.UserService
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.dsn, "db-dsn", "database.sqlite", "Data source name (DSN)")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := NewApplication(db)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func NewApplication(db *sql.DB) *Application {
	return &Application{
		accountService:  services.NewAccountService(db),
		merchantService: services.NewMerchantService(db),
		paymentService:  services.NewPaymentService(db),
		userService:     services.NewUserService(db),
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("database connection pool established")
	return db, nil
}
