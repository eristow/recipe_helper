import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import P from "./P";

describe("P", () => {
  it('renders the text "Home"', () => {
    render(<P>Home</P>);

    const homeText = screen.getByText("Home");

    expect(homeText).toBeInTheDocument();
  });
});
