package words

import "time"

func (r *WordRepository) CreateWord(word Word) error {
	if _, err := r.DB.DB.Exec(`
	INSERT INTO learning_words
	VALUES (?, ?, ?, ?, ?)
	`, word.ID, word.PackID, word.Word, time.Now(), time.Now()); err != nil {
		return err
	}

	return nil
}
