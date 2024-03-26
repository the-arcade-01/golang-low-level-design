package entity

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Todo struct {
	Id        string `dynamodbav:"id" json:"id"`
	Name      string `dynamodbav:"name" json:"name"`
	Completed bool   `dynamodbav:"completed" json:"completed"`
	Timestamp int64  `dynamodbav:"timestamp" json:"timestamp"`
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

func GetKey(id string) (map[string]types.AttributeValue, error) {
	itemId, err := attributevalue.Marshal(id)
	return map[string]types.AttributeValue{"id": itemId}, err
}
