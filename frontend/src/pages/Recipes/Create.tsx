import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import Input from "@/components/Input/Input";
import PageContainer from "@/components/PageContainer/PageContainer";
import TextArea from "@/components/TextArea/TextArea";
import { RecipeNoId } from "@/types/Recipe";
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
            <Input name="recipeName" type="text" placeholder="Recipe Name" />
          </InputContainer>
          <InputContainer label="Recipe Ingredients:">
            <TextArea
              name="recipeIngredients"
              placeholder="Enter ingredients, one per line"
            />
          </InputContainer>
          <InputContainer label="Recipe Steps:">
            <TextArea
              name="recipeSteps"
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
