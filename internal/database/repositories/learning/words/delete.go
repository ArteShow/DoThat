package words

func (r *WordRepository) DeleteWord(id string) error {
	if _, err := r.DB.DB.Exec(`
	DELETE FROM learning_words WHERE id = ?
	`, id); err != nil {
		return err
	}

	return nil
}
