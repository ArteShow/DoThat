package packs

func (r *PacksRepository) DeletePack(id string) error {
	_, err := r.DB.DB.Exec(`
	DELETE FROM learning_packs WHERE id = ?
	`, id)

	if err != nil {
		return err
	}

	return nil
}
