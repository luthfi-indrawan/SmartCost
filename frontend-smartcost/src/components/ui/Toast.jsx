import { useUiStore } from "../../stores/uiStore";
import { CheckCircle, XCircle, AlertCircle, X } from "lucide-react";
import { motion, AnimatePresence } from "motion/react";
import { cn } from "../../lib/utils";

const toastConfig = {
  success: {
    bg: "bg-primary-50",
    border: "border-primary-200",
    text: "text-primary-800",
    icon: CheckCircle,
    iconColor: "text-primary-500",
  },
  error: {
    bg: "bg-danger-50",
    border: "border-danger-200",
    text: "text-danger-800",
    icon: XCircle,
    iconColor: "text-danger-500",
  },
  warning: {
    bg: "bg-warning-50",
    border: "border-warning-200",
    text: "text-warning-800",
    icon: AlertCircle,
    iconColor: "text-warning-500",
  },
};

export function Toast() {
  const toast = useUiStore((state) => state.toast);

  if (!toast) return null;

  const config = toastConfig[toast.type] || toastConfig.success;
  const Icon = config.icon;

  return (
    <AnimatePresence>
      <motion.div
        initial={{ opacity: 0, y: -20, x: "-50%" }}
        animate={{ opacity: 1, y: 0, x: "-50%" }}
        exit={{ opacity: 0, y: -20, x: "-50%" }}
        transition={{ duration: 0.2 }}
        className={cn(
          "fixed top-4 left-1/2 z-[100] flex items-center gap-2 px-4 py-3 rounded-lg shadow-lg border",
          config.bg,
          config.border,
        )}
      >
        <Icon className={cn("w-5 h-5 flex-shrink-0", config.iconColor)} />
        <span className={cn("text-sm font-medium", config.text)}>
          {toast.message}
        </span>
      </motion.div>
    </AnimatePresence>
  );
}
