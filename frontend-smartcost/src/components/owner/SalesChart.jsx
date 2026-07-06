import { useMemo } from "react";
import { format, parseISO } from "date-fns";
import { id } from "date-fns/locale";

export function SalesChart({ data = [] }) {
  const maxValue = useMemo(() => {
    if (!data.length) return 0;
    return Math.max(...data.map((d) => d.revenue));
  }, [data]);

  if (!data.length) {
    return (
      <div className="h-48 flex items-center justify-center text-neutral-400 text-sm">
        Tidak ada data penjualan
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-end justify-between gap-2 h-48 px-2">
        {data.map((item, index) => {
          const height = maxValue > 0 ? (item.revenue / maxValue) * 100 : 0;
          const dayLabel = format(parseISO(item.date), "EEE", { locale: id });

          return (
            <div
              key={index}
              className="flex-1 flex flex-col items-center gap-2"
            >
              <div className="w-full flex items-end justify-center h-36">
                <div
                  className="w-full max-w-[32px] bg-primary-500 rounded-t-md transition-all duration-500 hover:bg-primary-600"
                  style={{ height: `${Math.max(height, 4)}%` }}
                  title={`${format(parseISO(item.date), "dd MMM")}: ${new Intl.NumberFormat(
                    "id-ID",
                    {
                      style: "currency",
                      currency: "IDR",
                      maximumFractionDigits: 0,
                    },
                  ).format(item.revenue)}`}
                />
              </div>
              <span className="text-[10px] text-neutral-500 font-medium uppercase">
                {dayLabel}
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
}
