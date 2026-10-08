package learning

import (
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"
	"github.com/google/uuid"
)

func (m *LearningManager) CreatePack(pack packs.LearningPack) (string, error) {
	id := uuid.NewString()
	pack.ID = id

	return id, m.PacksRepo.CreatePack(pack)
}

func (m *LearningManager) DeletePack(id string) error {
	return m.PacksRepo.DeletePack(id)
}

func (m *LearningManager) GetAllPacks() ([]packs.LearningPack, error) {
	return m.PacksRepo.GetAll()
}

func (m *LearningManager) GetPackByID(id string) (packs.LearningPack, error) {
	return m.PacksRepo.GetByID(id)
}
