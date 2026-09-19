package tasks

import (
	"database/sql"
	"time"

	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/tasks"
	"github.com/google/uuid"
)

type TaskManager struct {
	registry *TaskRegistry
}

func NewTaskManager(db *sql.DB) (*TaskManager, error) {
	repo := tasks.NewTaskRepository(&database.Database{DB: db})

	registry, err := LoadRegistry(repo)
	if err != nil {
		return &TaskManager{}, err
	}

	return &TaskManager{
		registry: NewTaskRegistry(repo, registry),
	}, nil
}

func LoadRegistry(repo *tasks.TaskRepository) (map[string]tasks.Task, error) {
	registry := map[string]tasks.Task{}

	dbRegistry, err := repo.GetAll()
	if err != nil {
		return make(map[string]tasks.Task), err
	}

	for _, task := range dbRegistry {
		registry[task.ID] = task
	}

	return registry, nil
}

func (m *TaskManager) GetError() error {
	return m.registry.GetError()
}

func (m *TaskManager) CreateTask(task tasks.Task) string {
	id := uuid.NewString()
	task.ID = id

	m.registry.CreateTask(task)
	return id
}

func (m *TaskManager) DeleteTask(id string) {
	m.registry.DeleteTask(id)
}

func (m *TaskManager) UpdateStatus(id, newStatus string) {
	m.registry.UpdateStatus(id, newStatus)
}

func (m *TaskManager) UpdateDeadline(id string, newDeadline time.Time) {
	m.registry.UpdateDeadline(id, newDeadline)
}

func (m *TaskManager) GetAll() []tasks.Task {
	return m.registry.GetAll()
}

func (m *TaskManager) GetByID(id string) tasks.Task {
	return m.registry.GetByID(id)
}
