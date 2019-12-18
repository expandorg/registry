package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MockContext struct{}

const JWT_SECRET = "abc"

func (ctx MockContext) Deadline() (deadline time.Time, ok bool) {
	return deadline, ok
}

func (ctx MockContext) Done() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}

func (ctx MockContext) Err() error {
	return context.DeadlineExceeded
}

func (ctx MockContext) Value(key interface{}) interface{} {
	switch fmt.Sprintf("%v", key) {
	case "auth":
		j, _ := GenerateJWT(1)
		return "Bearer " + j
	}
	return nil
}

func GenerateJWT(userID int) (string, error) {
	c := jwt.MapClaims{
		"exp": time.Now().Add(8760 * time.Hour).Unix(),
		"iss": "http://localhost:3000",
		"uid": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(JWT_SECRET))
}
