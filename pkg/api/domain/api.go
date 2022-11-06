package domain

import (
	"github.com/go-chi/chi/v5"
)

type Api interface {
	SetupRouter() chi.Router
}
