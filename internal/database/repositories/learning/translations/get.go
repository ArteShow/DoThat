package translations

func (r *TranslationsRepository) GetAll() ([]Translation, error) {
	var translations []Translation

	rows, err := r.DB.DB.Query(`
	SELECT * FROM learning_translations
	`)

	if err != nil {
		return []Translation{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var translation Translation

		if err := rows.Scan(
			&translation.ID,
			&translation.WordID,
			&translation.Language,
			&translation.Translation,
			&translation.CreatedAt,
			&translation.UpdatedAt,
		); err != nil {
			return []Translation{}, err
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

func (r *TranslationsRepository) GetByID(id string) (Translation, error) {
	row := r.DB.DB.QueryRow(`
	SELECT * FROM learning_translations WHERE id = ?
	`, id)

	var translation Translation

	if err := row.Scan(
		&translation.ID,
		&translation.WordID,
		&translation.Language,
		&translation.Translation,
		&translation.CreatedAt,
		&translation.UpdatedAt,
	); err != nil {
		return Translation{}, err
	}

	if err := row.Err(); err != nil {
		return Translation{}, row.Err()
	}

	return translation, nil
}
