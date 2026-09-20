package packs

import (
	"time"

	"github.com/ArteShow/DoThat/internal/database"
)

type PacksRepository struct {
	DB *database.Database
}

func NewPacksRepository(db *database.Database) *PacksRepository {
	return &PacksRepository{
		DB: db,
	}
}

type LearningPack struct {
	ID          string
	Name        string
	Description string
	Language    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
