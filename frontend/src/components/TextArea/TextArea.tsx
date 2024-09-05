import { cn } from "@/utils/cn";

export default function TextArea(
  props: React.TextareaHTMLAttributes<HTMLTextAreaElement>,
) {
  return (
    <textarea
      className={cn(
        "rounded-lg border border-solid border-neutral-800 bg-neutral-900 p-1",
        "h-32",
      )}
      {...props}
    />
  );
}
