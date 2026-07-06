import { forwardRef } from "react";
import { Search, X } from "lucide-react";
import { cn } from "../../lib/utils";

export const SearchBar = forwardRef(function SearchBar(
  {
    value,
    onChange,
    placeholder = "Cari produk...",
    className,
    onClear,
    ...props
  },
  ref,
) {
  return (
    <div
      className={cn(
        "relative flex items-center w-full h-11 bg-neutral-100 rounded-full",
        "transition-all duration-200",
        "focus-within:bg-white focus-within:ring-[3px] focus-within:ring-primary-500/15",
        className,
      )}
    >
      <Search className="absolute left-4 w-5 h-5 text-neutral-400 pointer-events-none" />
      <input
        ref={ref}
        type="text"
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        className="w-full h-full pl-11 pr-10 bg-transparent text-sm text-neutral-800 placeholder:text-neutral-400 focus:outline-none"
        {...props}
      />
      {value && (
        <button
          onClick={onClear}
          className="absolute right-3 p-1 rounded-full hover:bg-neutral-200 transition-colors"
          type="button"
        >
          <X className="w-4 h-4 text-neutral-400" />
        </button>
      )}
    </div>
  );
});
