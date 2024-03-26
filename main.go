package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/the-arcade-01/go-dynamodb-example/handlers"
)

type Server struct {
	Router *chi.Mux
}

func (server *Server) MountMiddlewares() {
	server.Router.Use(middleware.Logger)
}

func (server *Server) MountHandlers() {
	server.Router.Get("/greet", handlers.Greet)
}

func CreateNewServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	server.MountMiddlewares()
	server.MountHandlers()
	return server
}

func main() {
	server := CreateNewServer()
	log.Println("server running on port:8080")
	http.ListenAndServe(":8080", server.Router)
}
