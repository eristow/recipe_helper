import { Outlet } from "react-router-dom";
import Header from "./components/Header";

export default function Layout() {
  return (
    <div className="flex min-h-screen flex-col justify-start bg-neutral-900 align-top text-neutral-100">
      <Header />
      <Outlet />
    </div>
  );
}
