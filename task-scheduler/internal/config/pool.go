package config

import (
	"fmt"
	"sync"
	"task-scheduler/internal/models"
)

var (
	once sync.Once
	pool *WorkerPool
)

type WorkerPool struct {
	workers     int
	wg          sync.WaitGroup
	workerQueue chan *models.Task
}

func NewWorkerPool(workers int) *WorkerPool {
	once.Do(func() {
		pool := &WorkerPool{
			workers:     workers,
			workerQueue: make(chan *models.Task),
		}

		for i := range workers {
			pool.wg.Add(1)

			go func(i int) {
				defer pool.wg.Done()
				fmt.Printf("Worker pool %v initialized\n", i)

				for task := range pool.workerQueue {
					fmt.Printf("%v %v %v %v", task.ID, task.Name, task.Type, task.StartTime)
				}
			}(i)
		}
	})
	return pool
}

func (p *WorkerPool) AddJob(task *models.Task) {
	p.workerQueue <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.workerQueue)
}
