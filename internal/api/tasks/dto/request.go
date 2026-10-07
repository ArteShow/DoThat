package dto

import "time"

type CreateTaskRequest struct {
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

type DeleteTaskRequest struct {
	TaskID string `json:"task_id"`
}

type GetTaskByIDRequest struct {
	TaskID string `json:"task_id"`
}

type UpdateTaskStatusRequest struct {
	TaskID    string `json:"task_id"`
	NewStatus string `json:"new_status"`
}

type UpdateTaskDeadlineRequest struct {
	TaskID      string    `json:"task_id"`
	NewDeadline time.Time `json:"new_deadline"`
}
