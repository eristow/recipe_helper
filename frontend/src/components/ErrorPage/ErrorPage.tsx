import { isRouteErrorResponse, useRouteError } from "react-router-dom";

export default function ErrorPage() {
  const error = useRouteError();
  console.error(error);

  return (
    <div className="grid h-screen w-screen justify-center bg-neutral-100 text-neutral-800 dark:bg-neutral-800 dark:text-neutral-100">
      <div className="m-auto">
        <h1 className="mb-3">Oops!</h1>
        <p className="mb-3">Sorry, an unexpected error has occurred.</p>
        <p>
          <i>
            {isRouteErrorResponse(error)
              ? `${error.statusText} (${error.status}): ${error.data}`
              : "Unknown error message"}
          </i>
        </p>
      </div>
    </div>
  );
}
