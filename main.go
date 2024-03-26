package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/the-arcade-01/go-dynamodb-example/handlers"
)

type Server struct {
	Router       *chi.Mux
	DynamoClient *dynamodb.Client
}

func (server *Server) MountMiddlewares() {
	server.Router.Use(middleware.Logger)
}

func (server *Server) MountHandlers() {
	TODOS_TABLE := os.Getenv("DYNAMO_TODOS_TABLE")
	handlers := handlers.CreateHandlers(server.DynamoClient, TODOS_TABLE)
	server.Router.Get("/greet", handlers.Greet)
	server.Router.Get("/tables", handlers.GetTables)
	server.Router.Get("/todos", handlers.GetTodos)
	server.Router.Post("/todos", handlers.CreateItem)
	server.Router.Get("/todos/{id}", handlers.GetTodoById)
	server.Router.Delete("/todos/{id}", handlers.DeleteTodoById)
}

func CreateNewServer(client *dynamodb.Client) *Server {
	server := &Server{
		Router:       chi.NewRouter(),
		DynamoClient: client,
	}
	server.MountMiddlewares()
	server.MountHandlers()
	return server
}

func Config() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		log.Printf("[config] unable to load SDK config, error: %v, \n", err)
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	log.Println("[config] Dynamodb client loaded")
	return client, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("[init] error loading .env variables")
	}
}

func main() {
	client, err := Config()
	if err != nil {
		panic(err)
	}
	server := CreateNewServer(client)
	log.Println("[main] server running on port:8080")
	http.ListenAndServe(":8080", server.Router)
}
