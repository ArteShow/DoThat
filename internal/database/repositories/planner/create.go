package planner

func (r *PlannerRepository) CreateEntry(entry PlannerEntry) error {
	_, err := r.DB.Query(`
	INSERT INTO planner_entries 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, entry.ID,
		entry.TaskID,
		entry.Title,
		entry.Description,
		entry.StartTime,
		entry.EndTime,
		entry.CreatedAt,
		entry.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
