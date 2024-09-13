package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecipeHandler_Delete(t *testing.T) {
	store := database.NewDatastore()
	handler := rest.NewRecipeHandler(store, nil)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	req, err := http.NewRequest("DELETE", "/recipes/"+testRecipe.Id.String(), nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify that the recipe was deleted
	_, exists := store.GetRecipeById(testRecipe.Id.String())
	assert.False(t, exists)
}
