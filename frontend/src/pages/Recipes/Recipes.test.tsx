import { render, screen, waitFor } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import Recipes from "./Recipes";
import { describe, expect, it, vi } from "vitest";
import { Recipe } from "@/types/Recipe";

describe("Recipes", () => {
  it("should render page title", () => {
    render(
      <BrowserRouter>
        <Recipes />
      </BrowserRouter>,
    );

    const pageTitle = screen.getByText("Recipes");
    expect(pageTitle).toBeInTheDocument();
  });

  it("should render recipe list", async () => {
    const mockRecipes: Recipe[] = [
      {
        id: "1",
        name: "Pasta",
        ingredients: ["pasta", "sauce"],
        steps: ["cook on stove", "serve"],
      },
      {
        id: "2",
        name: "Pizza",
        ingredients: ["dough", "cheese", "toppings"],
        steps: ["cook in oven", "serve"],
      },
    ];

    vi.spyOn(globalThis, "fetch").mockResolvedValueOnce({
      json: async () => mockRecipes,
    } as Response);

    render(
      <BrowserRouter>
        <Recipes />
      </BrowserRouter>,
    );

    await waitFor(() => {
      expect(screen.getByText("Pasta")).toBeInTheDocument();
      expect(screen.getByText("Pizza")).toBeInTheDocument();
    });

    const recipeList = screen.getAllByTestId("recipe-list");

    expect(recipeList).toHaveLength(2);
  });
});
