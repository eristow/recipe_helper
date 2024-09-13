import { render, screen, waitFor } from "@testing-library/react";
import Ingredients from "./Ingredients";
import { beforeEach, describe, expect, test, vi } from "vitest";
import { userEvent } from "@testing-library/user-event";

describe("Ingredients Component", () => {
  beforeEach(() => {
    const mockRecipes = [
      {
        id: 1,
        name: "Recipe 1",
        ingredients: ["ingredient1"],
        steps: ["step1"],
      },
    ];

    vi.spyOn(globalThis, "fetch").mockResolvedValueOnce({
      json: async () => {
        // sleep for 1 second to simulate a slow response
        await new Promise((resolve) => setTimeout(resolve, 1000));
        return mockRecipes;
      },
    } as Response);
  });

  test("renders Ingredients component", () => {
    render(<Ingredients />);

    expect(screen.getByText("Ingredients")).toBeInTheDocument();
    expect(screen.getByTestId("ingredients-form")).toBeInTheDocument();
    expect(screen.getByTestId("ingredients-textarea")).toBeInTheDocument();
    expect(screen.getByTestId("submit-button")).toBeInTheDocument();
  });

  test("form submission works and sets pending state", async () => {
    const user = userEvent.setup();
    render(<Ingredients />);

    const textArea = screen.getByTestId("ingredients-textarea");
    const submitButton = screen.getByTestId("submit-button");

    await user.type(textArea, "ingredient1\ningredient2");
    await user.click(submitButton);

    await waitFor(() =>
      expect(screen.getByText("Generating Recipes...")).toBeInTheDocument(),
    );

    await waitFor(() =>
      expect(
        screen.queryByText("Generating Recipes..."),
      ).not.toBeInTheDocument(),
    );

    expect(screen.getByText("Suggested Recipes")).toBeInTheDocument();
  });

  test("generateRecipes function is called with correct ingredients list", async () => {
    const user = userEvent.setup();
    render(<Ingredients />);

    const textArea = screen.getByTestId("ingredients-textarea");
    const submitButton = screen.getByTestId("submit-button");

    await user.type(textArea, "ingredient1\ningredient2");
    await user.click(submitButton);

    await waitFor(() =>
      expect(globalThis.fetch).toHaveBeenCalledWith(
        expect.stringContaining("/recipes/generate"),
        expect.objectContaining({
          method: "POST",
          body: JSON.stringify({
            ingredientsList: ["ingredient1", "ingredient2"],
          }),
        }),
      ),
    );
  });

  test("displays recipes after fetching", async () => {
    const user = userEvent.setup();
    render(<Ingredients />);

    const textArea = screen.getByTestId("ingredients-textarea");
    const submitButton = screen.getByTestId("submit-button");

    await user.type(textArea, "ingredient1\ningredient2");
    await user.click(submitButton);

    await waitFor(() =>
      expect(screen.getByText("Generating Recipes...")).toBeInTheDocument(),
    );
    await waitFor(() =>
      expect(screen.getByText("Suggested Recipes")).toBeInTheDocument(),
    );
    expect(screen.getByText("Recipe 1")).toBeInTheDocument();
    expect(screen.getByText("* ingredient1")).toBeInTheDocument();
    expect(screen.getByText("1. step1")).toBeInTheDocument();
  });

  test("displays error message when fetch fails", async () => {
    vi.spyOn(globalThis, "fetch").mockRejectedValueOnce(
      new Error("Failed to fetch"),
    );

    const user = userEvent.setup();
    render(<Ingredients />);

    const submitButton = screen.getByTestId("submit-button");

    await user.click(submitButton);

    await waitFor(() =>
      expect(
        screen.getByText(
          "There was an error processing your request. Please try again.",
        ),
      ).toBeInTheDocument(),
    );
  });
});
