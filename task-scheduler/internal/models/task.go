package models

import (
	"math/rand"
	"time"
)

type TaskType string
type TaskPriority int

const (
	OneTime    TaskType = "OneTime"
	FixedRate  TaskType = "FixedRate"
	FixedDelay TaskType = "FixedDelay"
)

const (
	High TaskPriority = iota
	Medium
	Low
)

type Task struct {
	ID        int
	Name      string
	Creator   string
	CreatedAt time.Time
	Type      TaskType
	Retries   int
	StartTime time.Time
	Priority  TaskPriority
	Delay     int // is in seconds
}

func NewTask(id int, name, creator string, tp TaskType, retries int, startTime time.Time, priority TaskPriority, delay int) *Task {
	return &Task{
		ID:        id,
		Name:      name,
		Creator:   creator,
		CreatedAt: time.Now(),
		Type:      tp,
		Retries:   retries,
		StartTime: startTime,
		Priority:  priority,
		Delay:     delay,
	}
}

func (t *Task) Run() bool {
	return rand.Intn(2) == 0
}
