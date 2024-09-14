package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/recipe"
	"github.com/eristow/recipe_helper_backend/internal/rest"
	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/jackc/pgx/v5"
	"github.com/ollama/ollama/api"
)

type slashFix struct {
	mux http.Handler
}

func (h *slashFix) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for strings.Contains(r.URL.Path, "//") {
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "//", "/")
	}
	h.mux.ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()
	rootH := rest.NewRootHandler()

	ds := database.NewDatastore()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Println("Error creating client:", err)
		return
	}
	recipeH := rest.NewRecipeHandler(ds, client)

	// Add test recipes
	pancakeRecipe := recipe.NewRecipe(
		"Pancakes",
		[]string{"Flour", "Eggs", "Milk", "Sugar"},
		[]string{"Mix ingredients", "Cook on pan"},
	)
	pizzaRecipe := recipe.NewRecipe(
		"Pizza",
		[]string{"Flour", "Yeast", "Sugar", "Olive Oil", "Salt", "Tomato Sauce", "Mozzarella Cheese", "Toppings"},
		[]string{"Mix warm water, yeast, sugar. Cover and allow to rest for 5 min.", "Add flour, salt, olive oil. Mix until dough forms.", "Knead dough for 5 min.", "Cover and allow to rise for 1 hour.", "Preheat oven to 475F.", "Roll out dough.", "Add sauce, cheese, toppings.", "Bake for 13-15 min."},
	)

	ds.AddRecipe(pancakeRecipe)
	ds.AddRecipe(pizzaRecipe)

	// Add test entry into DB
	// TODO: get DB URL from env
	// databaseUrl := os.Getenv("DATABASE_URL")
	databaseUrl := "postgres://defaultuser:defaultpassword@database:5432/defaultdatabase"
	log.Println("Database URL: ", databaseUrl)
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS recipes (id uuid DEFAULT gen_random_uuid(), name TEXT, ingredients TEXT[], steps TEXT[], PRIMARY KEY (id))")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO recipes (name, ingredients, steps) VALUES ($1, $2, $3)", "Pancakes", []string{"Flour", "Eggs", "Milk", "Sugar"}, []string{"Mix ingredients", "Cook on pan"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}

	// Test getting from DB
	rows, err := conn.Query(context.Background(), "SELECT name, ingredients, steps FROM recipes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get recipes from table: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var ingredients []string
		var steps []string
		err = rows.Scan(&name, &ingredients, &steps)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse recipe: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(name, ingredients, steps)
	}

	mux.Handle("/", util.HandleCors(rootH))
	mux.Handle("/recipes/", util.HandleCors(recipeH))
	mux.Handle("/recipes", util.HandleCors(recipeH))
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &slashFix{mux}))
}
