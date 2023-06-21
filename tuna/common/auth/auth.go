package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lamt3/sushi/tuna/common/logger"
	uuid "github.com/satori/go.uuid"
)

var jwtKey = []byte("plum_secret")

type Claims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func BuildJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(8766 * time.Hour)
	// expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		logger.Error("error generating jwt for %s", userID)
		return "", err
	}

	return tokenString, err

}

type AuthToken struct {
	JWT    *jwt.Token
	UserID uuid.UUID
}

func DecryptJWT(tokenStr string) (*AuthToken, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		logger.Error("failed to decrypt jwt %s", tokenStr)
		return nil, err
	}

	userID := uuid.FromStringOrNil(claims.UserID)

	return &AuthToken{
		JWT:    tkn,
		UserID: userID,
	}, nil

}
