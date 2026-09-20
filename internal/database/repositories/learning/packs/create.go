package packs

func (r *PacksRepository) CreatePack(pack LearningPack) error {
	if _, err := r.DB.DB.Exec(`
	INSERT INTO learning_packs
	VALUES (?, ?, ?, ?, ?, ?)
	`, pack.ID, pack.Name, pack.Description, pack.Language, pack.CreatedAt, pack.UpdatedAt); err != nil {
		return err
	}

	return nil
}
