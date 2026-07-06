import { useNavigate } from "react-router-dom";
import { FileQuestion, Home } from "lucide-react";
import { Button } from "../../components/ui/Button";

export default function NotFoundPage() {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-wa-bg flex items-center justify-center p-4">
      <div className="text-center max-w-sm">
        <div className="w-20 h-20 bg-neutral-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <FileQuestion className="w-10 h-10 text-neutral-400" />
        </div>
        <h1 className="text-3xl font-bold text-neutral-800 mb-2">404</h1>
        <h2 className="text-lg font-semibold text-neutral-700 mb-2">
          Halaman Tidak Ditemukan
        </h2>
        <p className="text-sm text-neutral-500 mb-6">
          Halaman yang Anda cari tidak tersedia atau telah dipindahkan.
        </p>
        <Button onClick={() => navigate("/")} fullWidth>
          <Home className="w-4 h-4" />
          Kembali ke Beranda
        </Button>
      </div>
    </div>
  );
}
