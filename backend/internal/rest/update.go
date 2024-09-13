package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/google/uuid"
)

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update")

	recipeId := GetRecipeNameIdFromUrl(r)

	var updatedRecipe recipe.Recipe
	if err := json.NewDecoder(r.Body).Decode(&updatedRecipe); err != nil {
		log.Println("Error decoding updated recipe:", err)
		util.InternalServerError(w, r)
		return
	}

	id, err := uuid.Parse(recipeId)
	if err != nil {
		log.Println("Error parsing recipe ID:", err)
		util.InternalServerError(w, r)
		return
	}
	updatedRecipe.SetId(id)

	log.Printf("Updating recipe: %+v", updatedRecipe)

	h.store.UpdateRecipe(recipeId, &updatedRecipe)

	recipeJsonBytes, err := json.Marshal(updatedRecipe)
	if err != nil {
		log.Println("Error marshalling updated recipe:", err)
		util.InternalServerError(w, r)
		return
	}

	log.Printf("Updated recipe: %+v", updatedRecipe)

	w.WriteHeader(http.StatusOK)
	w.Write(recipeJsonBytes)
}
