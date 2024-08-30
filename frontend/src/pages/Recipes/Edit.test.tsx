import { screen, render } from "@testing-library/react";
import { MemoryRouter, useLocation, useParams } from "react-router-dom";
import { beforeEach, describe, expect, it, vi } from "vitest";
import Edit from "./Edit";

const recipe = {
  id: "1",
  name: "recipe-name",
  ingredients: ["ingredient-1", "ingredient-2"],
  steps: ["step-1", "step-2"],
};

describe("Edit", () => {
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

  it("should render the edit form", () => {
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
        <Edit />
      </MemoryRouter>,
    );

    const recipeName = screen.getByText("Edit recipe-name");

    expect(recipeName).toBeInTheDocument();
  });
});
