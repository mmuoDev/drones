package app_test

import (
	"drones/internal"
	"drones/internal/app"
	"drones/internal/jwt"
	"drones/internal/utils"
	"drones/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validToken   = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzEwOTQ0NjksImV4cCI6MTY2MjYzMDQ2OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsImRyb25lSUQiOiIwMjQ5NzQxOTE3In0.KgVdBOEx-DiTa5VeGxJGyXMz1oJAv0xE3JiyQ-pZ-y4"
	invalidToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzEwOTQ0NjksImV4cCI6MTY2MjYzMDQ2OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsInVzZXJJRCI6IjEuMjMifQ.LCckOdq5W2rq1EJQ-PV79rL_Ue1OMKXDFW46F4Iyqy4"
)

func testServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)
	return ts.URL, func() { ts.Close() }
}

func TestRetrieveLocationWithInvalidClaimsReturns401(t *testing.T) {
	retrieveJWTClaimsIsInvoked := false
	mockRetrieveJWTClaims := func(o *app.OptionalArgs) {
		o.GetClaims = func(rawJwt string) (jwt.Claims, error) {
			retrieveJWTClaimsIsInvoked = true
			claims := make(map[string]interface{})
			return claims, nil
		}
	}
	//optional args
	opts := []app.Options{
		mockRetrieveJWTClaims,
	}
	app := app.New(opts...)
	serverURL, cleanUpServer := testServer(app.Handler())
	defer cleanUpServer()

	x, y, z, velocity := "1.23", "2.44", "5.6", "2"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/locations?x=%s&y=%s&z=%s&velocity=%s", serverURL, x, y, z, velocity), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", invalidToken))
	client := &http.Client{}
	res, _ := client.Do(req)

	t.Run("Http Status Code is 401", func(t *testing.T) {
		assert.Equal(t, res.StatusCode, http.StatusUnauthorized)
	})

	t.Run("Retrieve claims from JWT is invoked", func(t *testing.T) {
		assert.True(t, retrieveJWTClaimsIsInvoked)
	})

}

func TestRetrieveLocationReturns200(t *testing.T) {
	retrieveJWTClaimsIsInvoked := false
	droneID := "0249741917"
	mockRetrieveJWTClaims := func(o *app.OptionalArgs) {
		o.GetClaims = func(rawJwt string) (jwt.Claims, error) {
			retrieveJWTClaimsIsInvoked = true
			claims := map[string]interface{}{"droneID": droneID}
			return claims, nil
		}
	}

	retrieveDrone := func(o *app.OptionalArgs) {
		o.RetrieveDrone = func(id string) (internal.Drone, error) {
			var drone internal.Drone
			_, err := utils.FileToStruct(filepath.Join("testdata", "retrieve_drone_details.json"), &drone)
			if err != nil {
				t.Log("unable to convert from file to struct")
			}
			return drone, nil
		}
	}
	//optional args
	opts := []app.Options{
		mockRetrieveJWTClaims,
		retrieveDrone,
	}

	app := app.New(opts...)
	serverURL, cleanUpServer := testServer(app.Handler())
	defer cleanUpServer()

	x, y, z, velocity := "1.23", "2.44", "5.6", "2"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/locations?x=%s&y=%s&z=%s&velocity=%s", serverURL, x, y, z, velocity), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	client := &http.Client{}
	res, _ := client.Do(req)

	t.Run("Http Status Code is 200", func(t *testing.T) {
		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("Retrieve claims from JWT is invoked", func(t *testing.T) {
		assert.True(t, retrieveJWTClaimsIsInvoked)
	})

	t.Run("Response Body is as expected", func(t *testing.T) {
		var (
			expectedResponse pkg.DNSResponse
			actualResponse   pkg.DNSResponse
		)
		json.NewDecoder(res.Body).Decode(&actualResponse)
		_, err := utils.FileToStruct(filepath.Join("testdata", "retrieve_location_response.json"), &expectedResponse)
		if err != nil {
			t.Log("unable to marshal file")
		}
		assert.Equal(t, expectedResponse, actualResponse)
	})
}
