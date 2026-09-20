package packs

func (r *PacksRepository) GetAll() ([]LearningPack, error) {
	var packs []LearningPack

	row, err := r.DB.DB.Query(`
	SELECT * FROM learning_packs
	`)

	if err != nil {
		return []LearningPack{}, err
	}
	defer row.Close()

	for row.Next() {
		var pack LearningPack
		if err = row.Scan(
			&pack.ID,
			&pack.Name,
			&pack.Description,
			&pack.Language,
			&pack.CreatedAt,
			&pack.UpdatedAt,
		); err != nil {
			return []LearningPack{}, err
		}

		packs = append(packs, pack)
	}

	return packs, nil
}

func (r *PacksRepository) GetByID(id string) (LearningPack, error) {
	rows := r.DB.DB.QueryRow(`
	SELECT * FROM learning_packs WHERE id = ?
	`, id)

	var pack LearningPack
	if err := rows.Scan(
		&pack.ID,
		&pack.Name,
		&pack.Description,
		&pack.Language,
		&pack.CreatedAt,
		&pack.UpdatedAt,
	); err != nil {
		return LearningPack{}, err
	}

	if err := rows.Err(); err != nil {
		return LearningPack{}, rows.Err()
	}

	return pack, nil
}
