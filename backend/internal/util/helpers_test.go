package util_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	util.InternalServerError(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "500 internal server error", rr.Body.String())
}

func TestNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	util.NotFound(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "404 not found", rr.Body.String())
}

func TestNewTrue(t *testing.T) {
	result := util.NewTrue()
	assert.NotNil(t, result)
	assert.True(t, *result)
}

func TestNewFalse(t *testing.T) {
	result := util.NewFalse()
	assert.NotNil(t, result)
	assert.False(t, *result)
}

func TestFirstN(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"hello", 2, "he"},
		{"hello", 5, "hello"},
		{"hello", 10, "hello"},
		{"", 2, ""},
	}

	for _, tt := range tests {
		result := util.FirstN(tt.input, tt.n)
		assert.Equal(t, tt.expected, result)
	}
}
