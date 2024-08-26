import H1 from "./H1";

export default function Header() {
  return (
    <div className="text-central grid grid-flow-col grid-cols-[20%_60%_20%] border-b-4 border-solid border-neutral-950 bg-neutral-900 p-4">
      <div className="m-auto">{/* TODO: add Home button */}</div>
      <div className="m-auto">
        <H1>Recipe Helper</H1>
      </div>
      <div className="m-auto"></div>
    </div>
  );
}
