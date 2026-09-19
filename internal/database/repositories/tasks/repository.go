package tasks

import (
	"time"

	"github.com/ArteShow/DoThat/internal/database"
)

type TaskRepository struct {
	DB *database.Database
}

func NewTaskRepository(db *database.Database) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

type Task struct {
	ID          string    `json:"task_it"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Duration    int       `json:"duration"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
