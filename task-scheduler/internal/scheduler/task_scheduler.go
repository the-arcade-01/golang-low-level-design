package scheduler

import (
	"container/heap"
	"fmt"
	"sync"
	"task-scheduler/internal/models"
	"task-scheduler/internal/queue"
)

type TaskScheduler struct {
	mtx           sync.Mutex
	priorityQueue *queue.Heap
}

func NewTaskScheduler() *TaskScheduler {
	pq := queue.NewHeap()
	heap.Init(pq)

	return &TaskScheduler{
		priorityQueue: pq,
	}
}

func (s *TaskScheduler) AddTask(task *models.Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	heap.Push(s.priorityQueue, task)
}

func (s *TaskScheduler) DisplayTasks() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	fmt.Println("All tasks in priority queue")
	for _, task := range *s.priorityQueue {
		fmt.Printf("%v %v %v %v %v\n", task.ID, task.Name, task.StartTime, task.Priority, task.Type)
	}
}

func (s *TaskScheduler) RemoveTask(task *models.Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

}

// just of debugging the ordering
func (s *TaskScheduler) Pop() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for s.priorityQueue.Len() > 0 {
		task := heap.Pop(s.priorityQueue).(*models.Task)
		fmt.Printf("%v %v %v %v %v\n", task.ID, task.Name, task.StartTime, task.Priority, task.Type)
	}
}
