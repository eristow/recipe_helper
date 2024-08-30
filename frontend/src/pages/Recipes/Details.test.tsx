import { render, screen } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { Details } from "./Details";
import { MemoryRouter, useLocation, useParams } from "react-router-dom";

const recipe = {
  id: "1",
  name: "recipe-name",
  ingredients: ["ingredient-1", "ingredient-2"],
  steps: ["step-1", "step-2"],
};

describe("Details", () => {
  beforeEach(() => {
    vi.mock("react-router-dom", async () => {
      const actual = await vi.importActual("react-router-dom");
      const mockUseParams = vi.fn();
      const mockUseLocation = vi.fn();
      const mockUseNavigate = vi.fn();

      return {
        ...actual,
        useParams: mockUseParams,
        useLocation: mockUseLocation,
        useNavigate: mockUseNavigate,
      };
    });
  });

  it("should render the details page", () => {
    vi.mocked(useParams).mockReturnValue({ recipeId: "1" });
    vi.mocked(useLocation).mockReturnValue({
      state: {
        recipe: recipe,
      },
      key: "",
      pathname: "",
      search: "",
      hash: "",
    });

    render(
      <MemoryRouter>
        <Details />
      </MemoryRouter>,
    );

    const recipeName = screen.getByText("recipe-name");
    const recipeIngredients = screen.getByText("ingredient-1");

    expect(recipeName).toBeInTheDocument();
    expect(recipeIngredients).toBeInTheDocument();
  });
});
