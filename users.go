package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService struct {
	store Store
}

func NewUserService(store Store) *UserService {
	return &UserService{store: store}
}

func (us *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/register", us.handleUserRegister).Methods(http.MethodPost)
}

func (us *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error reading request body"})
		return
	}
	defer r.Body.Close()

	var payload *User
	if err := json.Unmarshal(body, &payload); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error hashing password"})
		return
	}
	payload.Password = hashedPassword

	createdUser, err := us.store.CreateUser(payload)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := CreateAndSetAuthCookie(createdUser.Id, w)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating token"})
		return
	}

	WriteJson(w, http.StatusCreated, token)

}

func validateUserPayload(user *User) error {
	if user.Email == "" {
		return errors.New("user email is required")
	}
	if user.Password == "" {
		return errors.New("user password is required")
	}
	return nil
}
