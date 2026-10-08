package dto

import "time"

type CreatePackRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeletePackRequest struct {
	PackID string `json:"pack_id"`
}

type GetPackByIDRequest struct {
	PackID string `json:"pack_id"`
}
