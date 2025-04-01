package main

import (
	"task-scheduler/internal/models"
	"task-scheduler/internal/scheduler"
	"time"
)

func main() {
	scheduler := scheduler.NewTaskScheduler()

	fixedTime := time.Now().Add(time.Duration(4) * time.Second)

	task1 := models.NewTask(1, "Testing task1", "Meowth", models.OneTime, 3, fixedTime, models.Medium, 300)
	task2 := models.NewTask(2, "Testing task2", "Ben10", models.FixedDelay, 3, fixedTime, models.Low, 300)
	task3 := models.NewTask(3, "Testing task3", "Gwen", models.FixedRate, 3, fixedTime, models.High, 300)

	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)

	scheduler.Pop()
}
