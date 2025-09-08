package handlers

import (
	"encoding/json"
	"net/http"
	"simple-todo-app/internal/helpers"
	"simple-todo-app/internal/models"
	"simple-todo-app/internal/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (s *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var data models.UserRegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Request invalaid : "+err.Error(), nil)
		return
	}

	res, err := s.service.Register(data)
	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "User registered", res)
}

func (s *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data models.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Request invalaid :"+err.Error(), nil)
		return
	}

	res, err := s.service.Login(data)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Login success", res)
}

func (s *UserHandler) Info(w http.ResponseWriter, r *http.Request) {
	jwtInfo := r.Context().Value("userInfo").(*helpers.Claims)

	if jwtInfo == nil {
		helpers.ApiResponse(w, http.StatusBadRequest, "Profile invalid", nil)
		return
	}

	user, err := s.service.FindUserByEmail(jwtInfo.Email)

	if err != nil {
		helpers.ApiResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.ApiResponse(w, http.StatusOK, "Profile loaded", user)
}
