import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import H1 from "./H1";

describe("H1", () => {
  it("should render H1 component with correct text", () => {
    render(<H1>Test Heading</H1>);

    const headingElement = screen.getByText("Test Heading");

    expect(headingElement).toBeInTheDocument();
  });
});
