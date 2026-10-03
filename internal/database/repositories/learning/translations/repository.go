package translations

import (
	"time"

	"github.com/ArteShow/DoThat/internal/database"
)

type TranslationsRepository struct {
	DB *database.Database
}

func NewTranslationsRepository(db *database.Database) *TranslationsRepository {
	return &TranslationsRepository{
		DB: db,
	}
}

type Translation struct {
	ID          string
	WordID      string
	Language    string
	Translation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
