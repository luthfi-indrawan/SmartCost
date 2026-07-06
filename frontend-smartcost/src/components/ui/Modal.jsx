import { useEffect, useRef } from "react";
import { motion, AnimatePresence } from "motion/react";
import { X } from "lucide-react";
import { cn } from "../../lib/utils";

export function Modal({
  isOpen,
  onClose,
  title,
  children,
  size = "md",
  showClose = true,
}) {
  const overlayRef = useRef(null);

  const sizeClasses = {
    sm: "max-w-sm",
    md: "max-w-md",
    lg: "max-w-lg",
    xl: "max-w-xl",
    full: "max-w-full mx-4",
  };

  useEffect(() => {
    const handleEscape = (e) => {
      if (e.key === "Escape") onClose();
    };

    if (isOpen) {
      document.addEventListener("keydown", handleEscape);
      document.body.style.overflow = "hidden";
    }

    return () => {
      document.removeEventListener("keydown", handleEscape);
      document.body.style.overflow = "";
    };
  }, [isOpen, onClose]);

  return (
    <AnimatePresence>
      {isOpen && (
        <motion.div
          ref={overlayRef}
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          exit={{ opacity: 0 }}
          transition={{ duration: 0.2 }}
          className="fixed inset-0 z-50 flex items-end sm:items-center justify-center bg-black/40 p-0 sm:p-4"
          onClick={(e) => {
            if (e.target === overlayRef.current) onClose();
          }}
        >
          <motion.div
            initial={{ y: "100%" }}
            animate={{ y: 0 }}
            exit={{ y: "100%" }}
            transition={{ type: "spring", damping: 25, stiffness: 300 }}
            className={cn(
              "w-full bg-wa-surface rounded-t-[18px] sm:rounded-[14px] shadow-xl",
              "max-h-[90vh] overflow-y-auto",
              sizeClasses[size],
            )}
          >
            {/* Header */}
            {(title || showClose) && (
              <div className="flex items-center justify-between px-5 py-4 border-b border-neutral-100">
                {title && (
                  <h3 className="text-lg font-semibold text-neutral-800">
                    {title}
                  </h3>
                )}
                {showClose && (
                  <button
                    onClick={onClose}
                    className="p-1.5 rounded-lg hover:bg-neutral-100 transition-colors"
                  >
                    <X className="w-5 h-5 text-neutral-500" />
                  </button>
                )}
              </div>
            )}

            {/* Content */}
            <div className="p-5">{children}</div>
          </motion.div>
        </motion.div>
      )}
    </AnimatePresence>
  );
}
