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

type CreateWordRequest struct {
	PackID    string    `json:"pack_id"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteWordRequest struct {
	WordID string `json:"word_id"`
}

type GetWordByIDRequest struct {
	WordID string `json:"word_id"`
}

type GetWordsByPackIDRequest struct {
	PackID string `json:"pack_id"`
}

type CreateTranslationRequest struct {
	WordID      string    `json:"word_id"`
	Language    string    `json:"language"`
	Translation string    `json:"translation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeleteTranslationRequest struct {
	TranslationID string `json:"translation_id"`
}

type GetTranslationByIDRequest struct {
	TranslationID string `json:"translation_id"`
}

type GetTranslationsByWordIDRequest struct {
	WordID string `json:"word_id"`
}
