package dto

import "time"

type CreatePackResponse struct {
	PackID string `json:"pack_id"`
}

type PackResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllPacksResponse struct {
	Packs []PackResponse `json:"packs"`
}

type GetPackByIDResponse struct {
	Pack PackResponse `json:"pack"`
}
