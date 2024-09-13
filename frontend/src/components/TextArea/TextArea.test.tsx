import { render, screen, fireEvent } from "@testing-library/react";
import TextArea from "./TextArea";
import { describe, it, expect } from "vitest";

describe("TextArea Component", () => {
  it("should render TextArea component", () => {
    render(<TextArea />);

    const textAreaElement = screen.getByRole("textbox");

    expect(textAreaElement).toBeInTheDocument();
  });

  it("should accept user input", () => {
    render(<TextArea />);

    const textAreaElement = screen.getByRole("textbox");
    fireEvent.change(textAreaElement, { target: { value: "test input" } });

    expect(screen.getByDisplayValue("test input")).toBeInTheDocument();
  });
});
