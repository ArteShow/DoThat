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

func (h *PlannerHandler) DeleteEntryHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.DeletePlannerEntryRequest
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

	if err = h.PlannerManager.DeleteEntry(req.EntryID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) GetAllEntriesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := h.PlannerManager.GetAllEntries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetAllEntriesResponse{Entries: make([]dto.EntryResponse, len(entries))}
	for _, e := range entries {
		resp.Entries = append(resp.Entries, dto.EntryResponse{
			ID:          e.ID,
			TaskID:      e.TaskID,
			Title:       e.Title,
			Description: e.Description,
			StartTime:   e.StartTime,
			EndTime:     e.EndTime,
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PlannerHandler) GetEntryByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetEntryByIDRequest
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

	entry, err := h.PlannerManager.GetByID(req.EntryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
