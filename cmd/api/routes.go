package main

import (
	"encoding/json"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /transactions", func(w http.ResponseWriter, r *http.Request) {
		var paymentDto dto.PaymentDto
		err := json.NewDecoder(r.Body).Decode(&paymentDto)
		if err != nil {
			responseJson(w, http.StatusBadRequest, err.Error())
			return
		}

		paymentDto, err = app.paymentService.AddPayment(r.Context(), paymentDto)
		if err != nil {
			responseJson(w, http.StatusBadRequest, err.Error())
			return
		}

		responseJson(w, http.StatusCreated, paymentDto)
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var userDto dto.UserDto
		err := json.NewDecoder(r.Body).Decode(&userDto)
		if err != nil {
			responseJson(w, http.StatusBadRequest, err.Error())
			return
		}

		userDto, err = app.userService.CreateUser(r.Context(), userDto)

		if err != nil {
			responseJson(w, http.StatusBadRequest, err.Error())
			return
		}

		responseJson(w, http.StatusCreated, userDto)
	})

	return mux
}
