package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRecipeNameIdFromUrl(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"/recipes/123", "123"},
		{"/recipes/", ""},
		{"/recipes", ""},
		{"/recipes/abc", "abc"},
		{"/invalid/123", ""},
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", tt.url, nil)
		assert.NoError(t, err)

		result := rest.GetRecipeNameIdFromUrl(req)
		assert.Equal(t, tt.expected, result)
	}
}
func TestRootHandler(t *testing.T) {
	handler := rest.NewRootHandler()
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expected := "Welcome to the Recipe Helper Backend!"
	assert.Equal(t, expected, rr.Body.String())
}

func TestRecipeHandler_ServeHTTP(t *testing.T) {
	store := database.NewDatastore()

	// Add a recipe to the store
	store.AddRecipe(recipe.NewRecipe(
		"Test Recipe",
		[]string{"Ingredient 1", "Ingredient 2"},
		[]string{"Step 1", "Step 2"},
	))

	testRecipe, _ := store.GetRecipeByName("Test Recipe")
	recipeUuid := testRecipe.Id.String()

	handler := rest.NewRecipeHandler(store, nil)

	tests := []struct {
		method         string
		target         string
		body           interface{}
		expectedStatus int
	}{
		{"POST", "/recipes", recipe.Recipe{Name: "Test Recipe"}, http.StatusCreated},
		{"PUT", fmt.Sprintf("/recipes/%s", recipeUuid), recipe.Recipe{Name: "Updated Recipe"}, http.StatusOK},
		{"GET", "/recipes", nil, http.StatusOK},
		{"GET", fmt.Sprintf("/recipes/%s", recipeUuid), nil, http.StatusOK},
		{"DELETE", fmt.Sprintf("/recipes/%s", recipeUuid), nil, http.StatusOK},
		// TODO: Figure out how to mock Ollama client for this test
		// {"POST", "/recipes/generate", recipe.Ingredients{IngredientsList: []string{"Ingredient 1", "Ingredient 2"}}, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.target, func(t *testing.T) {
			var body []byte
			var err error
			if tt.body != nil {
				body, err = json.Marshal(tt.body)
				require.NoError(t, err)
			}

			req, err := http.NewRequest(tt.method, tt.target, bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
