package words

func (r *WordRepository) GetAll() ([]Word, error) {
	var words []Word

	rows, err := r.DB.DB.Query(`
	SELECT * FROM learning_words
	`)

	if err != nil {
		return []Word{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var word Word

		if err := rows.Scan(
			&word.ID,
			&word.PackID,
			&word.Word,
			&word.CreatedAt,
			&word.UpdatedAt,
		); err != nil {
			return []Word{}, err
		}

		words = append(words, word)
	}

	return words, nil
}

func (r *WordRepository) GetByID(id string) (Word, error) {
	row := r.DB.DB.QueryRow(`
	SELECT * FROM learning_words WHERE id = ?
	`, id)

	var word Word

	if err := row.Scan(
		&word.ID,
		&word.PackID,
		&word.Word,
		&word.CreatedAt,
		&word.UpdatedAt,
	); err != nil {
		return Word{}, err
	}

	if err := row.Err(); err != nil {
		return Word{}, row.Err()
	}

	return word, nil
}
