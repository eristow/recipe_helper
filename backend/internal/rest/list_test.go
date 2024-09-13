package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecipeHandler_List(t *testing.T) {
	store := database.NewDatastore()
	handler := rest.NewRecipeHandler(store, nil)

	req, err := http.NewRequest("GET", "/recipes", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var recipes []recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &recipes)
	require.NoError(t, err)
}
