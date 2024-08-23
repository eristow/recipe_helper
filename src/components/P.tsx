import { HTMLAttributes } from "react";

export default function p(props: HTMLAttributes<HTMLHeadingElement>) {
  return (
    <p className="text-lg font-medium text-neutral-200">{props.children}</p>
  );
}
