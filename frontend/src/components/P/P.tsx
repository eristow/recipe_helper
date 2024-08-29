import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export default function p(props: HTMLAttributes<HTMLHeadingElement>) {
  const pClasses = "text-lg font-medium";
  const combinedClasses = cn(pClasses, props.className);

  return <p className={combinedClasses}>{props.children}</p>;
}
