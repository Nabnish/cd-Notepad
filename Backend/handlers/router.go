package handlers

import (
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	 
	
)


func CreateRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			router.Get("/health", Health)
			router.Post("/generaterequest",createRequest)
		})
	})

	return router
}
