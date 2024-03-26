package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
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
	ResponseWithJSON(w, http.StatusOK, tables.TableNames)
}

func (handler Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(handler.table),
	}

	results, err := handler.client.Scan(context.TODO(), scanInput)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't list Todos, err: %v", err.Error()))
		return
	}

	var todos []*entity.Todo
	for _, item := range results.Items {
		todo := new(entity.Todo)
		err := attributevalue.UnmarshalMap(item, todo)
		if err != nil {
			ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't unmarshal the Todos, err: %v", err.Error()))
			return
		}
		todos = append(todos, todo)
	}

	ResponseWithJSON(w, http.StatusOK, todos)
}

func (handler Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var todoReqBody entity.TodoRequestBody
	if err := json.NewDecoder(r.Body).Decode(&todoReqBody); err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	todo := entity.CreateTodo(todoReqBody.Name, todoReqBody.Completed)

	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create Todos, err: %v", err.Error()))
		return
	}

	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(handler.table),
		Item:      item,
	}

	_, err = handler.client.PutItem(context.TODO(), putInput)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create Todos, err: %v", err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, todo)
}

func (handler Handler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	key, err := entity.GetKey(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(handler.table),
		Key:       key,
	}
	results, err := handler.client.GetItem(context.TODO(), getInput)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get Todo by Id: %v, err: %v", id, err.Error()))
		return
	}

	todo := new(entity.Todo)
	err = attributevalue.UnmarshalMap(results.Item, todo)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't unmarshal Todo by Id: %v, err: %v", id, err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusOK, todo)
}

func (handler Handler) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	key, err := entity.GetKey(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(handler.table),
		Key:       key,
	}

	results, err := handler.client.DeleteItem(context.TODO(), deleteInput)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete Todo by Id: %v, err: %v", id, err.Error()))
		return
	}

	ResponseWithJSON(w, http.StatusOK, results)
}

func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
