package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/util"
)

func (h *RecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("Get")
	recipeId := GetRecipeNameIdFromUrl(r)
	log.Printf("Recipe ID: %s", recipeId)
	if recipeId == "" {
		util.NotFound(w, r)
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
		util.InternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}
