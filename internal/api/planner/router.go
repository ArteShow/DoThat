package planner

import "github.com/go-chi/chi/v5"

func NewRouter(handler *PlannerHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", handler.CreateEntryHandler)
	r.Get("/", handler.GetAllEntriesHandler)
	r.Get("/by-id", handler.GetEntryByIDHandler)
	r.Delete("/", handler.DeleteEntryHandler)

	return r
}
