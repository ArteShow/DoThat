package database

func Init(migrationPath, dbPath string) (*Database, error) {
	db, err := Connect(dbPath)
	if err != nil {
		return &Database{}, err
	}

	if err = RunMigrations(db.DB, migrationPath); err != nil {
		return &Database{}, err
	}

	return db, nil
}
