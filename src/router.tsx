import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Recipes from "./pages/Recipes";
import Layout from "./layout";
import ErrorPage from "./components/ErrorPage";
import Home from "./pages/Home";
import { Details } from "./pages/Recipes/Details";
import Edit from "./pages/Recipes/Edit";

const router = createBrowserRouter([
  {
    element: <Layout />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/",
        element: <Home />,
      },
      {
        path: "/recipes",
        element: <Recipes />,
      },
      {
        path: "/recipes/:recipeId",
        element: <Details />,
      },
      {
        path: "/recipes/edit/:recipeId",
        element: <Edit />,
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
}
