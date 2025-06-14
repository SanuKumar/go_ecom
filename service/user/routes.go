package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Handle user login logic here

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload

	//check if user already exists
	// if it doesn't exist, create a new user
}
