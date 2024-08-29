import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Home from "./Home";
import { MemoryRouter } from "react-router-dom";

describe("Home Page", () => {
  it("should render the Home page with the correct heading", () => {
    render(
      <MemoryRouter>
        <Home />
      </MemoryRouter>,
    );

    const headingText = screen.getByText("View Recipes");

    expect(headingText).toBeInTheDocument();
  });
});
