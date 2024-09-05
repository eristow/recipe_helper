import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import H2 from "@/components/H2/H2";
import P from "@/components/P/P";
import PageContainer from "@/components/PageContainer/PageContainer";
import TextArea from "@/components/TextArea/TextArea";
import { Recipe } from "@/types/Recipe";
import { useState } from "react";

export default function Ingredients() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [pending, setPending] = useState<boolean>(false);

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    console.log(`Getting recipes based on ingredients...`);
    setPending(true);
    setRecipes([]);

    const ingredients = event.currentTarget.ingredients.value;

    const ingredientsList = ingredients.split("\n");
    console.log(ingredientsList);

    generateRecipes(ingredientsList);
  }

  async function generateRecipes(ingredientsList: string[]) {
    const backendUrl = import.meta.env.VITE_BACKEND_URL;

    try {
      const response = await fetch(`${backendUrl}/recipes/generate`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ingredientsList }),
      });

      const data: Recipe[] = (await response.json()) as Recipe[];
      console.log(data);

      setRecipes(data);
    } catch (error) {
      console.error(`Failed to generate recipes: ${error}`);
      return "";
    } finally {
      setPending(false);
    }
  }

  return (
    <>
      <PageContainer>
        <H1>Ingredients</H1>
        <form onSubmit={handleSubmit} data-testid="ingredients-form">
          <div className="grid grid-flow-row gap-2">
            <TextArea
              className={`h-96 w-96 ${pending ? "opacity-50" : ""}`}
              disabled={pending}
              name="ingredients"
              placeholder="Enter ingredients from your pantry or refrigerator, one per line"
            />
            <Button
              disabled={pending}
              type="submit"
              data-testid="submit-button"
            >
              Get Recipes
            </Button>
          </div>
        </form>
      </PageContainer>
      {(pending || recipes.length != 0) && (
        <PageContainer className="max-w-3xl">
          {pending && (
            <div>
              <H1 className="no-underline">Generating Recipes...</H1>
              <span className="relative mx-auto my-5 flex h-7 w-7">
                <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-blue-500 opacity-75"></span>
                <span className="relative inline-flex h-7 w-7 rounded-full bg-blue-600"></span>
              </span>
            </div>
          )}
          {/* TODO: figure out where the duplicate key is... */}
          {recipes.length != 0 && (
            <div>
              <H1 className="mb-6">Suggested Recipes</H1>
              {recipes.map((recipe: Recipe) => (
                <div className="mb-10" key={recipe.id}>
                  <H2 className="text-center">{recipe.name}</H2>
                  <div className="mb-5">
                    <H2>Ingredients:</H2>
                    {recipe.ingredients.map((ingredient, index) => (
                      <P key={recipe.id + ingredient + index}>* {ingredient}</P>
                    ))}
                  </div>
                  <div>
                    <H2>Steps:</H2>
                    {recipe.steps.map((step, index) => (
                      <>
                        <P
                          className="whitespace-pre-line"
                          key={recipe.id + step + index}
                        >
                          {index + 1}. {step}
                          {"\n"}
                        </P>
                      </>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          )}
        </PageContainer>
      )}
    </>
  );
}
