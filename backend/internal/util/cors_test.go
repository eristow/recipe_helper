package util_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestEnableCors(t *testing.T) {
	w := httptest.NewRecorder()
	util.EnableCors(w)

	expectedOrigin := "http://localhost:3000"
	actualOrigin := w.Header().Get("Access-Control-Allow-Origin")
	assert.Equal(t, expectedOrigin, actualOrigin)

	expectedMethods := "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	actualMethods := w.Header().Get("Access-Control-Allow-Methods")
	assert.Equal(t, expectedMethods, actualMethods)

	expectedHeaders := "Origin, Content-Type, X-Auth-Token"
	actualHeaders := w.Header().Get("Access-Control-Allow-Headers")
	assert.Equal(t, expectedHeaders, actualHeaders)
}

func TestHandleCors(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	corsHandler := util.HandleCors(handler)

	// Test OPTIONS request
	req, _ := http.NewRequest("OPTIONS", "/", nil)
	rr := httptest.NewRecorder()
	corsHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Test non-OPTIONS request
	req, _ = http.NewRequest("GET", "/", nil)
	rr = httptest.NewRecorder()
	corsHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
