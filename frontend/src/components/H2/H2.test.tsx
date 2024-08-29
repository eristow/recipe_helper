import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import H2 from "./H2";

describe("H2", () => {
  it("should render H2 component with correct text", () => {
    render(<H2>Test Heading</H2>);

    const headingElement = screen.getByText("Test Heading");

    expect(headingElement).toBeInTheDocument();
  });
});
