package learning

import (
	"errors"

	"github.com/ArteShow/DoThat/internal/database/repositories/learning/translations"
	"github.com/google/uuid"
)

func (m *LearningManager) CreateTransLation(translation translations.Translation) (string, error) {
	id := uuid.NewString()
	translation.ID = id

	if _, err := m.WordsRepo.GetByID(translation.WordID); err != nil {
		return "", errors.New("No such words with id: " + translation.WordID)
	}

	return id, m.TranslationRepo.CreateNewTranslation(translation)
}

func (m *LearningManager) DeleteTranslation(id string) error {
	return m.TranslationRepo.DeleteTranslation(id)
}

func (m *LearningManager) GetAllTranslations() ([]translations.Translation, error) {
	return m.TranslationRepo.GetAll()
}

func (m *LearningManager) GetTranslationByID(id string) (translations.Translation, error) {
	return m.TranslationRepo.GetByID(id)
}

func (m *LearningManager) GetWordTranslations(wordID string) ([]translations.Translation, error) {
	if _, err := m.WordsRepo.GetByID(wordID); err != nil {
		return []translations.Translation{}, err
	}

	allTranslations, err := m.TranslationRepo.GetAll()
	if err != nil {
		return []translations.Translation{}, err
	}

	var wordTranslations []translations.Translation
	for _, t := range allTranslations {
		if t.WordID == wordID {
			wordTranslations = append(wordTranslations, t)
		}
	}

	return wordTranslations, nil
}
