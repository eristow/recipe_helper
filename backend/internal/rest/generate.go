package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/ollama/ollama/api"
)

func (h *RecipeHandler) Generate(w http.ResponseWriter, r *http.Request) {
	log.Println("Generate")
	log.Println("Request body:", r.Body)

	var ingredientsList recipe.Ingredients
	if err := json.NewDecoder(r.Body).Decode(&ingredientsList); err != nil {
		log.Println("Error decoding ingredients list:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Println("Ingredients list:", ingredientsList)

	req := &api.GenerateRequest{
		Model:  "mistral",
		Prompt: fmt.Sprintf("Create 3 recipes given the following list of ingredients. The output should be formatted as an array of JSON objects like [{\"name\": \"\", \"ingredients\": [\"\", \"\"], \"steps\": [\"\", \"\"]}]. Do not number each recipe, and do not format like markdown. Ingredients List: %s", ingredientsList),

		Stream: util.NewFalse(),
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

	err := h.client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Println("Error generating recipe:", err)
		util.InternalServerError(w, r)
		return
	}

	log.Println("Generated recipe text:", util.FirstN(recipeText, 100))

	// Parse the recipeText into an array of Recipe objects
	var recipes []recipe.Recipe
	err = json.Unmarshal([]byte(recipeText), &recipes)
	if err != nil {
		log.Println("Error parsing recipe text:", err)
		util.InternalServerError(w, r)
		return
	}

	for i := range recipes {
		recipes[i].SetId(recipes[i].NewRecipeId())
	}

	log.Println("Parsed recipes:", recipes)

	recipesBytes, err := json.Marshal(recipes)
	if err != nil {
		log.Println("Error marshalling recipe text:", err)
		util.InternalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(recipesBytes)
}
