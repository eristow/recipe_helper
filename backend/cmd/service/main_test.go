package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
)

func TestSlashFix(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.URL.Path))
	})

	slashFixHandler := &slashFix{mux: handler}

	testCases := []struct {
		name           string
		path           string
		expectedPath   string
		expectedStatus int
	}{
		{"No double slashes", "/recipes/1", "/recipes/1", http.StatusOK},
		{"Double slashes", "/recipes//1", "/recipes/1", http.StatusOK},
		{"Multiple double slashes", "/recipes///1//2", "/recipes/1/2", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			slashFixHandler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}

			if rr.Body.String() != tc.expectedPath {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tc.expectedPath)
			}
		})
	}
}

func TestRoutes(t *testing.T) {
	ds := database.NewDatastore()
	rootH := rest.NewRootHandler()
	recipeH := rest.NewRecipeHandler(ds, nil)

	testRecipe := recipe.NewRecipe(
		"Test Recipe",
		[]string{"Ingredient 1", "Ingredient 2"},
		[]string{"Step 1", "Step 2"},
	)
	ds.AddRecipe(testRecipe)

	mux := http.NewServeMux()
	mux.Handle("/", rest.HandleCors(rootH))
	mux.Handle("/recipes/", rest.HandleCors(recipeH))
	mux.Handle("/recipes", rest.HandleCors(recipeH))

	testCases := []struct {
		name           string
		method         string
		path           string
		body           []byte
		expectedStatus int
	}{
		{"Root GET", "GET", "/", nil, http.StatusOK},
		{"Recipes List GET", "GET", "/recipes", nil, http.StatusOK},
		{"Recipe GET", "GET", "/recipes/" + testRecipe.Id.String(), nil, http.StatusOK},
		{"Recipe POST", "POST", "/recipes", []byte(`{"name":"New Recipe","ingredients":["Ingredient"],"steps":["Step"]}`), http.StatusCreated},
		{"Recipe PUT", "PUT", "/recipes/" + testRecipe.Id.String(), []byte(`{"name":"Updated Recipe","ingredients":["New Ingredient"],"steps":["New Step"]}`), http.StatusOK},
		{"Recipe DELETE", "DELETE", "/recipes/" + testRecipe.Id.String(), nil, http.StatusOK},
		{"Options request", "OPTIONS", "/recipes", nil, http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.path, bytes.NewBuffer(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			if tc.body != nil {
				req.Header.Set("Content-Type", "application/json")
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}
		})
	}
}
