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

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.TaskManager.GetAll()

	var resp dto.GetAllTasksResponse
	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Deadline:    t.Deadline,
			Duration:    t.Duration,
			Status:      t.Status,
			Priority:    t.Priority,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	var req dto.GetTaskByIDRequest
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

	task := h.TaskManager.GetByID(req.TaskID)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTaskStatusRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.TaskManager.UpdateStatus(req.TaskID, req.NewStatus)

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) UpdateTaskDeadline(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateTaskDeadlineRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.TaskManager.UpdateDeadline(req.TaskID, req.NewDeadline)

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) GetError(w http.ResponseWriter, r *http.Request) {
	taskErr := h.TaskManager.GetError()

	resp := dto.GetErrorResponse{Error: taskErr.Error()}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
