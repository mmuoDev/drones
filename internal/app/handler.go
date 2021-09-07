package app

import (
	"drones/internal/jwt"
	"drones/internal/workflow"
	"drones/pkg"
	"net/http"

	"github.com/mmuoDev/commons/httputils"
	"github.com/pkg/errors"
)

//RetrieveLocationHandler returns a http request to retrieve location
func RetrieveLocationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params pkg.DNSQueryParams
		if err := httputils.GetQueryParams(&params, r); err != nil {
			httputils.ServeError(errors.Wrapf(err, "handler - unable to decode query params"), w)
			return
		}

		//Retrieve sectorID from jwt
		token, err := jwt.GetTokenMetaData(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		sectorID := token.SectorID
		retrieveLocation := workflow.RetrieveLocation()
		res, err := retrieveLocation(params, sectorID)
		if err != nil {
			httputils.ServeError(err, w)
			return
		}
		httputils.ServeJSON(res, w)
	}
}
