### Recipe Helper Backend

This app is a recipe helper that allows users to add, view, and delete recipes. It is a Go application that uses the `net/http` package to create a web server. The app currently uses an in-memory "database" to store recipes.

## How to run the app

* docker build -t recipe_helper_backend .
* docker run -p 8080:8080 --name recipe_helper_backend recipe_helper_backend
  * Or run `./run.sh`

## TODO:

- [ ] Switch from storing recipes in in-memory to storing them in a database.

- [ ] Create swagger docs

## Done:
- [x] Write tests for main.go
- [x] Write tests for rest.go
- [x] Write tests for database.go
- [x] Write tests for recipes.go
- [x] Create PUT endpoint for updating recipes
- [x] Extract HTML into React front-end
- [x] Add logging
- [x] Create DELETE endpoint for deleting recipes
- [x] Return all recipes sorted by name
- [x] ID should be assigned and not 0000...
- [x] Get by ID instead of by name
- [x] Delete by ID instead of by name
