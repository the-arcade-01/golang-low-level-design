package task

import "sync"

type TaskService struct {
	mtx           sync.Mutex
	tasksMap      map[int]*Task
	taskScheduler *TaskScheduler
}

func NewTaskService(taskScheduler *TaskScheduler) *TaskService {
	return &TaskService{
		tasksMap:      make(map[int]*Task),
		taskScheduler: taskScheduler,
	}
}

func (t *TaskService) CreateTask(name, creator string, retries int, priority TaskPriority, tp TaskType, initialDelay, recurringDelay int) *Task {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	task := NewTask(len(t.tasksMap)+1, name, creator, retries, priority, tp, initialDelay, recurringDelay)
	t.tasksMap[task.ID] = task

	return task
}

func (t *TaskService) StopTask(taskID int) bool {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	task, ok := t.tasksMap[taskID]
	if !ok {
		return true
	}
	t.taskScheduler.Stop(task)
	return true
}
