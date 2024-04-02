package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (server *APIServer) Start() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	taskService := NewTasksService(server.store)
	taskService.RegisterRoutes(subRouter)

	userService := NewUserService(server.store)
	userService.RegisterRoutes(subRouter)

	log.Println("Starting API server on", server.addr)

	log.Fatal(http.ListenAndServe(server.addr, subRouter))
}
