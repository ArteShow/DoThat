package dto

import "time"

type CreateTaskResponse struct {
	TaskID string `json:"task_id"`
}

type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Duration    int       `json:"duration"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

type GetTaskByIDResponse struct {
	Task TaskResponse `json:"task"`
}

type GetErrorResponse struct {
	Error string `json:"error"`
}
