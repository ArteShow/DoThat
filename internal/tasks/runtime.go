package tasks

import (
	"sync"
	"time"

	"github.com/ArteShow/DoThat/internal/database/repositories/tasks"
)

type TaskRegistry struct {
	TaskRepository *tasks.TaskRepository
	Tasks          map[string]tasks.Task

	latestErr error
	mu        sync.Mutex
}

func NewTaskRegistry(repo *tasks.TaskRepository, registry map[string]tasks.Task) *TaskRegistry {
	return &TaskRegistry{
		TaskRepository: repo,
		mu:             sync.Mutex{},
		Tasks:          registry,
	}
}

func (r *TaskRegistry) CreateTask(task tasks.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Tasks[task.ID] = task
}

func (r *TaskRegistry) DeleteTask(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Tasks, id)
}

func (r *TaskRegistry) GetAll() []tasks.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	var tasks []tasks.Task
	for _, task := range r.Tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (r *TaskRegistry) GetByID(id string) tasks.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.Tasks[id]
}

func (r *TaskRegistry) UpdateStatus(id, status string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := r.Tasks[id]
	task.Status = status

	r.Tasks[id] = task
}

func (r *TaskRegistry) UpdateDeadline(id string, newDeadline time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := r.Tasks[id]
	task.Deadline = newDeadline

	r.Tasks[id] = task
}

func (r *TaskRegistry) GetError() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.latestErr
	r.latestErr = nil

	return err
}
