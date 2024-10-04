package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper/internal/database"
	"github.com/eristow/recipe_helper/internal/recipe"
	"github.com/google/uuid"
)

func TestRootHandler(t *testing.T) {
	handler := NewRootHandler()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Welcome to the Recipe Helper Backend!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRecipeHandler_List(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store)

	req, err := http.NewRequest("GET", "/recipes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var recipes []recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &recipes)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
}

func TestRecipeHandler_Get(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	req, err := http.NewRequest("GET", "/recipes/"+testRecipe.Id.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if returnedRecipe.Name != testRecipe.Name {
		t.Errorf("handler returned unexpected recipe: got %v want %v", returnedRecipe.Name, testRecipe.Name)
	}
}

func TestRecipeHandler_Create(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store)

	newRecipe := recipe.Recipe{Name: "New Test Recipe"}
	body, _ := json.Marshal(newRecipe)

	req, err := http.NewRequest("POST", "/recipes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if returnedRecipe.Name != newRecipe.Name {
		t.Errorf("handler returned unexpected recipe: got %v want %v", returnedRecipe.Name, newRecipe.Name)
	}

	if returnedRecipe.Id == uuid.Nil {
		t.Errorf("handler returned recipe with nil ID")
	}
}

func TestRecipeHandler_Update(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	updatedRecipe := recipe.Recipe{Name: "Updated Test Recipe"}
	body, _ := json.Marshal(updatedRecipe)

	req, err := http.NewRequest("PUT", "/recipes/"+testRecipe.Id.String(), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var returnedRecipe recipe.Recipe
	err = json.Unmarshal(rr.Body.Bytes(), &returnedRecipe)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if returnedRecipe.Name != updatedRecipe.Name {
		t.Errorf("handler returned unexpected recipe: got %v want %v", returnedRecipe.Name, updatedRecipe.Name)
	}

	if returnedRecipe.Id != testRecipe.Id {
		t.Errorf("handler returned recipe with different ID: got %v want %v", returnedRecipe.Id, testRecipe.Id)
	}
}

func TestRecipeHandler_Delete(t *testing.T) {
	store := database.NewDatastore()
	handler := NewRecipeHandler(store)

	testRecipe := &recipe.Recipe{Name: "Test Recipe"}
	testRecipe.SetId(testRecipe.NewRecipeId())
	store.AddRecipe(testRecipe)

	req, err := http.NewRequest("DELETE", "/recipes/"+testRecipe.Id.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the recipe was deleted
	_, exists := store.GetRecipeById(testRecipe.Id.String())
	if exists {
		t.Errorf("Recipe was not deleted from the store")
	}
}

func TestEnableCors(t *testing.T) {
	w := httptest.NewRecorder()
	enableCors(w)

	expectedOrigin := "http://localhost:3000"
	actualOrigin := w.Header().Get("Access-Control-Allow-Origin")
	if actualOrigin != expectedOrigin {
		t.Errorf("unexpected Access-Control-Allow-Origin header: got %v want %v", actualOrigin, expectedOrigin)
	}

	expectedMethods := "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	actualMethods := w.Header().Get("Access-Control-Allow-Methods")
	if actualMethods != expectedMethods {
		t.Errorf("unexpected Access-Control-Allow-Methods header: got %v want %v", actualMethods, expectedMethods)
	}

	expectedHeaders := "Origin, Content-Type, X-Auth-Token"
	actualHeaders := w.Header().Get("Access-Control-Allow-Headers")
	if actualHeaders != expectedHeaders {
		t.Errorf("unexpected Access-Control-Allow-Headers header: got %v want %v", actualHeaders, expectedHeaders)
	}
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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for OPTIONS: got %v want %v", status, http.StatusOK)
	}

	// Test non-OPTIONS request
	req, _ = http.NewRequest("GET", "/", nil)
	rr = httptest.NewRecorder()
	corsHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for GET: got %v want %v", status, http.StatusOK)
	}
}
