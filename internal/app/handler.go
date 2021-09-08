package app

import (
	"drones/internal/db"
	"drones/internal/jwt"
	"drones/internal/utils"
	"drones/internal/workflow"
	"drones/pkg"
	"net/http"
	"strings"

	//"github.com/mmuoDev/commons/httputils"
	"github.com/pkg/errors"
)

//RetrieveLocationHandler returns a http request to retrieve location
func RetrieveLocationHandler(getClaims jwt.GetClaimsFunc, retrieveDrone db.RetrieveDroneByIDFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var params pkg.DNSQueryParams
		if err := utils.GetQueryParams(&params, r); err != nil {
			utils.ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Retrieve sectorID from jwt
		token, err := getToken(r)
		if err != nil {
			utils.ServeError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := getClaims(token)
		if err != nil {
			utils.ServeError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		droneID, ok := claims["droneID"]
		if !ok {
			utils.ServeError(w, "droneID not in jwt claims", http.StatusUnauthorized)
			return
		}
		retrieveLocation := workflow.RetrieveLocation(retrieveDrone)
		res, err := retrieveLocation(params, droneID.(string))
		if err != nil {
			utils.ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.ServeJSON(res, w, http.StatusOK)
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
