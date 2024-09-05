import Button from "@/components/Button/Button";
import H1 from "@/components/H1/H1";
import PageContainer from "@/components/PageContainer/PageContainer";
import { Link } from "react-router-dom";

export default function Home() {
  return (
    <PageContainer className="grid gap-4 p-4">
      <Button>
        <Link to="/recipes">
          <H1>View Recipes</H1>
        </Link>
      </Button>
      <Button>
        <Link to="/ingredients">
          <H1>Input Ingredients</H1>
        </Link>
      </Button>
    </PageContainer>
  );
}
