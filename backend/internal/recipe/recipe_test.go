package recipe

import (
	"reflect"
	"testing"
)

var r *Recipe

func setup(_ testing.TB) func(tb testing.TB) {
	r = NewRecipe("Chocolate Cake", []string{"Flour", "Sugar", "Cocoa Powder"}, []string{"Mix ingredients", "Bake at 350F for 30 min"})

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
