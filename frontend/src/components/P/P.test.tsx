import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import P from "./P";

describe("P", () => {
  it('renders the text "Home"', () => {
    render(<P>Home</P>);

    const homeText = screen.getByText("Home");

    expect(homeText).toBeInTheDocument();
  });

  it("renders children", () => {
    render(
      <P>
        <div>Test</div>
      </P>,
    );

    expect(document.body.innerHTML).toMatch("<div>Test</div>");
  });
});
