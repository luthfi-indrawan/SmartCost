import { forwardRef } from "react";
import { cn } from "../../lib/utils";

const variants = {
  primary:
    "bg-primary-500 text-white hover:bg-primary-600 active:bg-primary-700 shadow-[0_1px_3px_rgba(76,175,80,0.3)] hover:shadow-[0_4px_12px_rgba(76,175,80,0.4)]",
  secondary:
    "bg-secondary-500 text-white hover:bg-secondary-600 active:bg-secondary-700 shadow-[0_1px_3px_rgba(255,152,0,0.3)] hover:shadow-[0_4px_12px_rgba(255,152,0,0.4)]",
  ghost:
    "bg-transparent text-primary-500 border border-primary-500 hover:bg-primary-50 hover:text-primary-600 hover:border-primary-600 active:bg-primary-100",
  danger: "bg-danger-500 text-white hover:bg-red-600 active:bg-red-700",
  outline:
    "bg-transparent text-neutral-700 border border-neutral-200 hover:bg-neutral-50",
};

const sizes = {
  sm: "px-3 py-1.5 text-sm h-8",
  md: "px-5 py-2.5 text-[15px] font-semibold h-11",
  lg: "px-6 py-3 text-base h-12",
  icon: "p-2 h-10 w-10",
};

export const Button = forwardRef(function Button(
  {
    children,
    variant = "primary",
    size = "md",
    isLoading = false,
    disabled = false,
    fullWidth = false,
    className,
    ...props
  },
  ref,
) {
  return (
    <button
      ref={ref}
      disabled={disabled || isLoading}
      className={cn(
        "inline-flex items-center justify-center gap-2 rounded-[10px] transition-all duration-200",
        "disabled:opacity-50 disabled:cursor-not-allowed disabled:shadow-none",
        "active:translate-y-0 hover:-translate-y-px",
        variants[variant],
        sizes[size],
        fullWidth && "w-full",
        className,
      )}
      {...props}
    >
      {isLoading ? (
        <>
          <span className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin" />
          <span>Memuat...</span>
        </>
      ) : (
        children
      )}
    </button>
  );
});
