import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

interface H1Props extends HTMLAttributes<HTMLHeadingElement> {
  children: React.ReactNode;
  className?: string;
}

export default function H1({ children, className, ...props }: H1Props) {
  return (
    <h1
      className={cn("m-2 text-center text-3xl font-bold underline", className)}
      {...props}
    >
      {children}
    </h1>
  );
}
