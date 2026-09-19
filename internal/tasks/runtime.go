package tasks

import (
	"database/sql"
	"sync"
	"time"

	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/tasks"
)

type TaskRegistry struct {
	TaskRepository *tasks.TaskRepository
	Tasks          map[string]tasks.Task
	mu             sync.Mutex
}

func NewTaskRegistry(db *sql.DB) *TaskRegistry {
	return &TaskRegistry{
		TaskRepository: tasks.NewTaskRepository(&database.Database{DB: db}),
		mu:             sync.Mutex{},
		Tasks:          make(map[string]tasks.Task),
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
