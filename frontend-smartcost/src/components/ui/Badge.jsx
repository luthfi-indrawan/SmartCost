import { cn } from "../../lib/utils";

const variants = {
  primary: "bg-primary-50 text-primary-800",
  secondary: "bg-secondary-50 text-secondary-800",
  success: "bg-success-50 text-success-700",
  danger: "bg-danger-50 text-danger-700",
  warning: "bg-warning-50 text-warning-700",
  info: "bg-info-50 text-info-700",
  neutral: "bg-neutral-100 text-neutral-600",
};

export function Badge({
  children,
  variant = "neutral",
  className,
  size = "sm",
}) {
  return (
    <span
      className={cn(
        "inline-flex items-center font-medium rounded-full",
        size === "sm" && "px-2 py-0.5 text-[11px]",
        size === "md" && "px-2.5 py-1 text-xs",
        variants[variant],
        className,
      )}
    >
      {children}
    </span>
  );
}
