import { Printer, CheckCircle, ArrowLeft, Plus } from "lucide-react";
import { formatRupiah, formatDateTime } from "../../lib/constants";
import { Button } from "../ui/Button";
import { useCartStore } from "../../stores/cartStore";

export function ReceiptView({ transaction, onNewTransaction, onPrint }) {
  const clearCart = useCartStore((state) => state.clearCart);

  const handleNewTransaction = () => {
    clearCart();
    onNewTransaction();
  };

  return (
    <div className="max-w-md mx-auto bg-white min-h-screen">
      {/* Header */}
      <div className="sticky top-0 bg-white border-b border-neutral-200 px-4 py-3 flex items-center gap-3">
        <button
          onClick={handleNewTransaction}
          className="p-2 rounded-lg hover:bg-neutral-100"
        >
          <ArrowLeft className="w-5 h-5 text-neutral-600" />
        </button>
        <h1 className="text-lg font-semibold text-neutral-800">
          Struk Penjualan
        </h1>
      </div>

      <div className="p-6 space-y-6">
        {/* Store info */}
        <div className="text-center space-y-1">
          <h2 className="text-xl font-bold text-neutral-900">SMART COST</h2>
          <p className="text-sm text-neutral-500">Struk Penjualan</p>
        </div>

        {/* Transaction info */}
        <div className="receipt-line pt-4 space-y-1 text-sm">
          <div className="flex justify-between">
            <span className="text-neutral-500">No. Transaksi</span>
            <span className="font-mono text-neutral-800">
              {transaction.transaction_code}
            </span>
          </div>
          <div className="flex justify-between">
            <span className="text-neutral-500">Waktu</span>
            <span className="text-neutral-800">
              {formatDateTime(transaction.created_at)}
            </span>
          </div>
          <div className="flex justify-between">
            <span className="text-neutral-500">Kasir</span>
            <span className="text-neutral-800">
              {transaction.cashier?.name}
            </span>
          </div>
        </div>

        {/* Items */}
        <div className="receipt-line pt-4 space-y-3">
          {transaction.items?.map((item, index) => (
            <div key={item.id || index} className="text-sm">
              <div className="flex justify-between">
                <span className="text-neutral-800">
                  {index + 1}. {item.product?.name}
                </span>
              </div>
              <div className="flex justify-between text-neutral-500 pl-4">
                <span>
                  {item.qty} x {formatRupiah(item.unit_price)}
                </span>
                <span className="font-mono-price">
                  {formatRupiah(item.subtotal)}
                </span>
              </div>
              {item.notes && (
                <p className="text-xs text-neutral-400 pl-4 italic">
                  Catatan: {item.notes}
                </p>
              )}
            </div>
          ))}
        </div>

        {/* Totals */}
        <div className="receipt-line pt-4 space-y-2">
          <div className="flex justify-between text-sm">
            <span className="text-neutral-600">Subtotal</span>
            <span className="font-mono-price text-neutral-800">
              {formatRupiah(transaction.summary?.subtotal)}
            </span>
          </div>
          {transaction.summary?.discount_amount > 0 && (
            <div className="flex justify-between text-sm">
              <span className="text-neutral-600">Diskon</span>
              <span className="font-mono-price text-success-600">
                -{formatRupiah(transaction.summary.discount_amount)}
              </span>
            </div>
          )}
          <div className="flex justify-between text-base font-bold pt-2">
            <span className="text-neutral-800">Total</span>
            <span className="font-mono-price text-neutral-900">
              {formatRupiah(transaction.summary?.total)}
            </span>
          </div>
        </div>

        {/* Payment */}
        {transaction.payment && (
          <div className="receipt-line pt-4 space-y-1 text-sm">
            <div className="flex justify-between">
              <span className="text-neutral-500">Metode</span>
              <span className="text-neutral-800 capitalize">
                {transaction.payment.method}
              </span>
            </div>
            <div className="flex justify-between">
              <span className="text-neutral-500">Dibayar</span>
              <span className="font-mono-price text-neutral-800">
                {formatRupiah(transaction.payment.amount_paid)}
              </span>
            </div>
            <div className="flex justify-between">
              <span className="text-neutral-500">Kembalian</span>
              <span className="font-mono-price text-success-600">
                {formatRupiah(transaction.payment.change)}
              </span>
            </div>
          </div>
        )}

        {/* Success indicator */}
        <div className="flex items-center justify-center gap-2 py-4">
          <CheckCircle className="w-5 h-5 text-primary-500" />
          <span className="text-sm font-medium text-primary-700">
            Transaksi Berhasil
          </span>
        </div>

        {/* Footer */}
        <div className="text-center text-xs text-neutral-400 pt-4 receipt-line">
          <p>Terima kasih telah berbelanja!</p>
          <p className="mt-1">SmartCost - Solusi POS UMKM</p>
        </div>
      </div>

      {/* Actions */}
      <div className="sticky bottom-0 bg-white border-t border-neutral-200 p-4 space-y-2 no-print">
        <Button fullWidth variant="secondary" onClick={onPrint}>
          <Printer className="w-4 h-4" />
          Cetak Struk
        </Button>
        <Button fullWidth onClick={handleNewTransaction}>
          <Plus className="w-4 h-4" />
          Transaksi Baru
        </Button>
      </div>
    </div>
  );
}
