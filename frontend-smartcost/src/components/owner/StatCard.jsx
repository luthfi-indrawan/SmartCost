import { ArrowUpRight, ArrowDownRight } from "lucide-react";
import { Card } from "../ui/Card";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";

export function StatCard({
  title,
  value,
  subtitle,
  trend,
  trendValue,
  icon: Icon,
  color = "primary",
}) {
  const colorMap = {
    primary: "bg-primary-50 text-primary-600",
    secondary: "bg-secondary-50 text-secondary-600",
    success: "bg-success-50 text-success-600",
    danger: "bg-danger-50 text-danger-600",
  };

  const isPositive = trend === "up";

  return (
    <Card padding="lg" className="h-full">
      <div className="flex items-start justify-between">
        <div
          className={cn(
            "w-10 h-10 rounded-[10px] flex items-center justify-center",
            colorMap[color],
          )}
        >
          <Icon className="w-5 h-5" />
        </div>
        {trend && (
          <div
            className={cn(
              "flex items-center gap-0.5 text-xs font-medium",
              isPositive ? "text-success-500" : "text-danger-500",
            )}
          >
            {isPositive ? (
              <ArrowUpRight className="w-3.5 h-3.5" />
            ) : (
              <ArrowDownRight className="w-3.5 h-3.5" />
            )}
            {trendValue}
          </div>
        )}
      </div>
      <div className="mt-4">
        <p className="text-2xl font-bold text-neutral-900 font-mono-price">
          {typeof value === "number" ? formatRupiah(value) : value}
        </p>
        <p className="text-sm text-neutral-500 mt-1">{title}</p>
        {subtitle && (
          <p className="text-xs text-neutral-400 mt-0.5">{subtitle}</p>
        )}
      </div>
    </Card>
  );
}
