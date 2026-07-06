import { useNavigate } from "react-router-dom";
import { ShieldAlert, ArrowLeft } from "lucide-react";
import { Button } from "../../components/ui/Button";

export default function ForbiddenPage() {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-wa-bg flex items-center justify-center p-4">
      <div className="text-center max-w-sm">
        <div className="w-20 h-20 bg-danger-50 rounded-full flex items-center justify-center mx-auto mb-6">
          <ShieldAlert className="w-10 h-10 text-danger-500" />
        </div>
        <h1 className="text-3xl font-bold text-neutral-800 mb-2">403</h1>
        <h2 className="text-lg font-semibold text-neutral-700 mb-2">
          Akses Ditolak
        </h2>
        <p className="text-sm text-neutral-500 mb-6">
          Anda tidak memiliki izin untuk mengakses halaman ini.
        </p>
        <Button onClick={() => navigate(-1)} variant="outline" fullWidth>
          <ArrowLeft className="w-4 h-4" />
          Kembali
        </Button>
      </div>
    </div>
  );
}
