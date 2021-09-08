package db

import (
	"drones/internal"
	"embed"
	"encoding/json"
	"io/fs"

	"github.com/pkg/errors"
)

//go:embed mockdata/*
var content embed.FS

//RetrieveDroneByIDFunc returns functionality to retrieve drone by id
type RetrieveDroneByIDFunc func(id string) (internal.Drone, error)

//RetrieveDroneByID mocks a retrieval by drone ID
func RetrieveDroneByID() RetrieveDroneByIDFunc {
	return func(id string) (internal.Drone, error) {
		var drone internal.Drone
		files, err := fs.ReadDir(content, "mockdata")
		if err != nil {
			return drone, errors.Wrap(err, "db - unable to read files to mockdata dir")
		}
		for _, file := range files {
			if file.Name() == "retrieve_drone_details.json" {
				bb, err := content.ReadFile("mockdata/" + file.Name())
				if err != nil {
					return drone, errors.Wrap(err, "db - unable to read file from mockdata dir")
				}
				if err := json.Unmarshal(bb, &drone); err != nil {
					return drone, errors.Wrap(err, "db - Unable to unmarshal struct")
				}
				return drone, nil
			}
		}
		return drone, nil
	}
}
