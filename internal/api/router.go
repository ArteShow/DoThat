package api

import (
	planner_handler "github.com/ArteShow/DoThat/internal/api/planner"
	task_handler "github.com/ArteShow/DoThat/internal/api/tasks"
	planer_service "github.com/ArteShow/DoThat/internal/planner"
	task_service "github.com/ArteShow/DoThat/internal/tasks"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter(taskManager *task_service.TaskManager, plannerManager *planer_service.PlannerManager) *Router {
	r := chi.NewRouter()

	taskHandler := task_handler.NewTaskHandler(taskManager)
	plannerHandler := planner_handler.NewPlannerHandler(plannerManager)

	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Mount("/", task_handler.NewTaskRouter(taskHandler))
		r.Mount("/", planner_handler.NewRouter(plannerHandler))
	})

	return &Router{
		Router: r,
	}
}
