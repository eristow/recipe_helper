import { cn } from "@/utils/cn";

export default function Input(
  props: React.InputHTMLAttributes<HTMLInputElement>,
) {
  return (
    <input
      className={cn(
        "rounded-lg border border-solid border-neutral-800 bg-neutral-900 p-1",
      )}
      {...props}
    />
  );
}
