import { useState, useEffect } from "react";
import H1 from "./components/H1";
import P from "./components/P";

type Recipe = {
  id: number;
  name: string;
  ingredients: string[];
  steps: string[];
};

function App() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);

  useEffect(() => {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    async function fetchRecipes() {
      // TODO: extract fetch to a service?
      try {
        const response = await fetch(`${backendUrl}/recipes`);
        const data: Recipe[] = await response.json();
        console.log(data);
        setRecipes(data);
      } catch (error) {
        console.error(`Failed to fetch recipes: ${error}`);
      }
    }

    fetchRecipes();
  }, []);

  return (
    <div className="flex h-screen w-screen flex-col justify-start bg-neutral-700 align-top">
      <div className="mx-auto mt-2 rounded-xl border-4 border-solid border-transparent bg-neutral-800 p-4">
        <H1>Recipe Helper</H1>
        <div>
          {recipes.map((recipe: Recipe) => (
            <div key={recipe.id}>
              <P>{recipe.id}</P>
              <P>{recipe.name}</P>
              <P>{recipe.ingredients.join(", ")}</P>
              <P>{recipe.steps.join(", ")}</P>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
