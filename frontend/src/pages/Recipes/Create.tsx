import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import PageContainer from "@/components/PageContainer/PageContainer";
import { RecipeNoId } from "@/types/Recipe";
import { cn } from "@/utils/cn";
import { useNavigate } from "react-router-dom";

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

const inputClassNames =
  "rounded-lg border border-solid border-neutral-800 bg-neutral-900 p-1";
const textAreaClassNames = cn(inputClassNames, "h-32");

export default function Create() {
  const navigate = useNavigate();

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    console.log(`Creating new recipe...`);

    const recipeName = event.currentTarget.recipeName.value;
    const recipeIngredients = event.currentTarget.recipeIngredients.value;
    const recipeSteps = event.currentTarget.recipeSteps.value;

    const ingredients = recipeIngredients.split("\n");
    const steps = recipeSteps.split("\n");

    const newRecipe: RecipeNoId = {
      name: recipeName,
      ingredients,
      steps,
    };
    console.log(newRecipe);

    createRecipe(newRecipe);

    navigate("/recipes");
  }

  async function createRecipe(newRecipe: RecipeNoId) {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    try {
      await fetch(`${backendUrl}/recipes/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newRecipe),
      });
    } catch (error) {
      console.error(`Failed to save recipe: ${error}`);
      return;
    }
  }

  return (
    <PageContainer>
      <H1>Create Recipe</H1>
      <form onSubmit={handleSubmit} data-testid="create-form">
        <div className="grid grid-flow-row gap-2">
          <InputContainer label="Recipe Name:">
            <input
              name="recipeName"
              className={inputClassNames}
              type="text"
              placeholder="Recipe Name"
            />
          </InputContainer>
          <InputContainer label="Recipe Ingredients:">
            <textarea
              name="recipeIngredients"
              className={textAreaClassNames}
              placeholder="Enter ingredients, one per line"
            />
          </InputContainer>
          <InputContainer label="Recipe Steps:">
            <textarea
              name="recipeSteps"
              className={textAreaClassNames}
              placeholder="Enter steps, one per line"
            />
          </InputContainer>
          <Button type="submit" data-testid="submit-button">
            Save
          </Button>
        </div>
      </form>
    </PageContainer>
  );
}
