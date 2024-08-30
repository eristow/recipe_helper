import { render, screen } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import ErrorPage from "./ErrorPage";
import {
  ErrorResponse,
  isRouteErrorResponse,
  useRouteError,
} from "react-router-dom";

const errorReturn: ErrorResponse = {
  statusText: "Not Found",
  status: 404,
  data: "Resource not found",
};

describe("ErrorPage", () => {
  beforeEach(() => {
    vi.mock("react-router-dom", async () => {
      const actual = await vi.importActual("react-router-dom");
      const mockUseRouteError = vi.fn();
      const mockIsRouteErrorResponse = vi.fn();

      return {
        ...actual,
        useRouteError: mockUseRouteError,
        isRouteErrorResponse: mockIsRouteErrorResponse,
      };
    });
  });

  it("should render error page", () => {
    render(<ErrorPage />);

    const oops = screen.getByText("Oops!");

    expect(oops).toBeInTheDocument();
  });

  it("should render error page with error message", () => {
    vi.mocked(useRouteError).mockReturnValue(errorReturn);
    vi.mocked(isRouteErrorResponse).mockReturnValue(true);

    render(<ErrorPage />);

    const errorMessage = screen.getByText(
      `${errorReturn.statusText} (${errorReturn.status}): ${errorReturn.data}`,
    );

    expect(errorMessage).toBeInTheDocument();
  });

  it("should render error page with unknown error message", () => {
    vi.mocked(useRouteError).mockReturnValue({});
    vi.mocked(isRouteErrorResponse).mockReturnValue(false);

    render(<ErrorPage />);

    const errorMessage = screen.getByText("Unknown error message");

    expect(errorMessage).toBeInTheDocument();
  });
});
