package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	pkgErr "github.com/pkg/errors"
)

//TokenMetaData represents metadata of the token
type TokenMetaData struct {
	SectorID string
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

//verifyToken verifies signing method
func verifyToken(r *http.Request) (*jwt.Token, error) {
	ts, err := getToken(r)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

//GetTokenMetaData returns a token's metadata
func GetTokenMetaData(r *http.Request) (*TokenMetaData, error) {
	token, err := verifyToken(r)
	if err != nil {
		return &TokenMetaData{}, pkgErr.Wrap(err, "jwt - unable to retrieve token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sID, ok := claims["sector_id"].(string)
		if !ok {
			return nil, fmt.Errorf("sectorID not found in token")
		}
		return &TokenMetaData{
			SectorID: sID,
		}, nil
	}
	return &TokenMetaData{}, err
}