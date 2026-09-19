package tasks

func (r *TaskRepository) DeleteTask(id string) error {
	_, err := r.DB.DB.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}
