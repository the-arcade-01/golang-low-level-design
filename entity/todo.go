package entity

import "github.com/google/uuid"

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Timestamp int64     `json:"timestamp"`
}

type TodoRequestBody struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
