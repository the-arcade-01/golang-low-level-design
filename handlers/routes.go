package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/the-arcade-01/go-dynamodb-example/entity"
)

type Handler struct {
	client *dynamodb.Client
	table  string
}

func CreateHandlers(client *dynamodb.Client, table string) *Handler {
	return &Handler{
		client: client,
		table:  table,
	}
}

func (handler Handler) Greet(w http.ResponseWriter, r *http.Request) {
	ResponseWithJSON(w, http.StatusOK, "Hello, World!!")
}

func (handler Handler) GetTables(w http.ResponseWriter, r *http.Request) {
	tables, err := handler.client.ListTables(context.TODO(), nil)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't list tables, err: %v", err.Error()))
		return
	}
	ResponseWithJSON(w, http.StatusOK, tables)
}

func (handler Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(handler.table),
	}
	todos, err := handler.client.Scan(context.TODO(), scanInput)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't list Todos, err: %v", err.Error()))
		return
	}
	ResponseWithJSON(w, http.StatusOK, todos)
}

func (handler Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var todoReqBody entity.TodoRequestBody
	if err := json.NewDecoder(r.Body).Decode(&todoReqBody); err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	todo := &entity.Todo{
		Id:        uuid.New(),
		Name:      todoReqBody.Name,
		Completed: todoReqBody.Completed,
	}
	// if err != nil {
	// 	ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create Todos, err: %v", err.Error()))
	// 	return
	// }
	putInput := &dynamodb.PutItemInput{
		TableName: &handler.table,
		Item:      nil,
	}
	handler.client.PutItem(context.TODO(), putInput)
	ResponseWithJSON(w, http.StatusAccepted, todo)
}

func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
