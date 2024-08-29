import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Layout from "./layout";
import { MemoryRouter } from "react-router-dom";

describe("Layout Component", () => {
  it("should render the Layout component with the correct content", () => {
    render(
      <MemoryRouter>
        <Layout />
      </MemoryRouter>,
    );

    const contentElement = screen.getByText("Home");

    expect(contentElement).toBeInTheDocument();
  });
});
