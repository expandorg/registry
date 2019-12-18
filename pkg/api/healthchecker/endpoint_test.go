package healthchecker

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthyResponse(t *testing.T) {
	want := `{"healthy":true}`

	actual, _ := json.Marshal(HealthyResponse{true})
	assert.Equal(t, want, string(actual))
}
