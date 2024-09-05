import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import H2 from "@/components/H2/H2";
import PageContainer from "@/components/PageContainer/PageContainer";
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
        <Button className="p-4" color="green">
          <Link to="/recipes/create">
            <H2 className="m-auto">Create New Recipe</H2>
          </Link>
        </Button>
        <H1>Recipes</H1>
        <div className="grid grid-cols-2 gap-2 sm:grid-cols-3 md:grid-cols-4">
          {recipes.map((recipe: Recipe) => (
            <Button
              className="py-6"
              color="blue"
              key={recipe.id}
              data-testid="recipe-list"
            >
              <Link to={`/recipes/${recipe.id}`} state={{ recipe }}>
                <H2 className="m-auto">{recipe.name}</H2>
              </Link>
            </Button>
          ))}
        </div>
      </div>
    </PageContainer>
  );
}
