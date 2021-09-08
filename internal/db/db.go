package db

import (
	"drones/internal"
)

//RetrieveDroneByIDFunc returns functionality to retrieve drone by id
type RetrieveDroneByIDFunc func(id string) (internal.Drone, error)

//RetrieveDroneByID mocks a retrieval by drone ID
func RetrieveDroneByID() RetrieveDroneByIDFunc {
	return func(id string) (internal.Drone, error) {
		//mock drone details
		drone := internal.Drone{
			ID:       "2846798240975",
			SectorID: 12,
		}
		return drone, nil
	}
}
