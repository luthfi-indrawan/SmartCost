import { useParams, useNavigate } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { transactionsApi } from "../../api";
import { ReceiptView } from "../../components/cashier/ReceiptView";
import { Button } from "../../components/ui/Button";
import { ArrowLeft, Printer } from "lucide-react";

export default function ReceiptPage() {
  const { transactionId } = useParams();
  const navigate = useNavigate();

  const { data, isLoading } = useQuery({
    queryKey: ["transaction", transactionId],
    queryFn: async () => {
      const response = await transactionsApi.getById(transactionId);
      return response.data.result;
    },
  });

  const handlePrint = () => {
    window.print();
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-white flex items-center justify-center">
        <div className="animate-pulse space-y-4 w-full max-w-md px-4">
          <div className="h-8 bg-neutral-200 rounded w-1/2 mx-auto" />
          <div className="h-64 bg-neutral-200 rounded" />
        </div>
      </div>
    );
  }

  if (!data) {
    return (
      <div className="min-h-screen bg-white flex items-center justify-center">
        <div className="text-center">
          <p className="text-neutral-500">Transaksi tidak ditemukan</p>
          <Button onClick={() => navigate("/kasir")} className="mt-4">
            Kembali ke Kasir
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-white">
      <div className="sticky top-0 bg-white border-b border-neutral-200 px-4 py-3 flex items-center gap-3 no-print">
        <button
          onClick={() => navigate("/kasir")}
          className="p-2 rounded-lg hover:bg-neutral-100"
        >
          <ArrowLeft className="w-5 h-5 text-neutral-600" />
        </button>
        <h1 className="text-lg font-semibold text-neutral-800">
          Struk Penjualan
        </h1>
        <div className="flex-1" />
        <Button variant="ghost" size="sm" onClick={handlePrint}>
          <Printer className="w-4 h-4" />
          Cetak
        </Button>
      </div>

      <ReceiptView
        transaction={data}
        onNewTransaction={() => navigate("/kasir")}
        onPrint={handlePrint}
      />
    </div>
  );
}
