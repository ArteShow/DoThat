package api

import (
	learning_handler "github.com/ArteShow/DoThat/internal/api/learning"
	planner_handler "github.com/ArteShow/DoThat/internal/api/planner"
	task_handler "github.com/ArteShow/DoThat/internal/api/tasks"
	learning_service "github.com/ArteShow/DoThat/internal/learning"
	planer_service "github.com/ArteShow/DoThat/internal/planner"
	task_service "github.com/ArteShow/DoThat/internal/tasks"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter(taskManager *task_service.TaskManager, plannerManager *planer_service.PlannerManager, learningManager *learning_service.LearningManager) *Router {
	r := chi.NewRouter()

	taskHandler := task_handler.NewTaskHandler(taskManager)
	plannerHandler := planner_handler.NewPlannerHandler(plannerManager)
	learningHandler := learning_handler.NewLearningHandler(learningManager)

	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Mount("/", task_handler.NewTaskRouter(taskHandler))
		r.Mount("/", planner_handler.NewRouter(plannerHandler))
		r.Mount("/", learning_handler.NewRouter(learningHandler))
	})

	return &Router{
		Router: r,
	}
}
