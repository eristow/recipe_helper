import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  className?: string;
  type?: "button" | "submit" | "reset";
}

export const buttonClasses =
  "rounded-lg border border-solid border-neutral-700 bg-neutral-600 px-4 py-2 text-white shadow-md";

export default function Button({
  children,
  className,
  type,
  ...props
}: ButtonProps) {
  const combinedClasses = cn(buttonClasses, className);

  return (
    <button className={combinedClasses} type={type} {...props}>
      {children}
    </button>
  );
}
