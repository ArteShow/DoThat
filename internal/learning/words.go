package learning

import (
	"errors"

	"github.com/ArteShow/DoThat/internal/database/repositories/learning/words"
)

func (m *LearningManager) CreateWord(word words.Word) error {
	if _, err := m.PacksRepo.GetByID(word.PackID); err != nil {
		return errors.New("No pack found with pack id: " + word.PackID)
	}

	return m.WordsRepo.CreateWord(word)
}

func (m *LearningManager) DeleteWord(id string) error {
	return m.WordsRepo.DeleteWord(id)
}

func (m *LearningManager) GetAllWords() ([]words.Word, error) {
	return m.WordsRepo.GetAll()
}

func (m *LearningManager) GetWordByID(id string) (words.Word, error) {
	return m.WordsRepo.GetByID(id)
}

func (m *LearningManager) GetWordsByPackID(packID string) ([]words.Word, error) {
	if _, err := m.PacksRepo.GetByID(packID); err != nil {
		return []words.Word{}, err
	}

	allWords, err := m.WordsRepo.GetAll()
	if err != nil {
		return []words.Word{}, err
	}

	var packWords []words.Word
	for _, w := range allWords {
		if w.PackID == packID {
			packWords = append(packWords, w)
		}
	}

	return packWords, nil
}
