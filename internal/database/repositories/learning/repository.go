package learning

import (
	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/words"
)

type LearningRepositories struct {
	PacksRepo *packs.PacksRepository
	WordRepo  *words.WordRepository
}

func NewLearningRepositories(db *database.Database) *LearningRepositories {
	return &LearningRepositories{
		PacksRepo: packs.NewPacksRepository(db),
		WordRepo:  words.NewWordRepository(db),
	}
}
