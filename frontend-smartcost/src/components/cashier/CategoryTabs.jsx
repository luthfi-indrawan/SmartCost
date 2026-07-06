import { cn } from "../../lib/utils";

export function CategoryTabs({ categories, activeCategory, onSelect }) {
  return (
    <div className="sticky top-14 z-30 bg-wa-bg border-b border-neutral-200">
      <div className="flex gap-2 px-4 py-3 overflow-x-auto scrollbar-hide">
        <button
          onClick={() => onSelect("")}
          className={cn(
            "flex-shrink-0 px-4 py-2 rounded-full text-sm font-medium transition-colors",
            !activeCategory
              ? "bg-primary-500 text-white shadow-[0_1px_3px_rgba(76,175,80,0.3)]"
              : "bg-white text-neutral-600 border border-neutral-200 hover:bg-neutral-50",
          )}
        >
          Semua
        </button>
        {categories.map((category) => (
          <button
            key={category.id}
            onClick={() => onSelect(category.id)}
            className={cn(
              "flex-shrink-0 px-4 py-2 rounded-full text-sm font-medium transition-colors",
              activeCategory === category.id
                ? "bg-primary-500 text-white shadow-[0_1px_3px_rgba(76,175,80,0.3)]"
                : "bg-white text-neutral-600 border border-neutral-200 hover:bg-neutral-50",
            )}
          >
            {category.name}
          </button>
        ))}
      </div>
    </div>
  );
}
