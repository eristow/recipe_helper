import H1 from "@/components/H1";
import H2 from "@/components/H2";
import PageContainer from "@/components/PageContainer";
import { Recipe } from "@/types/Recipe";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

export default function Recipes() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);

  useEffect(() => {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    // TODO: cache this?
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
    <PageContainer>
      <div className="grid grid-flow-row justify-center">
        <H1>Recipes</H1>
        <Link className="p-4 text-blue-700" to="/recipes/create">
          <H2>Create New Recipe</H2>
        </Link>
        {recipes.map((recipe: Recipe) => (
          <div
            className="mb-2 flex min-w-56 justify-center rounded-xl border-4 border-solid border-blue-800 p-4 align-middle"
            key={recipe.id}
          >
            <Link
              className="p-4"
              to={`/recipes/${recipe.id}`}
              state={{ recipe }}
            >
              <H2 className="m-auto">{recipe.name}</H2>
            </Link>
          </div>
        ))}
      </div>
    </PageContainer>
  );
}
