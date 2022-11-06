package service

import (
	"bytes"
	"encoding/json"
	"github.com/noon2dusk/go-things/db"
	"net/http"
)

func (s *service) marshalData(data interface{}) ([]byte, error) {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return marshalledData, nil
}

func (s *service) readBody(r *http.Request) ([]byte, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *service) writeSuccess(w http.ResponseWriter) {
	data, _ := s.marshalData(SuccessfulResponse{Success: true})
	w.Write(data)
}

type SuccessfulResponse struct {
	Success bool
}

func (s *service) createUpdateSpaceshipStruct(spaceship UpdateSpaceship, spaceshipDB db.Spaceship) db.UpdateSpaceshipParams {
	var updateSpaceshipParams db.UpdateSpaceshipParams
	if spaceship.Name != nil {
		updateSpaceshipParams.Name = *spaceship.Name
	} else {
		updateSpaceshipParams.Name = spaceshipDB.Name
	}

	if spaceship.Class != nil {
		updateSpaceshipParams.Class = *spaceship.Class
	} else {
		updateSpaceshipParams.Class = spaceshipDB.Class
	}

	if spaceship.Crew != nil {
		updateSpaceshipParams.Crew = *spaceship.Crew
	} else {
		updateSpaceshipParams.Crew = spaceshipDB.Crew
	}

	if spaceship.Image != nil {
		updateSpaceshipParams.Image = *spaceship.Image
	} else {
		updateSpaceshipParams.Image = spaceshipDB.Image
	}

	if spaceship.Name != nil {
		updateSpaceshipParams.Name = *spaceship.Name
	} else {
		updateSpaceshipParams.Name = spaceshipDB.Name
	}

	if spaceship.Value != nil {
		updateSpaceshipParams.Value = *spaceship.Value
	} else {
		updateSpaceshipParams.Value = spaceshipDB.Value
	}

	if spaceship.Status != nil {
		updateSpaceshipParams.Status = *spaceship.Status
	} else {
		updateSpaceshipParams.Status = spaceshipDB.Status
	}

	if spaceship.Armaments != nil {
		updateSpaceshipParams.Armaments = *spaceship.Armaments
	} else {
		updateSpaceshipParams.Armaments = spaceshipDB.Armaments
	}
	updateSpaceshipParams.ID = spaceshipDB.ID
	return updateSpaceshipParams
}
