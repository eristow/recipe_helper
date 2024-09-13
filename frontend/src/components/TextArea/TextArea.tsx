import { cn } from "@/utils/cn";

interface TextAreaProps
  extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
  className?: string;
}

export default function TextArea({ className, ...props }: TextAreaProps) {
  return (
    <textarea
      className={cn(
        "h-32 rounded-lg border border-solid border-neutral-800 bg-neutral-900 p-1",
        className,
      )}
      {...props}
    />
  );
}
