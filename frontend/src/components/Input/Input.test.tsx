import { render, screen, fireEvent } from "@testing-library/react";
import Input from "./Input";
import { describe, it, expect } from "vitest";

describe("Input Component", () => {
  it("should render Input component", () => {
    render(<Input />);

    const inputElement = screen.getByRole("textbox");

    expect(inputElement).toBeInTheDocument();
  });

  it("should accept user input", () => {
    render(<Input />);

    const inputElement = screen.getByRole("textbox");
    fireEvent.change(inputElement, { target: { value: "test input" } });

    expect(screen.getByDisplayValue("test input")).toBeInTheDocument();
  });
});
