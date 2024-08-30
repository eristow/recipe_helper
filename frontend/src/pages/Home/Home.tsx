import H1 from "@/components/H1/H1";
import PageContainer from "@/components/PageContainer/PageContainer";
import { Link } from "react-router-dom";

export default function Home() {
  return (
    <PageContainer>
      <Link to="/recipes">
        <H1>View Recipes</H1>
      </Link>
    </PageContainer>
  );
}
