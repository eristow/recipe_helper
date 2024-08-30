import { Link } from "react-router-dom";
import H1 from "@/components/H1/H1";
import P from "@/components/P/P";
import { buttonClasses } from "@/components/Button/Button";

export default function Header() {
  return (
    <div className="text-central grid grid-flow-col grid-cols-[20%_60%_20%] border-b-4 border-solid border-neutral-950 bg-neutral-900 p-4">
      <div className="m-auto">
        <Link className={buttonClasses} to="/">
          <P>Home</P>
        </Link>
      </div>
      <div className="m-auto">
        <H1>Recipe Helper</H1>
      </div>
      <div className="m-auto">
        <Link className={buttonClasses} to="/recipes">
          <P>Recipes</P>
        </Link>
      </div>
    </div>
  );
}
