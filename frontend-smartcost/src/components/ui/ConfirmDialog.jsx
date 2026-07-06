import { Modal } from "./Modal";
import { Button } from "./Button";
import { AlertTriangle } from "lucide-react";

export function ConfirmDialog({
  isOpen,
  onClose,
  onConfirm,
  title = "Konfirmasi",
  message = "Apakah Anda yakin?",
  confirmLabel = "Ya, Lanjutkan",
  cancelLabel = "Batal",
  variant = "danger",
  isLoading = false,
}) {
  const handleConfirm = () => {
    onConfirm();
    if (!isLoading) onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="sm">
      <div className="flex flex-col items-center text-center py-2">
        <div className="w-14 h-14 rounded-full bg-danger-50 flex items-center justify-center mb-4">
          <AlertTriangle className="w-7 h-7 text-danger-500" />
        </div>
        <h3 className="text-lg font-semibold text-neutral-800 mb-2">{title}</h3>
        <p className="text-sm text-neutral-500 mb-6">{message}</p>
        <div className="flex gap-3 w-full">
          <Button
            variant="outline"
            fullWidth
            onClick={onClose}
            disabled={isLoading}
          >
            {cancelLabel}
          </Button>
          <Button
            variant={variant}
            fullWidth
            onClick={handleConfirm}
            isLoading={isLoading}
          >
            {confirmLabel}
          </Button>
        </div>
      </div>
    </Modal>
  );
}
