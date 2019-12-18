package authentication

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gemsorg/registry/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func mockHeadRequest(t *testing.T, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Errorf("Expected request method: %q, but got: %q", http.MethodHead, r.Method)
		}
		w.WriteHeader(statusCode)
	}))
}

func TestExtractAuthorizationHeaderFromContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), authKey, "Bearer 123")
	actualToken, _ := extractAuthorizationHeaderFromContext(ctx)
	if actualToken != "123" {
		t.Fatalf("Expected token: %q, but got: %q", "123", actualToken)
	}
}

func TestParseAuthData(t *testing.T) {
	os.Setenv("JWT_SECRET", mock.JWT_SECRET)
	ctx := mock.MockContext{}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    AuthData
		wantErr bool
	}{
		{
			"it returns auth data",
			args{ctx},
			AuthData{1591960106, "http://localhost:3000", 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAuthData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAuthData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Issuer, got.Issuer)
			assert.Equal(t, tt.want.UserID, got.UserID)
		})
	}
}
