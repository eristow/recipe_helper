import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import Header from "./Header";
import { expect, describe, it } from "vitest";

describe("Header", () => {
  it("should render header component", () => {
    render(
      <BrowserRouter>
        <Header />
      </BrowserRouter>,
    );

    const homeLink = screen.getByText("Home");
    const title = screen.getByText("Recipe Helper");
    const recipeLink = screen.getByText("Recipes");

    expect(homeLink).toBeInTheDocument();
    expect(title).toBeInTheDocument();
    expect(recipeLink).toBeInTheDocument();
  });
});
