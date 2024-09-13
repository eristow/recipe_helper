package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/util"
)

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create")
	var newRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&newRecipe); err != nil {
		log.Println("Error decoding new recipe:", err)
		util.InternalServerError(w, r)
		return
	}

	newRecipe.SetId(newRecipe.NewRecipeId())

	log.Printf("Adding new recipe: %+v", newRecipe)

	h.store.AddRecipe(&newRecipe)

	recipeJsonBytes, err := json.Marshal(newRecipe)
	if err != nil {
		log.Println("Error marshalling new recipe:", err)
		util.InternalServerError(w, r)
		return
	}

	log.Printf("Added new recipe: %+v", newRecipe)

	w.WriteHeader(http.StatusCreated)
	w.Write(recipeJsonBytes)
}
