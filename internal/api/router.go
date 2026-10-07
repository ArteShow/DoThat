package api

import (
	task_handler "github.com/ArteShow/DoThat/internal/api/tasks"
	task_service "github.com/ArteShow/DoThat/internal/tasks"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter(taskManager *task_service.TaskManager) *Router {
	r := chi.NewRouter()

	taskHandler := task_handler.NewTaskHandler(taskManager)

	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Mount("/", task_handler.NewTaskRouter(taskHandler))
	})

	return &Router{
		Router: r,
	}
}
