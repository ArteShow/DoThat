package learning

import (
	"database/sql"

	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/translations"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/words"
)

type LearningManager struct {
	WordsRepo       *words.WordRepository
	PacksRepo       *packs.PacksRepository
	TranslationRepo *translations.TranslationsRepository
}

func NewLearningManager(db *sql.DB) *LearningManager {
	return &LearningManager{
		WordsRepo: words.NewWordRepository(
			&database.Database{DB: db},
		),
		PacksRepo: packs.NewPacksRepository(
			&database.Database{DB: db},
		),
		TranslationRepo: translations.NewTranslationsRepository(
			&database.Database{DB: db},
		),
	}
}
