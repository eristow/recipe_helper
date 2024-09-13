package rest

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/eristow/recipe_helper_backend/internal/database"
	"github.com/eristow/recipe_helper_backend/internal/util"
	"github.com/ollama/ollama/api"
)

var (
	listRecipeRe     = regexp.MustCompile(`^\/recipes\/*$`)
	getRecipeRe      = regexp.MustCompile(`^\/recipes\/([a-zA-Z0-9\-]+)\/?$`)
	createRecipeRe   = regexp.MustCompile(`^\/recipes\/*$`)
	generateRecipeRe = regexp.MustCompile(`^\/recipes\/generate\/*$`)
)

func GetRecipeNameIdFromUrl(r *http.Request) string {
	matches := getRecipeRe.FindStringSubmatch(r.URL.Path)
	if matches == nil || len(matches) < 2 {
		return ""
	}
	return matches[1]
}

type OllamaClient interface {
	Generate(ctx context.Context, req *api.GenerateRequest, respFunc api.GenerateResponseFunc) error
}

type RootHandler struct{}
type RecipeHandler struct {
	store  *database.Datastore
	client OllamaClient
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func NewRecipeHandler(store *database.Datastore, client OllamaClient) *RecipeHandler {
	return &RecipeHandler{store: store, client: client}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: make this return a list of endpoints? or just use swagger/OpenAPI?
	util.EnableCors(w)
	log.Printf("Root: %s", r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Recipe Helper Backend!"))
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(w)
	log.Printf("Recipes router: %s", r.Method)
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listRecipeRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getRecipeRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createRecipeRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPut && getRecipeRe.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodDelete && getRecipeRe.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	case r.Method == http.MethodPost && generateRecipeRe.MatchString(r.URL.Path):
		h.Generate(w, r)
		return
	default:
		log.Println("Route not found")
		util.NotFound(w, r)
		return
	}
}
