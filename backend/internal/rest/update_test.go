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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecipeHandler_Update(t *testing.T) {
	store := database.NewDatastore()
	handler := rest.NewRecipeHandler(store, nil)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	updatedRecipe := recipe.Recipe{Name: "Updated Test Recipe"}
	body, _ := json.Marshal(updatedRecipe)

	req, err := http.NewRequest("PUT", "/recipes/"+testRecipe.Id.String(), bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	require.NoError(t, err)

	assert.Equal(t, updatedRecipe.Name, returnedRecipe.Name)
	assert.Equal(t, testRecipe.Id, returnedRecipe.Id)
}
