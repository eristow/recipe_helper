# Recipe Helper
This app is a recipe helper that allows users to add, view, and delete recipes. It will also allow users to input ingredients from their fridge/pantry and suggest recipes based on those ingredients.

## Front-End
The front-end is a React application created with Vite and TypeScript. Styling is done with Tailwind CSS.

## Back-End
The back-end is a Go application that uses the `net/http` package to create a web server. The app currently uses an in-memory "database" to store recipes.

## Running the App
* `dc up`
* To run the FE and BE containers outside of docker-compose:
	* `./frontend/run.sh`
	* `./backend/run.sh`
	* Look at README.md in each directory for more information.
