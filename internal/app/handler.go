package app

import (
	"drones/internal/jwt"
	"drones/internal/workflow"
	"drones/pkg"
	"net/http"
	"strings"

	"github.com/mmuoDev/commons/httputils"
	"github.com/pkg/errors"
)

//RetrieveLocationHandler returns a http request to retrieve location
func RetrieveLocationHandler(getClaims jwt.GetClaimsFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params pkg.DNSQueryParams
		if err := httputils.GetQueryParams(&params, r); err != nil {
			httputils.ServeError(errors.Wrapf(err, "handler - unable to decode query params"), w)
			return
		}

		//Retrieve sectorID from jwt
		token, err := getToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, err := getClaims(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		sectorID, ok := claims["sectorID"]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		retrieveLocation := workflow.RetrieveLocation()
		res, err := retrieveLocation(params, sectorID.(string))
		if err != nil {
			httputils.ServeError(err, w)
			return
		}
		httputils.ServeJSON(res, w)
	}
}

//getToken returns token from the header
func getToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	s := strings.Split(bearerToken, " ")
	if len(s) == 2 {
		return s[1], nil
	}
	return "", errors.New("jwt - Token not found")
}
