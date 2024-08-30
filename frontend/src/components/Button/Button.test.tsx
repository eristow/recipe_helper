import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Button from "./Button";

describe("Button Component", () => {
  it("should render the Button component with the correct text", () => {
    render(<Button>Click Me</Button>);

    const buttonElement = screen.getByText(/Click Me/i);

    expect(buttonElement).toBeInTheDocument();
  });

  it("should call the onClick handler when clicked", () => {
    const handleClick = vi.fn();

    render(<Button onClick={handleClick}>Click Me</Button>);

    const buttonElement = screen.getByText(/Click Me/i);
    fireEvent.click(buttonElement);

    expect(handleClick).toHaveBeenCalledTimes(1);
  });
});
