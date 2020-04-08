package util

import (
	"api-dashboard/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

// ParseToken parsing token
func ParseToken(token string) (*models.ClaimsData, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &models.ClaimsData{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.ClaimsData); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
