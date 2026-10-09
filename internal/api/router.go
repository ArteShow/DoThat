package api

import (
	"net/http"

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

	r.Mount("/api/v1/tasks", task_handler.NewTaskRouter(taskHandler))
	r.Mount("/api/v1/planner", planner_handler.NewRouter(plannerHandler))
	r.Mount("/api/v1/learning", learning_handler.NewRouter(learningHandler))

	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return &Router{
		Router: r,
	}
}
