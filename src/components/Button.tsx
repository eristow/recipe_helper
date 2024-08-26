import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export default function Button(props: HTMLAttributes<HTMLButtonElement>) {
  const buttonClasses = "bg-primary-500 text-white px-4 py-2 rounded-lg";
  const combinedClasses = cn(buttonClasses, props.className);

  return <button className={combinedClasses}>{props.children}</button>;
}
