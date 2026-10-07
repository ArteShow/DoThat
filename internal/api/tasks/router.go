package tasks

import "github.com/go-chi/chi/v5"

func NewTaskRouter(handler *TaskHandler) *chi.Mux {
	r := chi.NewMux()

	r.Post("/", handler.CreateTask)
	r.Delete("/", handler.DeleteTask)
	r.Get("/", handler.GetAllTasks)
	r.Get("/by-id", handler.GetTaskByID)
	r.Patch("/status", handler.UpdateTaskStatus)
	r.Patch("/deadline", handler.UpdateTaskDeadline)
	r.Get("/error", handler.GetError)

	return r
}
