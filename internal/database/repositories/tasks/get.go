package tasks

func (r *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task

	rows, err := r.DB.DB.Query(`SELECT * FROM tasks`)
	if err != nil {
		return []Task{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task

		if err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Deadline,
			&task.Duration,
			&task.Priority,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return []Task{}, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(id string) (Task, error) {
	row := r.DB.DB.QueryRow(`SELECT * FROM tasks WHERE id = ?`, id)

	var task Task
	if err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Deadline,
		&task.Duration,
		&task.Priority,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	); err != nil {
		return Task{}, err
	}

	return task, nil
}
