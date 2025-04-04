package task

import (
	"container/heap"
	"fmt"
	"sync"
	"task-scheduler/internal/config"
	"time"
)

type Scheduler interface {
	Schedule(task *Task)
}

var (
	schedulerOnce sync.Once
	taskScheduler *TaskScheduler
)

type TaskScheduler struct {
	mtx       *sync.Mutex
	cond      *sync.Cond
	pool      *config.WorkerPool
	taskQueue *TaskHeap
}

func NewTaskScheduler(pool *config.WorkerPool) *TaskScheduler {
	schedulerOnce.Do(func() {
		taskScheduler = &TaskScheduler{
			mtx:       &sync.Mutex{},
			pool:      pool,
			taskQueue: NewTaskHeap(),
		}
		taskScheduler.cond = sync.NewCond(taskScheduler.mtx)

		go taskScheduler.Execute()
	})
	return taskScheduler
}

func (t *TaskScheduler) Schedule(task *Task) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	fmt.Printf("Task: %v Scheduled\n", task.Name)
	heap.Push(t.taskQueue, task)
	t.cond.Signal()
}

func (t *TaskScheduler) Stop(task *Task) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	for i, item := range *t.taskQueue {
		if item == task {
			fmt.Printf("Task: %v Stopped\n", task.Name)
			heap.Remove(t.taskQueue, i)
			break
		}
	}
}

func (t *TaskScheduler) Execute() {
	for {
		t.mtx.Lock()
		if t.taskQueue.Len() == 0 {
			t.cond.Wait()
		}

		task := heap.Pop(t.taskQueue).(*Task)
		if task.ExecutionTime.After(time.Now()) {
			heap.Push(t.taskQueue, task)

			t.mtx.Unlock()
			time.Sleep(time.Until(task.ExecutionTime))
			continue
		}

		t.mtx.Unlock()
		t.pool.Add(task)
	}
}
