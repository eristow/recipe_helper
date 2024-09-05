import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export default function p(props: HTMLAttributes<HTMLHeadingElement>) {
  return (
    <p className={cn("text-lg font-medium", props.className)}>
      {props.children}
    </p>
  );
}
