package learning

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/DoThat/internal/api/learning/dto"
	"github.com/ArteShow/DoThat/internal/database/repositories/learning/translations"
)

func (h *LearningHandler) CreateTranslationHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTranslationRequest
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

	id, err := h.LearningManager.CreateTransLation(translations.Translation{
		WordID:      req.WordID,
		Language:    req.Language,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
		Translation: req.Translation,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.CreateTranslationResponse{TranslationID: id}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LearningHandler) DeleteTranslationHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteTranslationRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.LearningManager.DeleteTranslation(req.TranslationID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetAllTranslationsHandler(w http.ResponseWriter, r *http.Request) {
	translations, err := h.LearningManager.GetAllTranslations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetAllTranslationsResponse{Translations: make([]dto.TranslationResponse, len(translations))}
	for _, t := range translations {
		resp.Translations = append(resp.Translations, dto.TranslationResponse{
			ID:          t.ID,
			WordID:      t.WordID,
			Language:    t.Language,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			Translation: t.Translation,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetTranslationByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetTranslationByIDRequest
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

	translation, err := h.LearningManager.GetTranslationByID(req.TranslationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetTranslationByIDResponse{Translation: dto.TranslationResponse{
		ID:          translation.ID,
		WordID:      translation.WordID,
		Language:    translation.Language,
		CreatedAt:   translation.CreatedAt,
		UpdatedAt:   translation.UpdatedAt,
		Translation: translation.Translation,
	}}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LearningHandler) GetTranslationByWordIDHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GetTranslationsByWordIDRequest
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

	translations, err := h.LearningManager.GetWordTranslations(req.WordID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetTranslationsByWordIDResponse{Translations: make([]dto.TranslationResponse, len(translations))}
	for _, t := range translations {
		resp.Translations = append(resp.Translations, dto.TranslationResponse{
			ID:          t.ID,
			WordID:      t.WordID,
			Language:    t.Language,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			Translation: t.Translation,
		})
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
