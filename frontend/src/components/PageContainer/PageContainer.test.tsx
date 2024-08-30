import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import PageContainer from "./PageContainer";

describe("PageContainer", () => {
  it("should render children", () => {
    render(
      <PageContainer>
        <div>Test</div>
      </PageContainer>,
    );

    const testText = screen.getByText("Test");

    expect(testText).toBeInTheDocument();
  });
});
