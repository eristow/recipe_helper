import { Link } from "react-router-dom";
import H1 from "@/components/H1/H1";
import P from "@/components/P/P";
import Button from "@/components/Button/Button";

export default function Header() {
  return (
    <div className="text-central border-b-1 grid grid-flow-col grid-cols-[20%_60%_20%] border-solid border-neutral-500 bg-neutral-900 p-4 shadow shadow-neutral-600">
      <Button className="m-auto">
        <Link to="/">
          <P>Home</P>
        </Link>
      </Button>
      <div className="m-auto">
        <H1>Recipe Helper</H1>
      </div>
      <Button className="m-auto">
        <Link to="/recipes">
          <P>Recipes</P>
        </Link>
      </Button>
    </div>
  );
}
