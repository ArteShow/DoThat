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

func (h *LearningHandler) GetAllPacksHandler(w http.ResponseWriter, r *http.Request) {
	packs, err := h.LearningManager.GetAllPacks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetAllPacksResponse{Packs: make([]dto.PackResponse, len(packs))}
	for _, p := range packs {
		resp.Packs = append(resp.Packs, dto.PackResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Language:    p.Language,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetPackByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetPackByIDRequest
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

	pack, err := h.LearningManager.GetPackByID(req.PackID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetPackByIDResponse{Pack: dto.PackResponse{
		ID:          pack.ID,
		Name:        pack.Name,
		Description: pack.Description,
		CreatedAt:   pack.CreatedAt,
		UpdatedAt:   pack.UpdatedAt,
		Language:    pack.Language,
	}}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
