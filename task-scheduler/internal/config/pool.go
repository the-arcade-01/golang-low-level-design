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
	workerQueue chan models.Runnable
	wg          sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	once.Do(func() {
		pool = &WorkerPool{
			workers:     workers,
			workerQueue: make(chan models.Runnable),
		}

		for i := range pool.workers {
			pool.wg.Add(1)
			fmt.Printf("Worker routine initialized, ID: %v\n", i)

			go func(i int) {
				defer pool.wg.Done()

				for run := range pool.workerQueue {
					run.Run()
				}
			}(i)
		}
	})
	return pool
}

func (p *WorkerPool) Add(run models.Runnable) {
	p.workerQueue <- run
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.workerQueue)
}
