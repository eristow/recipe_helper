import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export default function H2(props: HTMLAttributes<HTMLHeadingElement>) {
  const h2Classes = "mb-2 text-2xl font-bold";
  const combinedClasses = cn(h2Classes, props.className);

  return <h2 className={combinedClasses}>{props.children}</h2>;
}
