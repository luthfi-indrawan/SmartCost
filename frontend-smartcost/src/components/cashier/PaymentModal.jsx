import { useState } from "react";
import { Modal } from "../ui/Modal";
import { Button } from "../ui/Button";
import { Input } from "../ui/Input";
import { formatRupiah } from "../../lib/constants";
import { Banknote, QrCode, ArrowRightLeft } from "lucide-react";
import { cn } from "../../lib/utils";

const paymentMethods = [
  { id: "cash", label: "Tunai", icon: Banknote },
  { id: "qris", label: "QRIS", icon: QrCode },
  { id: "transfer", label: "Transfer", icon: ArrowRightLeft },
];

export function PaymentModal({ isOpen, onClose, total, onConfirm }) {
  const [method, setMethod] = useState("cash");
  const [amountPaid, setAmountPaid] = useState("");
  const [isProcessing, setIsProcessing] = useState(false);

  const paidAmount = parseInt(amountPaid.replace(/\D/g, "")) || 0;
  const change = paidAmount - total;

  const handleQuickAmount = (amount) => {
    setAmountPaid(amount.toString());
  };

  const quickAmounts = [
    total,
    Math.ceil(total / 1000) * 1000,
    Math.ceil(total / 5000) * 5000,
    Math.ceil(total / 10000) * 10000,
    50000,
    100000,
  ]
    .filter((v, i, a) => a.indexOf(v) === i)
    .slice(0, 4);

  const handleSubmit = async () => {
    if (method === "cash" && paidAmount < total) return;
    setIsProcessing(true);
    await onConfirm({
      method,
      amount_paid: method === "cash" ? paidAmount : total,
      change: method === "cash" ? change : 0,
    });
    setIsProcessing(false);
    setAmountPaid("");
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Pembayaran" size="md">
      <div className="space-y-5">
        {/* Total */}
        <div className="text-center py-4 bg-primary-50 rounded-[14px]">
          <p className="text-sm text-primary-700 mb-1">Total Pembayaran</p>
          <p className="text-3xl font-bold text-primary-800 font-mono-price">
            {formatRupiah(total)}
          </p>
        </div>

        {/* Payment methods */}
        <div>
          <label className="block text-sm font-medium text-neutral-700 mb-2">
            Metode Pembayaran
          </label>
          <div className="grid grid-cols-3 gap-2">
            {paymentMethods.map((pm) => (
              <button
                key={pm.id}
                onClick={() => setMethod(pm.id)}
                className={cn(
                  "flex flex-col items-center gap-2 p-3 rounded-[10px] border-2 transition-all",
                  method === pm.id
                    ? "border-primary-500 bg-primary-50 text-primary-700"
                    : "border-neutral-200 bg-white text-neutral-600 hover:border-neutral-300",
                )}
              >
                <pm.icon className="w-6 h-6" />
                <span className="text-xs font-medium">{pm.label}</span>
              </button>
            ))}
          </div>
        </div>

        {/* Cash input */}
        {method === "cash" && (
          <div className="space-y-3">
            <Input
              label="Jumlah Dibayar"
              type="text"
              value={amountPaid}
              onChange={(e) => {
                const val = e.target.value.replace(/\D/g, "");
                setAmountPaid(val);
              }}
              placeholder="0"
              className="text-lg font-bold text-center"
            />

            {/* Quick amounts */}
            <div className="flex flex-wrap gap-2">
              {quickAmounts.map((amount) => (
                <button
                  key={amount}
                  onClick={() => handleQuickAmount(amount)}
                  className="px-3 py-1.5 text-xs font-medium bg-neutral-100 text-neutral-700 rounded-full hover:bg-primary-50 hover:text-primary-700 transition-colors"
                >
                  {formatRupiah(amount)}
                </button>
              ))}
            </div>

            {/* Change */}
            {paidAmount > 0 && (
              <div
                className={cn(
                  "p-3 rounded-[10px] text-center",
                  change >= 0 ? "bg-success-50" : "bg-danger-50",
                )}
              >
                <p className="text-sm text-neutral-600">Kembalian</p>
                <p
                  className={cn(
                    "text-xl font-bold font-mono-price",
                    change >= 0 ? "text-success-600" : "text-danger-600",
                  )}
                >
                  {formatRupiah(change)}
                </p>
              </div>
            )}
          </div>
        )}

        {/* Confirm button */}
        <Button
          fullWidth
          disabled={method === "cash" && paidAmount < total}
          isLoading={isProcessing}
          onClick={handleSubmit}
        >
          {method === "cash"
            ? `Bayar ${formatRupiah(total)}`
            : `Konfirmasi ${formatRupiah(total)}`}
        </Button>
      </div>
    </Modal>
  );
}
