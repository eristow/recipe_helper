import { useState, useEffect } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

type Recipe = {
  id: number;
  name: string;
  ingredients: string[];
  steps: string[];
};

function App() {
  const [count, setCount] = useState(0);
  const [recipes, setRecipes] = useState<Recipe[]>([]);

  const fetchRecipes = async () => {
    // TODO: extract URL to .env file
    // TODO: extract fetch to a service
    try {
      const response = await fetch("http://localhost:8080/recipes");
      const data: Recipe[] = await response.json();
      console.log(data);
      setRecipes(data);
    } catch (error) {
      console.error(`Failed to fetch recipes: ${error}`);
    }
  };

  useEffect(() => {
    fetchRecipes();
  }, []);

  return (
    <>
      <div className="flex justify-center">
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1 className="text-3xl font-bold underline">Vite + React</h1>
      <div>
        {recipes.map((recipe: Recipe) => (
          <div key={recipe.id}>
            <p>{recipe.id}</p>
            <p>{recipe.name}</p>
            <p>{recipe.ingredients.join(", ")}</p>
            <p>{recipe.steps.join(", ")}</p>
          </div>
        ))}
      </div>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
