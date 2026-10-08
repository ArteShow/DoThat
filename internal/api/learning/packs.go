package learning

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/DoThat/internal/api/learning/dto"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/packs"
)

func (h *LearningHandler) CreatePackHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePackRequest
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

	id, err := h.LearningManager.CreatePack(packs.LearningPack{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
		Language:    req.Language,
	})

	resp := dto.CreatePackResponse{PackID: id}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LearningHandler) DeletePackHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.DeletePackRequest
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

	if err = h.LearningManager.DeletePack(req.PackID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
