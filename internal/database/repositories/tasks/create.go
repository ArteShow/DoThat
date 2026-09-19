package tasks

func (r *TaskRepository) CreateTask(task Task) error {
	_, err := r.DB.DB.Exec(`
	INSERT INTO tasks
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		task.ID,
		task.Title,
		task.Description,
		task.Deadline,
		task.Duration,
		task.Priority,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
