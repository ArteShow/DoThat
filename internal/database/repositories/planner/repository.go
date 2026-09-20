package planner

import (
	"database/sql"
	"time"

	"github.com/ArteShow/DoThat/internal/database"
)

type PlannerRepository struct {
	DB *sql.DB
}

func NewPlannerRepository(db *database.Database) *PlannerRepository {
	return &PlannerRepository{
		DB: db.DB,
	}
}

type PlannerEntry struct {
	ID          string
	TaskID      string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
