import { forwardRef } from "react";
import { cn } from "../../lib/utils";

export const Input = forwardRef(function Input(
  {
    label,
    error,
    helperText,
    icon: Icon,
    className,
    containerClassName,
    ...props
  },
  ref,
) {
  return (
    <div className={cn("w-full", containerClassName)}>
      {label && (
        <label className="block text-sm font-medium text-neutral-700 mb-1.5">
          {label}
        </label>
      )}
      <div className="relative">
        {Icon && (
          <div className="absolute left-4 top-1/2 -translate-y-1/2 pointer-events-none">
            <Icon className="w-5 h-5 text-neutral-400" />
          </div>
        )}
        <input
          ref={ref}
          className={cn(
            "w-full h-12 px-4 text-[15px] bg-white border rounded-[10px]",
            "placeholder:text-neutral-400 text-neutral-800",
            "transition-all duration-200",
            "focus:outline-none focus:border-primary-500 focus:ring-[3px] focus:ring-primary-500/15",
            "hover:border-neutral-300",
            "disabled:bg-neutral-100 disabled:text-neutral-400 disabled:cursor-not-allowed",
            error &&
              "border-danger-500 bg-danger-50 focus:border-danger-500 focus:ring-danger-500/15",
            Icon && "pl-11",
            className,
          )}
          {...props}
        />
      </div>
      {error && <p className="mt-1.5 text-sm text-danger-500">{error}</p>}
      {helperText && !error && (
        <p className="mt-1.5 text-sm text-neutral-500">{helperText}</p>
      )}
    </div>
  );
});
