package workflow

import (
	"drones/pkg"
	"strconv"
	"github.com/pkg/errors"
)

//RetrieveLocationFunc returns functionality to retrieve location
type RetrieveLocationFunc func(params pkg.DNSQueryParams, sectorID string) (pkg.DNSResponse, error)

//RetrieveLocation retrieves location 
func RetrieveLocation() RetrieveLocationFunc {
	return func(params pkg.DNSQueryParams, sectorID string) (pkg.DNSResponse, error) {
		sID, err := stringToFloat(sectorID, 64)
		if err != nil {
			return pkg.DNSResponse{}, errors.Wrapf(err, "workflow - error converting=%s to float", sectorID)
		}
		loc := params.XCoord*sID + params.YCoord*sID + params.ZCoord*sID + params.Velocity
		return pkg.DNSResponse{
			Location: loc,
		}, nil
	}
}

//stringToFloat converts a string to float
func stringToFloat(s string, b int) (float64, error) {
	f, err := strconv.ParseFloat(s, b)
	if err != nil {
		return 0, err
	}
	return f, nil
}
