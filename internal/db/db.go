package db

import (
	"drones/internal"
	"drones/internal/utils"
	"path/filepath"

	"github.com/pkg/errors"
)

//RetrieveDroneByIDFunc returns functionality to retrieve drone by id
type RetrieveDroneByIDFunc func(id string) (internal.Drone, error)

//RetrieveDroneByID mocks a retrieval by drone ID
func RetrieveDroneByID() RetrieveDroneByIDFunc {
	return func(id string) (internal.Drone, error) {
		var drone internal.Drone
		_, err := utils.FileToStruct(filepath.Join("testdata", "retrieve_drone_details.json"), &drone)
		if err != nil {
			return drone, errors.Wrapf(err, "db - unable to convert file to struct")
		}
		return drone, nil
	}
}
