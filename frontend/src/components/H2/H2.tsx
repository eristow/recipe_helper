import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export default function H2(props: HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h2 className={cn("mb-2 text-2xl font-bold", props.className)}>
      {props.children}
    </h2>
  );
}
