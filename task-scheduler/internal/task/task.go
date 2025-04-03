package task

import (
	"fmt"
	"time"
)

type TaskType string
type TaskPriority int

const (
	High TaskPriority = iota
	Medium
	Low
)

const (
	OneTime    TaskType = "OneTime"
	FixedDelay TaskType = "FixedDelay"
	FixedRate  TaskType = "FixedRate"
)

type Task struct {
	ID             int
	Name           string
	Creator        string
	MaxRetries     int
	Priority       TaskPriority
	Type           TaskType
	RecurringDelay int
	ExecutionTime  time.Time
	CreatedAt      time.Time
}

func NewTask(id int, name, creator string, retries int, priority TaskPriority, tp TaskType, initialDelay, recurringDelay int) *Task {
	return &Task{
		ID:             id,
		Name:           name,
		Creator:        creator,
		MaxRetries:     retries,
		Priority:       priority,
		Type:           tp,
		RecurringDelay: recurringDelay,
		ExecutionTime:  time.Now().Add(time.Duration(initialDelay) * time.Second),
		CreatedAt:      time.Now(),
	}
}

func (t *Task) Run() {
	if t.Type == FixedRate {
		t.UpdateExecutionTime(time.Now().Add(time.Duration(t.RecurringDelay) * time.Second))
		taskScheduler.Schedule(t)
	}

	// performing execution of the task
	t.execute()

	if t.Type == FixedDelay {
		t.UpdateExecutionTime(time.Now().Add(time.Duration(t.RecurringDelay) * time.Second))
		taskScheduler.Schedule(t)
	}
}

func (t *Task) execute() {
	fmt.Printf("Task Executed:\n")
	fmt.Printf("ID: %d\n", t.ID)
	fmt.Printf("Name: %s\n", t.Name)
	fmt.Printf("Creator: %s\n", t.Creator)
	fmt.Printf("Max Retries: %d\n", t.MaxRetries)
	fmt.Printf("Priority: %s\n", map[TaskPriority]string{High: "High", Medium: "Medium", Low: "Low"}[t.Priority])
	fmt.Printf("Type: %s\n", t.Type)
	fmt.Printf("Recurring Delay: %d seconds\n", t.RecurringDelay)
	fmt.Printf("Execution Time: %s\n", t.ExecutionTime.Format(time.RFC1123))
	fmt.Printf("Created At: %s\n", t.CreatedAt.Format(time.RFC1123))
}

func (t *Task) UpdateID(id int) *Task {
	t.ID = id
	return t
}

func (t *Task) UpdateExecutionTime(newTime time.Time) *Task {
	t.ExecutionTime = newTime
	return t
}
