package main

import "github.com/go-chi/chi"

func SetupRouter() *chi.Mux {

	router := chi.NewRouter()

	//Публичные ссылки
	publicRouter := chi.NewRouter()
	publicRouter.Post("/login", handleLogin)
	publicRouter.Post("/registration", nil) //добавь хендлер
	router.Mount("/", publicRouter)

	//Приватные ссылки
	protectedRouter := chi.NewRouter()
	protectedRouter.Use(JWTAuthMiddleware)
	protectedRouter.Post("/address/search", HandleSearch)
	protectedRouter.Post("/address/geocode", HandleGeocode)
	router.Mount("/api", protectedRouter)

	return router
}
