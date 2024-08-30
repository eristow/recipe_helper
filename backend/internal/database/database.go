package database

import (
	"sort"
	"sync"

	"github.com/eristow/recipe_helper_backend/internal/recipe"
)

// In-memory datastore
type Datastore struct {
	m  map[string]recipe.Recipe
	mu *sync.RWMutex
}

func NewDatastore() *Datastore {
	return &Datastore{
		m:  make(map[string]recipe.Recipe),
		mu: &sync.RWMutex{},
	}
}

func (ds *Datastore) AddRecipe(r *recipe.Recipe) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	key := r.Id.String()
	ds.m[key] = *r
}

func (ds *Datastore) UpdateRecipe(key string, r *recipe.Recipe) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.m[key] = *r
}

func (ds *Datastore) GetRecipeByName(key string) (*recipe.Recipe, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var recipe recipe.Recipe
	var exists bool

	for _, r := range ds.m {
		if r.Name == key {
			recipe = r
			exists = true
			break
		}
	}

	return &recipe, exists
}

func (ds *Datastore) GetRecipeById(key string) (*recipe.Recipe, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	recipe, exists := ds.m[key]

	return &recipe, exists
}

func (ds *Datastore) ListRecipes() (recipes []recipe.Recipe) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, r := range ds.m {
		recipes = append(recipes, r)
	}

	sort.Slice(recipes, func(i, j int) bool {
		return recipes[i].Name < recipes[j].Name
	})

	return
}

func (ds *Datastore) DeleteRecipe(key string) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	delete(ds.m, key)
}
