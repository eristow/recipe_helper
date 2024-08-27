import { HTMLAttributes } from "react";

export default function PageContainer(props: HTMLAttributes<HTMLDivElement>) {
  return (
    <div className="mx-auto mt-2 min-w-80 rounded-xl border-4 border-solid border-transparent bg-neutral-800 px-4 pb-4">
      {props.children}
    </div>
  );
}
