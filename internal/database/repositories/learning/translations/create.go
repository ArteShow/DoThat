package translations

import "time"

func (r *TranslationsRepository) CreateNewTranslation(translation Translation) error {
	if _, err := r.DB.DB.Exec(`
	INSERT INTO learning_translations
	VALUES (?, ?, ?, ?, ?, ?)
	`, translation.ID, translation.WordID, translation.Language, translation.Translation, time.Now(), time.Now()); err != nil {
		return err
	}

	return nil
}
