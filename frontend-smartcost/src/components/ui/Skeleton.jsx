import { cn } from "../../lib/utils";

export function Skeleton({ className, count = 1 }) {
  return (
    <>
      {Array.from({ length: count }).map((_, i) => (
        <div
          key={i}
          className={cn("animate-pulse bg-neutral-200 rounded-lg", className)}
        />
      ))}
    </>
  );
}

export function ProductCardSkeleton() {
  return (
    <div className="bg-white rounded-[14px] p-3 border border-neutral-200 space-y-3">
      <div className="aspect-square bg-neutral-200 rounded-[10px]" />
      <div className="space-y-2">
        <div className="h-4 bg-neutral-200 rounded w-3/4" />
        <div className="flex justify-between">
          <div className="h-5 bg-neutral-200 rounded w-1/3" />
          <div className="h-5 bg-neutral-200 rounded w-1/4" />
        </div>
      </div>
    </div>
  );
}

export function ListSkeleton({ count = 5 }) {
  return (
    <div className="space-y-3">
      {Array.from({ length: count }).map((_, i) => (
        <div
          key={i}
          className="flex items-center gap-3 p-3 bg-white rounded-[10px] border border-neutral-200"
        >
          <div className="w-10 h-10 bg-neutral-200 rounded-lg flex-shrink-0" />
          <div className="flex-1 space-y-2">
            <div className="h-4 bg-neutral-200 rounded w-1/2" />
            <div className="h-3 bg-neutral-200 rounded w-1/3" />
          </div>
        </div>
      ))}
    </div>
  );
}
