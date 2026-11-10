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

			router.Route("/request",func(router chi.Router){
				router.Post("/generate",createRequest)
				router.Get("/display/all",GetAllRequest)
				router.Get("display/{id}",GetRequestID)
				router.Delete("delete/{id}",DeleteRequest)
			
			})

			router.Route("/loadbalancer",func(router chi.Router){
				router.Post("/generate",CreateLoadBalancer)
				router.Get("/display/all",GetAllLoadBalancer)
				router.Get("display/{id}",GetLoadBalancerID)
				router.Delete("delete/{id}",DeleteLoadBalancer)
				router.Get("/display/servnodes/all",DisplayAllServerNodes)
				router.Get("display/requestcount",DisplayRequestCount)
				router.Put("update/servernode",AddServerNode)
				router.Put("update/delete/servernode",DeleteServerNode)
				
			})

		})
	})

	return router
}
