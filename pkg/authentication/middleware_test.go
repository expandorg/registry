package authentication

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/expandorg/registry/pkg/mock"
)

func TestAuthorizationHeaderIsInsertedInContextWithBearer(t *testing.T) {
	os.Setenv("JWT_SECRET", mock.JWT_SECRET)
	expectedToken, _ := mock.GenerateJWT(1)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualToken, err := GetAuthFromContext(r.Context())
		if err != nil {
			t.Fatalf("Expected token in context")
		}

		if actualToken != "Bearer "+expectedToken {
			t.Fatalf("Unexpected token: %q, but got: %q", expectedToken, actualToken)
		}
	})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+expectedToken)

	AuthMiddleware(handler).ServeHTTP(nil, req)
}

func TestGetAuthFromContextReturnsError(t *testing.T) {
	if _, err := GetAuthFromContext(context.Background()); err == nil {
		t.Fatalf("Expected error when retrieving auth from empty context")
	}
}
