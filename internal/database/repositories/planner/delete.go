package planner

func (r *PlannerRepository) DeleteEntry(id string) error {
	_, err := r.DB.Exec(`
	DELETE FROM planner_entries
	WHERE id = ?
	`, id)

	if err != nil {
		return err
	}

	return nil
}
