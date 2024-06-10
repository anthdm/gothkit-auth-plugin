package handlers

import (
	"net/http"

	"auth/app/types"

	"github.com/anthdm/gothkit/kit"
)

func HandleAuthentication(w http.ResponseWriter, r *http.Request) (kit.Auth, error) {
	return types.AuthUser{}, nil
}
