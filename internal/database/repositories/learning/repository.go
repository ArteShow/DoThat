package learning

import (
	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"
)

type LearningRepositories struct {
	PacksRepo *packs.PacksRepository
}

func NewLearningRepositories(db *database.Database) *LearningRepositories {
	return &LearningRepositories{
		PacksRepo: packs.NewPacksRepository(db),
	}
}
