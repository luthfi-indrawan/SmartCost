import { PackageOpen, Search, Inbox } from "lucide-react";
import { cn } from "../../lib/utils";

const icons = {
  search: Search,
  empty: PackageOpen,
  inbox: Inbox,
};

export function EmptyState({
  title = "Tidak ada data",
  description = "Data yang Anda cari tidak ditemukan.",
  icon = "empty",
  action,
}) {
  const Icon = icons[icon];

  return (
    <div className="flex flex-col items-center justify-center py-12 px-4 text-center">
      <div className="w-16 h-16 rounded-full bg-neutral-100 flex items-center justify-center mb-4">
        <Icon className="w-8 h-8 text-neutral-400" />
      </div>
      <h3 className="text-base font-semibold text-neutral-700 mb-1">{title}</h3>
      <p className="text-sm text-neutral-500 max-w-xs">{description}</p>
      {action && <div className="mt-4">{action}</div>}
    </div>
  );
}
