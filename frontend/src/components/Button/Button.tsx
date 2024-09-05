import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  className?: string;
  type?: "button" | "submit" | "reset";
  color?: string;
}

export default function Button({
  children,
  className,
  type,
  color = "neutral",
  ...props
}: ButtonProps) {
  return (
    <button
      className={cn(
        `rounded-xl border-4 border-solid border-neutral-800 bg-${color}-600 px-4 py-2 text-white shadow shadow-neutral-700 transition duration-200 hover:border-${color}-900 hover:bg-${color}-700 justify-center align-middle`,
        className,
      )}
      type={type}
      {...props}
    >
      {children}
    </button>
  );
}
