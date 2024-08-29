import { HTMLAttributes } from "react";

export default function H1(props: HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h1 className="m-2 text-center text-3xl font-bold underline">
      {props.children}
    </h1>
  );
}
