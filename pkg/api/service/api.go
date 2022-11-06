package service

import (
	"github.com/go-chi/chi/v5"
	"github.com/noon2dusk/go-things/db"
	"github.com/noon2dusk/go-things/pkg/api/domain"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Queries *db.Queries
	Logger  *logrus.Entry
}

func New(config Config) domain.Api {
	return &service{config.Queries, config.Logger}
}

type service struct {
	queries *db.Queries
	logger  *logrus.Entry
}

func (s *service) SetupRouter() chi.Router {
	router := chi.NewRouter()

	// TODO implement validation
	//router.Use(validate)

	router.Route("/spaceship", func(router chi.Router) {
		router.Post("/", s.addSpaceship)
		router.Get("/{id}", s.getSpaceship)
		router.Patch("/{id}", s.updateSpaceship)
		router.Delete("/{id}", s.deleteSpaceship)
	})

	router.Get("/spaceships", s.getSpaceships)

	return router
}
