package learning

import "github.com/go-chi/chi/v5"

func NewRouter(handler *LearningHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/packs", func(r chi.Router) {
		r.Post("/", handler.CreatePackHandler)
		r.Delete("/", handler.DeletePackHandler)
		r.Get("/", handler.GetAllPacksHandler)
		r.Get("/by-id", handler.GetPackByIDHandler)
	})

	r.Route("/words", func(r chi.Router) {
		r.Post("/", handler.CreateWordHandler)
		r.Delete("/", handler.DeleteWordHandler)
		r.Get("/", handler.GetAllWordsHandler)
		r.Get("/by-id", handler.GetWordByIDHandler)
		r.Get("/by-pack-id", handler.GetWordsByPackIDHandler)
	})

	r.Route("/translations", func(r chi.Router) {
		r.Post("/", handler.CreateTranslationHandler)
		r.Delete("/", handler.DeleteTranslationHandler)
		r.Get("/", handler.GetAllTranslationsHandler)
		r.Get("/by-id", handler.GetTranslationByIDHandler)
		r.Get("/by-word-id", handler.GetTranslationByWordIDHandler)
	})

	return r
}
