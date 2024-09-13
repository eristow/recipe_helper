package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
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

func TestRootHandler(t *testing.T) {
	handler := NewRootHandler()
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expected := "Welcome to the Recipe Helper Backend!"
	assert.Equal(t, expected, rr.Body.String())
}

func TestRecipeHandler_List(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store, nil)

	req, err := http.NewRequest("GET", "/recipes", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var recipes []recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &recipes)
	require.NoError(t, err)
}

func TestRecipeHandler_Get(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store, nil)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	req, err := http.NewRequest("GET", "/recipes/"+testRecipe.Id.String(), nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	require.NoError(t, err)

	assert.Equal(t, testRecipe.Name, returnedRecipe.Name)
}

func TestRecipeHandler_Create(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store, nil)

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

func TestRecipeHandler_Update(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store, nil)

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

func TestRecipeHandler_Delete(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store, nil)

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

func TestRecipeHandler_Generate(t *testing.T) {
	store := database.NewDatastore()

	// Mock the API client
	mockClient := new(MockClient)

	// Mock the Generate function
	mockClient.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	handler := NewRecipeHandler(store, mockClient)

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

func TestEnableCors(t *testing.T) {
	w := httptest.NewRecorder()
	enableCors(w)

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

	corsHandler := HandleCors(handler)

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
