import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
  className?: string;
  type?: "button" | "submit" | "reset";
  color?: string;
  disabled?: boolean;
}

export default function Button({
  children,
  className,
  type,
  color = "neutral",
  disabled,
  ...props
}: ButtonProps) {
  const colorToClass: { [key: string]: string } = {
    neutral:
      "bg-neutral-600 hover:bg-neutral-700 border-neutral-800 hover:border-neutral-900",
    red: "bg-red-600 hover:bg-red-700 border-red-800 hover:border-red-900",
    green:
      "bg-green-600 hover:bg-green-700 border-green-800 hover:border-green-900",
    blue: "bg-blue-600 hover:bg-blue-700 border-blue-800 hover:border-blue-900",
    yellow:
      "bg-yellow-600 hover:bg-yellow-700 border-yellow-800 hover:border-yellow-900",
  };

  return (
    <button
      className={cn(
        `justify-center rounded-xl border-4 border-solid px-4 py-2 align-middle shadow shadow-neutral-700 transition duration-200 ${colorToClass[color]}`,
        className,
      )}
      type={type}
      disabled={disabled}
      {...props}
    >
      {children}
    </button>
  );
}
