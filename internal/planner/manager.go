package planner

import (
	"database/sql"

	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/database/repositories/planner"
	"github.com/google/uuid"
)

type PlannerManager struct {
	PlannerRepo *planner.PlannerRepository
}

func NewPlannerManager(db *sql.DB) *PlannerManager {
	return &PlannerManager{
		PlannerRepo: planner.NewPlannerRepository(&database.Database{
			DB: db,
		}),
	}
}

func (m *PlannerManager) CreateEntry(entry planner.PlannerEntry) (string, error) {
	id := uuid.NewString()
	entry.ID = id

	if err := m.PlannerRepo.CreateEntry(entry); err != nil {
		return "", err
	}

	return id, nil
}

func (m *PlannerManager) DeleteEntry(id string) error {
	return m.PlannerRepo.DeleteEntry(id)
}

func (m *PlannerManager) GetAllEntries() ([]planner.PlannerEntry, error) {
	return m.PlannerRepo.GetAll()
}

func (m *PlannerManager) GetByID(id string) (planner.PlannerEntry, error) {
	return m.PlannerRepo.GetByID(id)
}
