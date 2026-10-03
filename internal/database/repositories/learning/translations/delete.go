package translations

func (r *TranslationsRepository) DeleteTranslation(id string) error {
	if _, err := r.DB.DB.Exec(`
	DELETE FROM learning_translations WHERE id = ?
	`, id); err != nil {
		return err
	}

	return nil
}
