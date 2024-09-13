package rest

import (
	"log"
	"net/http"
)

func (h *RecipeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete")

	recipeId := GetRecipeNameIdFromUrl(r)

	log.Printf("Deleting recipe: %s", recipeId)

	h.store.DeleteRecipe(recipeId)

	log.Printf("Deleted recipe: %s", recipeId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted recipe: " + recipeId))
}
