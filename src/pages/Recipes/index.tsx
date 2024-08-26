import Button from "@/components/Button";
import H1 from "@/components/H1";
import H2 from "@/components/H2";
import P from "@/components/P";
import PageContainer from "@/components/PageContainer";
import Recipe from "@/types/Recipe";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

export default function Recipes() {
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

  async function deleteRecipe(id: string) {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    try {
      await fetch(`${backendUrl}/recipes/${id}`, {
        method: "DELETE",
      });
      setRecipes((prevRecipes) =>
        prevRecipes.filter((recipe) => recipe.id !== id),
      );
    } catch (error) {
      console.error(`Failed to delete recipe: ${error}`);
    }
  }

  return (
    <PageContainer>
      <div className="grid grid-flow-row">
        <H1>Recipes</H1>
        {recipes.map((recipe: Recipe) => (
          // TODO: the text in the divs should be centered
          <div
            className="mb-2 flex justify-between rounded-xl border-4 border-solid border-blue-800 p-4"
            key={recipe.id}
          >
            <Link
              className="p-4"
              to={`/recipes/${recipe.id}`}
              state={{ recipe }}
            >
              <H2 className="m-auto">{recipe.name}</H2>
            </Link>
            <Link className="p-4" to={`/recipes/edit/${recipe.id}`}>
              <P>Edit</P>
            </Link>
            <Button onClick={() => deleteRecipe(recipe.id)}>
              <P>Delete</P>
            </Button>
          </div>
        ))}
      </div>
    </PageContainer>
  );
}
