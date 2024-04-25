package routers

import (
	"database/sql"
	"net/http"
)

func Routes(db *sql.DB) *http.ServeMux {
	transactionRouter := NewTransactionRouter(db)
	userRouter := NewUserRouter(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /transactions", transactionRouter.AddPayment)
	mux.HandleFunc("POST /users", userRouter.CreateUser)

	return mux
}
