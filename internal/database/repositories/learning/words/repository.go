package words

import (
	"time"

	"github.com/ArteShow/DoThat/internal/database"
)

type WordRepository struct {
	DB *database.Database
}

func NewWordRepository(db *database.Database) *WordRepository {
	return &WordRepository{
		DB: db,
	}
}

type Word struct {
	ID        string
	PackID    string
	Word      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
