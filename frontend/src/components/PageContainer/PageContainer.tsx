import { cn } from "@/utils/cn";
import { HTMLAttributes } from "react";

interface PageContainerProps extends HTMLAttributes<HTMLDivElement> {
  children: React.ReactNode;
  className?: string;
}

export default function PageContainer({
  children,
  className,
  ...props
}: PageContainerProps) {
  return (
    <div
      className={cn(
        "mx-auto mt-2 min-w-80 rounded-xl border-4 border-solid border-transparent bg-neutral-800 px-4 pb-4 shadow shadow-neutral-600",
        className,
      )}
      {...props}
    >
      {children}
    </div>
  );
}
