package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecipeHandler_Create(t *testing.T) {
	store := database.NewDatastore()
	handler := rest.NewRecipeHandler(store, nil)

	newRecipe := recipe.Recipe{Name: "New Test Recipe"}
	body, _ := json.Marshal(newRecipe)

	req, err := http.NewRequest("POST", "/recipes", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	require.NoError(t, err)

	assert.Equal(t, newRecipe.Name, returnedRecipe.Name)
	assert.NotEqual(t, uuid.Nil, returnedRecipe.Id)
}
