package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/routers"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	port int
	dsn  string
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

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: routers.Routes(db),
	}

	log.Printf("Starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
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
