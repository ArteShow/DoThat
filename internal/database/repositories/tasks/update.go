package tasks

import "time"

func (r *TaskRepository) UpdateStatus(id, status string) error {
	_, err := r.DB.DB.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, status, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) UpdateDeadline(newDeadline time.Time, id string) error {
	_, err := r.DB.DB.Exec(`UPDATE tasks SET deadline = ? WHERE id = ?`, newDeadline, id)
	if err != nil {
		return err
	}

	return nil
}
