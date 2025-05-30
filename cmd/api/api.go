package api

import (
	"database/sql"
	"github/SanuKumar/go_ecom/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter() // Gorilla Mux router
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler() // Assuming you have a user package with a NewHandler function
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
