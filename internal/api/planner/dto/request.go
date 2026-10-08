package dto

import "time"

type CreatePlannerEntryRequest struct {
	TaskID      string    `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeletePlannerEntryRequest struct {
	EntryID string `json:"entry_id"`
}

type GetEntryByIDRequest struct {
	EntryID string `json:"entry_id"`
}
