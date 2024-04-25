package routers

import (
	"database/sql"
	"encoding/json"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/services"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"net/http"
)

type UserRouter struct {
	userService *services.UserService
}

func NewUserRouter(db *sql.DB) *UserRouter {
	return &UserRouter{
		userService: services.NewUserService(db),
	}
}

func (u *UserRouter) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		responseJson(w, http.StatusBadRequest, err.Error())
		return
	}

	userDto, err = u.userService.CreateUser(r.Context(), userDto)

	if err != nil {
		responseJson(w, http.StatusBadRequest, err.Error())
		return
	}

	responseJson(w, http.StatusCreated, userDto)
}
