export type Recipe = {
  id: string;
  name: string;
  ingredients: string[];
  steps: string[];
};

export type RecipeNoId = Omit<Recipe, "id">;
