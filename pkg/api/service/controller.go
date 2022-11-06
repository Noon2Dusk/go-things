package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/noon2dusk/go-things/db"
	"net/http"
	"strconv"
)

func (s *service) getSpaceships(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	spaceships, err := s.queries.GetSpaceShips(ctx)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	data, err := s.marshalData(spaceships)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	w.Write(data)
}

func (s *service) getSpaceship(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	spaceship, err := s.queries.GetSpaceShipById(ctx, int32(intId))
	data, err := s.marshalData(spaceship)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	w.Write(data)
}

type AddSpaceship struct {
	Name      string             `json:"name"`
	Class     string             `json:"class"`
	Crew      int32              `json:"crew"`
	Image     string             `json:"image"`
	Value     string             `json:"value"`
	Status    db.SpaceshipStatus `json:"status"`
	Armaments string             `json:"armaments"`
}

func (s *service) addSpaceship(w http.ResponseWriter, r *http.Request) {
	body, err := s.readBody(r)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	var spaceship AddSpaceship
	err = json.Unmarshal(body, &spaceship)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	err = s.queries.InsertSpaceship(r.Context(), db.InsertSpaceshipParams{
		Name:      spaceship.Name,
		Class:     spaceship.Class,
		Crew:      spaceship.Crew,
		Image:     spaceship.Image,
		Value:     spaceship.Value,
		Status:    spaceship.Status,
		Armaments: spaceship.Armaments,
	})
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	s.writeSuccess(w)
}

type UpdateSpaceship struct {
	Name      *string             `json:"name"`
	Class     *string             `json:"class"`
	Crew      *int32              `json:"crew"`
	Image     *string             `json:"image"`
	Value     *string             `json:"value"`
	Status    *db.SpaceshipStatus `json:"status"`
	Armaments *string             `json:"armaments"`
}

func (s *service) updateSpaceship(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	spaceshipDB, err := s.queries.GetSpaceShipById(ctx, int32(intId))
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	body, err := s.readBody(r)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	var spaceship UpdateSpaceship
	err = json.Unmarshal(body, &spaceship)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	err = s.queries.UpdateSpaceship(r.Context(), s.createUpdateSpaceshipStruct(spaceship, spaceshipDB))
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	s.writeSuccess(w)
}

func (s *service) deleteSpaceship(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	err = s.queries.DeleteSpaceshipById(ctx, int32(intId))
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	s.writeSuccess(w)
}
