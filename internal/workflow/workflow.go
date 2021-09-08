package workflow

import (
	"drones/internal/db"
	"drones/pkg"

	"github.com/pkg/errors"
)

//RetrieveLocationFunc returns functionality to retrieve location
type RetrieveLocationFunc func(params pkg.DNSQueryParams, sectorID string) (pkg.DNSResponse, error)

//RetrieveLocation retrieves location
func RetrieveLocation(retrieveDrone db.RetrieveDroneByIDFunc) RetrieveLocationFunc {
	return func(params pkg.DNSQueryParams, droneID string) (pkg.DNSResponse, error) {
		drone, err := retrieveDrone(droneID)
		if err != nil {
			return pkg.DNSResponse{}, errors.Wrapf(err, "workflow - error retrieving drone details for id=%s", droneID)
		}
		sectorID := float64(drone.SectorID)
		loc := params.XCoord*sectorID + params.YCoord*sectorID + params.ZCoord*sectorID + params.Velocity
		return pkg.DNSResponse{
			Location: loc,
		}, nil
	}
}
