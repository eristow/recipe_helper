package recipe_test

import (
	"reflect"
	"testing"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

var r *recipe.Recipe

func setup(_ testing.TB) func(tb testing.TB) {
	r = recipe.NewRecipe("Chocolate Cake", []string{"Flour", "Sugar", "Cocoa Powder"}, []string{"Mix ingredients", "Bake at 350F for 30 min"})

	return func(tb testing.TB) {
		r = nil
	}
}

func TestNewRecipeId(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	newUuid := r.NewRecipeId()

	if reflect.TypeOf(newUuid).String() != "uuid.UUID" {
		t.Errorf("Expected type uuid.UUID, got %v", reflect.TypeOf(newUuid))
	}
}
