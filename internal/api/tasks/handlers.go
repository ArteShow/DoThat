package tasks

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/DoThat/internal/api/tasks/dto"
	"github.com/ArteShow/DoThat/internal/database/repositories/tasks"
	manager "github.com/ArteShow/DoThat/internal/tasks"
)

type TaskHandler struct {
	TaskManager *manager.TaskManager
}

func NewTaskHandler(m *manager.TaskManager) *TaskHandler {
	return &TaskHandler{
		TaskManager: m,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest
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

	id := h.TaskManager.CreateTask(tasks.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
		Priority:    req.Priority,
		Duration:    req.Duration,
		Status:      req.Status,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	})

	if err = json.NewEncoder(w).Encode(dto.CreateTaskResponse{TaskID: id}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteTaskRequest
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

	h.TaskManager.DeleteTask(req.TaskID)

	w.WriteHeader(http.StatusOK)
}
