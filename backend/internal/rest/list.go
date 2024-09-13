package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/util"
)

func (h *RecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Println("List")
	recipes := h.store.ListRecipes()

	recipesJsonBytes, err := json.Marshal(recipes)
	if err != nil {
		log.Println("Error marshalling recipes:", err)
		util.InternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipesJsonBytes)
}
