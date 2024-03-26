package entity

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        string `dynamodbav:"id"`
	Name      string `dynamodbav:"name"`
	Completed bool   `dynamodbav:"completed"`
	Timestamp int64  `dynamodbav:"timestamp"`
}

type TodoRequestBody struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func CreateTodo(name string, completed bool) *Todo {
	return &Todo{
		Id:        uuid.New().String(),
		Name:      name,
		Completed: completed,
		Timestamp: time.Now().Unix(),
	}
}
