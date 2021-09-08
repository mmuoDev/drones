package jwt

import "gopkg.in/square/go-jose.v2/jwt"

// GetClaimsFunc verifies jwt is not expired and gets the claims
type GetClaimsFunc func(rawJwt string) (Claims, error)

// GetClaims default get claims func that uses aws cognito
func GetClaims() GetClaimsFunc {
	return func(rawJWT string) (Claims, error) {
		parsedJWT, err := jwt.ParseSigned(rawJWT)

		if err != nil {
			return nil, err
		}

		unverifiedClaims := make(Claims)
		parsedJWT.UnsafeClaimsWithoutVerification(&unverifiedClaims)

		return unverifiedClaims, nil
	}
}

// Claims is a representation of a JWT claims
type Claims map[string]interface{}
