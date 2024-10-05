package database

import (
	"reflect"
	"testing"

	"github.com/eristow/recipe_helper/internal/recipe"
)

var ds *Datastore
var r *recipe.Recipe

func setup(_ testing.TB) func(tb testing.TB) {
	ds = NewDatastore()
	r = recipe.NewRecipe("Chocolate Cake", []string{"Flour", "Sugar", "Cocoa Powder"}, []string{"Mix ingredients", "Bake at 350F for 30 min"})
	ds.AddRecipe(r)

	return func(tb testing.TB) {
		ds = nil
		r = nil
	}
}

func TestAddRecipe(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	if addedRecipe, exists := ds.m[r.Id.String()]; !exists {
		t.Errorf("AddRecipe() did not add the recipe")
	} else if !reflect.DeepEqual(addedRecipe, *r) {
		t.Errorf("AddRecipe() added incorrect recipe: got %v, want %v", addedRecipe, *r)
	}
}

func TestUpdateRecipe(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	updatedRecipe := &recipe.Recipe{
		Id:          r.Id,
		Name:        "Chocolate Cake",
		Ingredients: []string{"Flour", "Sugar", "Cocoa Powder", "Eggs"},
		Steps:       []string{"Mix ingredients", "Bake at 350F for 30 min", "Add eggs"},
	}
	ds.UpdateRecipe(r.Id.String(), updatedRecipe)

	if updated, exists := ds.m[r.Id.String()]; !exists {
		t.Errorf("UpdateRecipe() did not update the recipe")
	} else if !reflect.DeepEqual(updated, *updatedRecipe) {
		t.Errorf("UpdateRecipe() updated incorrect recipe: got %v, want %v", updated, *updatedRecipe)
	}
}

func TestGetRecipeByName(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	if got, exists := ds.GetRecipeByName("Chocolate Cake"); !exists {
		t.Errorf("GetRecipeByName() did not return the recipe")
	} else if !reflect.DeepEqual(got, r) {
		t.Errorf("GetRecipeByName() returned incorrect recipe: got %v, want %v", got, r)
	}
}

func TestGetRecipeById(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	if got, exists := ds.GetRecipeById(r.Id.String()); !exists {
		t.Errorf("GetRecipeById() did not return the recipe")
	} else if !reflect.DeepEqual(got, r) {
		t.Errorf("GetRecipeById() returned incorrect recipe: got %v, want %v", got, r)
	}
}

func TestListRecipes(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	recipes := ds.ListRecipes()

	if len(recipes) != 1 {
		t.Errorf("ListRecipes() returned incorrect number of recipes: got %d, want %d", len(recipes), 1)
	}
}

func TestDeleteRecipe(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ds.DeleteRecipe(r.Id.String())

	if _, exists := ds.m["Chocolate Cake"]; exists {
		t.Errorf("DeleteRecipe() did not delete the recipe")
	} else if len(ds.m) != 0 {
		t.Errorf("DeleteRecipe() did not delete the recipe")
	}
}
