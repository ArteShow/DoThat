package learning

import "github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"

func (m *LearningManager) CreatePack(pack packs.LearningPack) error {
	return m.PacksRepo.CreatePack(pack)
}

func (m *LearningManager) DeletePack(id string) error {
	return m.PacksRepo.DeletePack(id)
}

func (m *LearningManager) GetAll() ([]packs.LearningPack, error) {
	return m.PacksRepo.GetAll()
}

func (m *LearningManager) GetByID(id string) (packs.LearningPack, error) {
	return m.PacksRepo.GetByID(id)
}
