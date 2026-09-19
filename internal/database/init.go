package database

func Init(path string) (*Database, error) {
	db, err := Connect(path)
	if err != nil {
		return &Database{}, err
	}

	if err = RunMigrations(db.DB, path); err != nil {
		return &Database{}, err
	}

	return db, nil
}
