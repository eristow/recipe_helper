import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

export const buttonClasses =
  "rounded-lg border border-solid border-neutral-700 bg-neutral-600 px-4 py-2 text-white shadow-md";

export default function Button({
  children,
  className,
  ...props
}: HTMLAttributes<HTMLButtonElement>) {
  const combinedClasses = cn(buttonClasses, className);

  return (
    <button className={combinedClasses} {...props}>
      {children}
    </button>
  );
}
