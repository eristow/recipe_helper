package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/google/uuid"
	"github.com/ollama/ollama/api"
)

var (
	listRecipeRe     = regexp.MustCompile(`^\/recipes\/*$`)
	getRecipeRe      = regexp.MustCompile(`^\/recipes\/([a-zA-Z0-9\-]+)\/?$`)
	createRecipeRe   = regexp.MustCompile(`^\/recipes\/*$`)
	generateRecipeRe = regexp.MustCompile(`^\/recipes\/generate\/*$`)
)

type RootHandler struct{}
type RecipeHandler struct {
	store *database.Datastore
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
}

func HandleCors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(w)
			w.WriteHeader(http.StatusOK)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

func firstN(s string, n int) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func NewRecipeHandler(store *database.Datastore) *RecipeHandler {
	return &RecipeHandler{store: store}
}

func getRecipeNameIdFromUrl(r *http.Request) string {
	matches := getRecipeRe.FindStringSubmatch(r.URL.Path)
	if matches == nil || len(matches) < 2 {
		return ""
	}
	return matches[1]
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: make this return a list of endpoints? or just use swagger/OpenAPI?
	enableCors(w)
	log.Printf("Root: %s", r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Recipe Helper Backend!"))
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	log.Printf("Recipes router: %s", r.Method)
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listRecipeRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getRecipeRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createRecipeRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPut && getRecipeRe.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodDelete && getRecipeRe.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	case r.Method == http.MethodPost && generateRecipeRe.MatchString(r.URL.Path):
		h.Generate(w, r)
		return
	default:
		log.Println("Route not found")
		notFound(w, r)
		return
	}
}

func (h *RecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Println("List")
	recipes := h.store.ListRecipes()

	recipesJsonBytes, err := json.Marshal(recipes)
	if err != nil {
		log.Println("Error marshalling recipes:", err)
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipesJsonBytes)
}

func (h *RecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("Get")
	recipeId := getRecipeNameIdFromUrl(r)
	log.Printf("Recipe ID: %s", recipeId)
	if recipeId == "" {
		notFound(w, r)
		return
	}

	recipe, exists := h.store.GetRecipeById(recipeId)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("recipe not found"))
		return
	}

	recipeJsonBytes, err := json.Marshal(recipe)
	if err != nil {
		log.Println("Error marshalling recipe:", err)
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create")
	var newRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&newRecipe); err != nil {
		log.Println("Error decoding new recipe:", err)
		internalServerError(w, r)
		return
	}

	newRecipe.SetId(newRecipe.NewRecipeId())

	log.Printf("Adding new recipe: %+v", newRecipe)

	h.store.AddRecipe(&newRecipe)

	recipeJsonBytes, err := json.Marshal(newRecipe)
	if err != nil {
		log.Println("Error marshalling new recipe:", err)
		internalServerError(w, r)
		return
	}

	log.Printf("Added new recipe: %+v", newRecipe)

	w.WriteHeader(http.StatusCreated)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update")

	recipeId := getRecipeNameIdFromUrl(r)

	var updatedRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&updatedRecipe); err != nil {
		log.Println("Error decoding updated recipe:", err)
		internalServerError(w, r)
		return
	}

	id, err := uuid.Parse(recipeId)
	if err != nil {
		log.Println("Error parsing recipe ID:", err)
		internalServerError(w, r)
		return
	}
	updatedRecipe.SetId(id)

	log.Printf("Updating recipe: %+v", updatedRecipe)

	h.store.UpdateRecipe(recipeId, &updatedRecipe)

	recipeJsonBytes, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Println("Error marshalling updated recipe:", err)
		internalServerError(w, r)
		return
	}

	log.Printf("Updated recipe: %+v", updatedRecipe)

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}

func (h *RecipeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete")

	recipeId := getRecipeNameIdFromUrl(r)

	log.Printf("Deleting recipe: %s", recipeId)

	h.store.DeleteRecipe(recipeId)

	log.Printf("Deleted recipe: %s", recipeId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted recipe: " + recipeId))
}

func (h *RecipeHandler) Generate(w http.ResponseWriter, r *http.Request) {
	log.Println("Generate")
	log.Println("Request body:", r.Body)

	var ingredientsList recipe.Ingredients
	if err := json.NewDecoder(r.Body).Decode(&ingredientsList); err != nil {
		log.Println("Error decoding ingredients list:", err)
		internalServerError(w, r)
		return
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Println("Error creating client:", err)
		internalServerError(w, r)
		return
	}

	log.Println("Ingredients list:", ingredientsList)

	req := &api.GenerateRequest{
		Model:  "mistral",
		Prompt: fmt.Sprintf("Create 3 recipes given the following list of ingredients. The output should be formatted as an array of JSON objects like [{\"name\": \"\", \"ingredients\": [\"\", \"\"], \"steps\": [\"\", \"\"]}]. Do not number each recipe, and do not format like markdown. Ingredients List: %s", ingredientsList),

		Stream: newFalse(),
	}

	ctx := context.Background()
	recipeText := ""

	respFunc := func(resp api.GenerateResponse) error {
		log.Println("Done:", resp.Done)
		log.Println("DoneReason:", resp.DoneReason)
		if resp.Done && resp.DoneReason != "stop" {
			return errors.New(resp.DoneReason)
		}

		recipeText = resp.Response
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Println("Error generating recipe:", err)
		internalServerError(w, r)
		return
	}

	log.Println("Generated recipe text:", firstN(recipeText, 100))

	// Parse the recipeText into an array of Recipe objects
	var recipes []recipe.Recipe
	err = json.Unmarshal([]byte(recipeText), &recipes)
	if err != nil {
		log.Println("Error parsing recipe text:", err)
		internalServerError(w, r)
		return
	}

	for i := range recipes {
		recipes[i].SetId(recipes[i].NewRecipeId())
	}

	log.Println("Parsed recipes:", recipes)

	recipesBytes, err := json.Marshal(recipes)
	if err != nil {
		log.Println("Error marshalling recipe text:", err)
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipesBytes)
}

func internalServerError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 internal server error"))
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
