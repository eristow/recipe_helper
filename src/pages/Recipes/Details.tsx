import H1 from "@/components/H1";
import H2 from "@/components/H2";
import PageContainer from "@/components/PageContainer";
import Recipe from "@/types/Recipe";
import { useLocation, useParams } from "react-router-dom";

interface LocationState {
  recipe: Recipe;
}

export function Details() {
  const { recipeId } = useParams();
  const { recipe } = useLocation().state as LocationState;

  if (!recipe) {
    // TODO: handle getting recipe by id from backend
  }

  console.log(recipeId);
  console.log(recipe);

  return (
    <PageContainer>
      <H1>Recipe Details</H1>
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
