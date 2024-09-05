import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import H2 from "@/components/H2/H2";
import P from "@/components/P/P";
import PageContainer from "@/components/PageContainer/PageContainer";
import { Recipe } from "@/types/Recipe";
import { useEffect, useState } from "react";
import { Link, useLocation, useNavigate, useParams } from "react-router-dom";

interface LocationState {
  recipe: Recipe;
}

export function Details() {
  const { recipeId } = useParams();
  const locationState = useLocation().state as LocationState;
  const navigate = useNavigate();
  const [recipe, setRecipe] = useState<Recipe>(locationState.recipe);

  useEffect(() => {
    // TODO: handle getting recipe by id from backend
    async function getRecipeById(id: string) {
      console.log(`Getting recipe with id: ${id}`);
      setRecipe({} as Recipe);
    }

    if (recipeId && !recipe) {
      getRecipeById(recipeId);
    }
  }, []);

  async function deleteRecipe(id: string) {
    console.log(`Deleting recipe with id: ${id}`);
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    try {
      await fetch(`${backendUrl}/recipes/${id}`, {
        method: "DELETE",
      });
    } catch (error) {
      console.error(`Failed to delete recipe: ${error}`);
      return;
    }

    navigate("/recipes");
  }

  return (
    <PageContainer>
      <H1>{`${recipe.name}`}</H1>
      <div className="flex justify-between">
        <Button>
          <Link to={`/recipes/edit/${recipe.id}`} state={{ recipe }}>
            <P>Edit</P>
          </Link>
        </Button>
        <Button onClick={() => deleteRecipe(recipe.id)}>
          <P>Delete</P>
        </Button>
      </div>
      <div className="mb-2 list-disc">
        <H2>Ingredients:</H2>
        {recipe.ingredients.map((ingredient: string) => (
          <li key={ingredient}>{ingredient}</li>
        ))}
      </div>
      <div className="mb-2 list-decimal">
        <H2>Steps:</H2>
        {recipe.steps.map((step: string) => (
          <li key={step}>{step}</li>
        ))}
      </div>
    </PageContainer>
  );
}
