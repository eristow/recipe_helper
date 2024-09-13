package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/google/uuid"
	"github.com/ollama/ollama/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mocking the API client and response
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Generate(ctx context.Context, req *api.GenerateRequest, respFunc api.GenerateResponseFunc) error {
	args := m.Called(ctx, req, respFunc)

	// Simulate the response function being called with the correct DoneReason
	respFunc(api.GenerateResponse{
		Done:       true,
		DoneReason: "stop",
		Response:   `[{"name": "Recipe 1", "ingredients": ["Ingredient 1", "Ingredient 2"], "steps": ["Step 1", "Step 2"]}, {"name": "Recipe 2", "ingredients": ["Ingredient 3", "Ingredient 4"], "steps": ["Step 3", "Step 4"]}, {"name": "Recipe 3", "ingredients": ["Ingredient 5", "Ingredient 6"], "steps": ["Step 5", "Step 6"]}]`,
	})

	return args.Error(0)
}

func TestRecipeHandler_Generate(t *testing.T) {
	store := database.NewDatastore()

	// Mock the API client
	mockClient := new(MockClient)

	// Mock the Generate function
	mockClient.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	handler := rest.NewRecipeHandler(store, mockClient)

	ingredients := recipe.Ingredients{
		IngredientsList: []string{
			"Ingredient 1",
			"Ingredient 2",
			"Ingredient 3",
		},
	}
	body, _ := json.Marshal(ingredients)

	req, err := http.NewRequest("POST", "/recipes/generate", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var recipes []recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &recipes)
	require.NoError(t, err)

	assert.Equal(t, 3, len(recipes))

	for _, r := range recipes {
		assert.NotEqual(t, uuid.Nil, r.Id)
	}
}
