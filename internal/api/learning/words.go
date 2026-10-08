package learning

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/DoThat/internal/api/learning/dto"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/words"
)

func (h *LearningHandler) CreateWordHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateWordRequest
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

	id, err := h.LearningManager.CreateWord(words.Word{
		Word:      req.Word,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		PackID:    req.PackID,
	})

	resp := dto.CreateWordResponse{WordID: id}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LearningHandler) DeleteWordHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteWordRequest
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

	if err = h.LearningManager.DeleteWord(req.WordID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetAllWordsHandler(w http.ResponseWriter, r *http.Request) {
	words, err := h.LearningManager.GetAllWords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetAllWordsResponse{Words: make([]dto.WordResponse, len(words))}
	for _, w := range words {
		resp.Words = append(resp.Words, dto.WordResponse{
			ID:        w.ID,
			Word:      w.Word,
			PackID:    w.PackID,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetWordByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetWordByIDRequest
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

	word, err := h.LearningManager.GetWordByID(req.WordID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetWordByIDResponse{Word: dto.WordResponse{
		ID:        word.ID,
		Word:      word.Word,
		PackID:    word.PackID,
		CreatedAt: word.CreatedAt,
		UpdatedAt: word.UpdatedAt,
	}}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetWordsByPackIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetWordsByPackIDRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	words, err := h.LearningManager.GetWordsByPackID(req.PackID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetWordsByPackIDResponse{Words: make([]dto.WordResponse, len(words))}
	for _, w := range words {
		resp.Words = append(resp.Words, dto.WordResponse{
			ID:        w.ID,
			Word:      w.Word,
			PackID:    w.PackID,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
