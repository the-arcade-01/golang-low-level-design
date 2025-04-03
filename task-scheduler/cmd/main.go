package main

import (
	"task-scheduler/internal/config"
	"task-scheduler/internal/task"
	"time"
)

func main() {
	pool := config.NewWorkerPool(5)

	taskScheduler := task.NewTaskScheduler(pool)
	taskService := task.NewTaskService(taskScheduler)

	task1 := taskService.CreateTask("Pikachu Capture", "Team Rocket", 3, task.High, task.FixedRate, 60, 180)
	task2 := taskService.CreateTask("Defeat Vilgax", "Ben 10", 3, task.Low, task.FixedDelay, 30, 120)
	task3 := taskService.CreateTask("Eat Ramen", "Naruto", 3, task.Low, task.OneTime, 10, 0)

	taskScheduler.Schedule(task1)
	taskScheduler.Schedule(task2)
	taskScheduler.Schedule(task3)

	go func() {
		time.Sleep(time.Duration(40) * time.Second)
		taskService.StopTask(task2.ID)
	}()

	pool.Wait()
}
