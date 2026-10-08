package planner

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/DoThat/internal/api/planner/dto"
	planer_repository "github.com/ArteShow/DoThat/internal/database/repositories/planner"
	planer_service "github.com/ArteShow/DoThat/internal/planner"
)

type PlannerHandler struct {
	PlannerManager *planer_service.PlannerManager
}

func NewPlannerHandler(m *planer_service.PlannerManager) *PlannerHandler {
	return &PlannerHandler{
		PlannerManager: m,
	}
}

func (h *PlannerHandler) CreateEntryHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePlannerEntryRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.PlannerManager.CreateEntry(planer_repository.PlannerEntry{
		TaskID:      req.TaskID,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.CreatePlannerEntryResponse{EntryID: id}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
