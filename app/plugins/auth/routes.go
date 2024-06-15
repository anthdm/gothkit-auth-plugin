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

	authConfig := kit.AuthenticationConfig{
		AuthFunc:    AuthenticateUser,
		RedirectURL: "/login",
	}

	router.Group(func(auth chi.Router) {
		auth.Use(kit.WithAuthentication(authConfig, true))
		auth.Get("/profile", kit.Handler(HandleProfileShow))
		auth.Put("/profile", kit.Handler(HandleProfileUpdate))
	})
}
