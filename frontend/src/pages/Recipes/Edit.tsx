import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import Input from "@/components/Input/Input";
import PageContainer from "@/components/PageContainer/PageContainer";
import TextArea from "@/components/TextArea/TextArea";
import { Recipe } from "@/types/Recipe";
import { useLocation, useNavigate, useParams } from "react-router-dom";

// TODO: move this to a shared location for Edit and Details?
interface LocationState {
  recipe: Recipe;
}

function InputContainer({
  label,
  children,
}: {
  label: string;
  children: React.ReactNode;
}) {
  return (
    <div className="grid grid-flow-row">
      <label>{label}</label>
      {children}
    </div>
  );
}

export default function Edit() {
  const { recipeId } = useParams();
  const state = useLocation().state as LocationState;
  const navigate = useNavigate();
  let recipe: Recipe = { id: "0", name: "", ingredients: [], steps: [] };
  console.log(state);

  if (!state?.recipe) {
    // TODO: handle getting recipe by id from backend
    console.log(`Getting recipe with id: ${recipeId}`);
  } else {
    recipe = state.recipe;
  }

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    console.log(`Saving recipe with id: ${recipeId}`);

    const recipeName = event.currentTarget.recipeName.value;
    const recipeIngredients = event.currentTarget.recipeIngredients.value;
    const recipeSteps = event.currentTarget.recipeSteps.value;

    const ingredients = recipeIngredients.split("\n");
    const steps = recipeSteps.split("\n");

    const updatedRecipe = {
      ...recipe,
      name: recipeName,
      ingredients,
      steps,
    };
    console.log(updatedRecipe);

    saveRecipe(updatedRecipe);

    navigate("/recipes");
  }

  async function saveRecipe(updatedRecipe: Recipe) {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    try {
      await fetch(`${backendUrl}/recipes/${recipeId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedRecipe),
      });
    } catch (error) {
      console.error(`Failed to save recipe: ${error}`);
      return;
    }
  }

  return (
    <PageContainer>
      <H1>Edit {recipe.name}</H1>
      <form onSubmit={handleSubmit}>
        <div className="grid grid-flow-row gap-2">
          <InputContainer label="Recipe Name:">
            <Input name="recipeName" type="text" defaultValue={recipe.name} />
          </InputContainer>
          <InputContainer label="Recipe Ingredients:">
            <TextArea
              name="recipeIngredients"
              defaultValue={recipe.ingredients.join("\n")}
            />
          </InputContainer>
          <InputContainer label="Recipe Steps:">
            <TextArea
              name="recipeSteps"
              defaultValue={recipe.steps.join("\n")}
            />
          </InputContainer>
          <Button type="submit">Save</Button>
        </div>
      </form>
    </PageContainer>
  );
}
