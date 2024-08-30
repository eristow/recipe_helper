import { describe, expect, it } from "vitest";
import Create from "./Create";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";

describe("Create", () => {
  it("should render the create form", () => {
    render(
      <MemoryRouter>
        <Create />
      </MemoryRouter>,
    );

    const createForm = screen.getByTestId("create-form");

    expect(createForm).toBeInTheDocument();
  });

  // // TODO: fix test
  // it("should call handleSubmit when the form is submitted", () => {
  //   render(
  //     <MemoryRouter>
  //       <Create />
  //     </MemoryRouter>,
  //   );

  //   const submitButton = screen.getByTestId("submit-button");
  //   const handleSubmitMock = vi.fn();

  //   vi.spyOn(Create.prototype, "handleSubmit").mockImplementation(
  //     handleSubmitMock,
  //   );

  //   fireEvent.click(submitButton);

  //   expect(handleSubmitMock).toHaveBeenCalled();
  // });

  // // TODO: fix test
  // it("should call createRecipe with the correct arguments", () => {
  //   render(
  //     <MemoryRouter>
  //       <Create />
  //     </MemoryRouter>,
  //   );

  //   const submitButton = screen.getByTestId("submit-button");
  //   const createRecipeMock = vi.fn();

  //   vi.spyOn(Create.prototype, "createRecipe").mockImplementation(
  //     createRecipeMock,
  //   );

  //   fireEvent.click(submitButton);

  //   expect(createRecipeMock).toHaveBeenCalled();
  // });
});
