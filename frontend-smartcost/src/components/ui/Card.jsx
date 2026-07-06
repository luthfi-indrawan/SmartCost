import { cn } from "../../lib/utils";

export function Card({
  children,
  className,
  padding = "normal",
  hover = false,
  onClick,
}) {
  const paddingMap = {
    none: "",
    sm: "p-3",
    normal: "p-4",
    lg: "p-5",
  };

  return (
    <div
      onClick={onClick}
      className={cn(
        "bg-wa-surface rounded-[14px] border border-neutral-200",
        "shadow-[0_1px_3px_rgba(0,0,0,0.05)]",
        paddingMap[padding],
        hover &&
          "cursor-pointer hover:border-primary-500 hover:bg-primary-50 hover:shadow-[0_4px_12px_rgba(76,175,80,0.15)] transition-all duration-200",
        onClick && "cursor-pointer active:scale-[0.98]",
        className,
      )}
    >
      {children}
    </div>
  );
}
