package auth

import (
	"github.com/anthdm/gothkit/kit"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(router chi.Router) {
	router.Get("/login", kit.Handler(HandleAuthIndex))
	router.Post("/login", kit.Handler(HandleAuthCreate))

	router.Get("/signup", kit.Handler(HandleSignupIndex))
	router.Post("/signup", kit.Handler(HandleSignupCreate))
}
