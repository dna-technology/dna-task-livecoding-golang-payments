package main

import (
	"database/sql"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/routers"
	"net/http"
)

func routes(db *sql.DB) *http.ServeMux {
	transactionRouter := routers.NewTransactionRouter(db)
	userRouter := routers.NewUserRouter(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /transactions", transactionRouter.AddPayment)
	mux.HandleFunc("POST /users", userRouter.CreateUser)

	return mux
}
