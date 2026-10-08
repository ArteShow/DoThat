package dto

import "time"

type CreatePlannerEntryResponse struct {
	EntryID string `json:"entry_id"`
}

type EntryResponse struct {
	ID          string    `json:"entry_id"`
	TaskID      string    `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllEntriesResponse struct {
	Entries []EntryResponse `json:"entries"`
}

type GetEntryByIDResponse struct {
	Entry EntryResponse `json:"entry"`
}
