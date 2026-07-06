import { useState } from "react";
import { Modal } from "../ui/Modal";
import { Button } from "../ui/Button";
import { Input } from "../ui/Input";
import { useCartStore } from "../../stores/cartStore";

export function HoldBillModal({ isOpen, onClose, onConfirm }) {
  const [note, setNote] = useState("");
  const [isProcessing, setIsProcessing] = useState(false);

  const handleSubmit = async () => {
    if (!note.trim()) return;
    setIsProcessing(true);
    await onConfirm(note.trim());
    setIsProcessing(false);
    setNote("");
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Simpan Pesanan" size="sm">
      <div className="space-y-4">
        <p className="text-sm text-neutral-500">
          Pesanan akan disimpan dan stok langsung dipotong. Masukkan catatan
          untuk mengidentifikasi pesanan.
        </p>

        <Input
          label="Catatan Pesanan"
          value={note}
          onChange={(e) => setNote(e.target.value)}
          placeholder="Contoh: Meja 5 / Andi"
          maxLength={100}
        />

        <div className="flex gap-3">
          <Button variant="outline" fullWidth onClick={onClose}>
            Batal
          </Button>
          <Button
            fullWidth
            disabled={!note.trim()}
            isLoading={isProcessing}
            onClick={handleSubmit}
          >
            Simpan Hold Bill
          </Button>
        </div>
      </div>
    </Modal>
  );
}
