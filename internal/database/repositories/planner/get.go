package planner

func (r *PlannerRepository) GetAll() ([]PlannerEntry, error) {
	var entries []PlannerEntry

	rows, err := r.DB.Query(`SELECT * FROM planner_entries`)

	if err != nil {
		return []PlannerEntry{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var entry PlannerEntry

		if err = rows.Scan(
			&entry.ID,
			&entry.TaskID,
			&entry.Title,
			&entry.Description,
			&entry.StartTime,
			&entry.EndTime,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		); err != nil {
			return []PlannerEntry{}, err
		}

		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		return []PlannerEntry{}, err
	}

	return entries, err
}

func (r *PlannerRepository) GetByID(id string) (PlannerEntry, error) {
	row := r.DB.QueryRow(`SELECT * FROM planner_entries WHERE id = ?`, id)

	var entry PlannerEntry
	if err := row.Scan(
		&entry.ID,
		&entry.TaskID,
		&entry.Title,
		&entry.Description,
		&entry.StartTime,
		&entry.EndTime,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		return PlannerEntry{}, err
	}

	return entry, nil
}
