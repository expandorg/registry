package authentication

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	authenticationHeaderName = "Authorization"
	ExpirationKey            = "exp"
	UserIDKey                = "uid"
	IssuerKey                = "iss"
)

type AuthData struct {
	Expiration int64
	Issuer     string
	UserID     uint64
}

func ParseAuthData(ctx context.Context) (AuthData, error) {
	tokenString, err := extractAuthorizationHeaderFromContext(ctx)

	if err != nil {
		return AuthData{}, err
	}

	claims, err := parseJWT(tokenString)

	if err != nil {
		return AuthData{}, err
	}

	return AuthData{
		int64(claims[ExpirationKey].(float64)),
		claims[IssuerKey].(string),
		uint64(claims[UserIDKey].(float64)),
	}, nil
}

// extractAuthorizationHeaderFromContext finds and extracts the Authorization JWT from a context.
func extractAuthorizationHeaderFromContext(ctx context.Context) (string, error) {
	jwt, err := extractJWTFromContext(ctx)

	jwt = strings.Split(jwt, "Bearer ")[1]

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func extractJWTFromContext(ctx context.Context) (string, error) {
	errMissingAuthorizationHeader := fmt.Errorf("missing %q header", authenticationHeaderName)
	t, err := GetAuthFromContext(ctx)

	if len(t) == 0 || err != nil {
		return "", errMissingAuthorizationHeader
	}
	return t, nil
}

func parser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Unexpected signing method")
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, parser)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Unable to parse JWT")
}
