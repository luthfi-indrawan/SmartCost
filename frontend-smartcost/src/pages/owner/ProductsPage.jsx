import { useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Plus,
  Search,
  Package,
  Pencil,
  Trash2,
  AlertTriangle,
} from "lucide-react";
import { useProducts } from "../../hooks/useProducts";
import { useCategories } from "../../hooks/useCategories";
import { useDebounce } from "../../hooks/useDebounce";
import { SearchBar } from "../../components/ui/SearchBar";
import { Button } from "../../components/ui/Button";
import { Card } from "../../components/ui/Card";
import { Badge } from "../../components/ui/Badge";
import { Modal } from "../../components/ui/Modal";
import { Input } from "../../components/ui/Input";
import { ConfirmDialog } from "../../components/ui/ConfirmDialog";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { formatRupiah, getStockBadgeConfig } from "../../lib/constants";
import { cn } from "../../lib/utils";

export default function ProductsPage() {
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("");
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [productToDelete, setProductToDelete] = useState(null);

  const debouncedSearch = useDebounce(searchQuery, 300);

  const { products, isLoading, deleteProduct, isDeleting } = useProducts({
    search: debouncedSearch,
    category_id: selectedCategory,
    page_size: 50,
  });

  const { categories } = useCategories();

  const handleDelete = (product) => {
    setProductToDelete(product);
    setShowDeleteConfirm(true);
  };

  const confirmDelete = () => {
    if (productToDelete) {
      deleteProduct(productToDelete.id);
      setShowDeleteConfirm(false);
      setProductToDelete(null);
    }
  };

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6">
      {/* Header */}
      <div className="flex flex-col sm:flex-row gap-3">
        <SearchBar
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onClear={() => setSearchQuery("")}
          placeholder="Cari produk, SKU, atau barcode..."
          className="flex-1"
        />
        <Button
          onClick={() => navigate("/produk/new")}
          className="flex-shrink-0"
        >
          <Plus className="w-4 h-4" />
          Tambah Produk
        </Button>
      </div>

      {/* Category filter */}
      <div className="flex gap-2 overflow-x-auto pb-1">
        <button
          onClick={() => setSelectedCategory("")}
          className={cn(
            "flex-shrink-0 px-3 py-1.5 rounded-full text-xs font-medium transition-colors",
            !selectedCategory
              ? "bg-primary-500 text-white"
              : "bg-white text-neutral-600 border border-neutral-200",
          )}
        >
          Semua
        </button>
        {categories.map((cat) => (
          <button
            key={cat.id}
            onClick={() => setSelectedCategory(cat.id)}
            className={cn(
              "flex-shrink-0 px-3 py-1.5 rounded-full text-xs font-medium transition-colors",
              selectedCategory === cat.id
                ? "bg-primary-500 text-white"
                : "bg-white text-neutral-600 border border-neutral-200",
            )}
          >
            {cat.name}
          </button>
        ))}
      </div>

      {/* Product list */}
      {isLoading ? (
        <ListSkeleton count={5} />
      ) : products.length === 0 ? (
        <EmptyState
          title="Belum ada produk"
          description="Tambahkan produk pertama Anda untuk mulai berjualan."
          icon="empty"
          action={
            <Button onClick={() => navigate("/produk/new")} size="sm">
              <Plus className="w-4 h-4" />
              Tambah Produk
            </Button>
          }
        />
      ) : (
        <div className="space-y-3">
          {products.map((product) => {
            const stockConfig = getStockBadgeConfig(product.stock_status);

            return (
              <Card
                key={product.id}
                padding="normal"
                hover
                onClick={() => navigate(`/produk/${product.id}`)}
                className={cn(
                  "flex items-center gap-4 cursor-pointer",
                  product.stock_status === "MINUS" &&
                    "border-l-4 border-l-danger-500",
                )}
              >
                {/* Image placeholder */}
                <div className="w-14 h-14 bg-neutral-100 rounded-[10px] flex items-center justify-center flex-shrink-0">
                  <Package className="w-6 h-6 text-neutral-400" />
                </div>

                {/* Info */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <h4 className="text-sm font-semibold text-neutral-800 truncate">
                      {product.name}
                    </h4>
                    <Badge
                      variant={
                        product.stock_status === "SAFE"
                          ? "primary"
                          : product.stock_status === "LOW"
                            ? "warning"
                            : "danger"
                      }
                      size="sm"
                    >
                      {product.stock} {product.unit}
                    </Badge>
                  </div>
                  <div className="flex items-center gap-3 text-xs text-neutral-500">
                    <span className="font-mono-price font-semibold text-neutral-700">
                      {formatRupiah(
                        product.effective_price || product.base_price,
                      )}
                    </span>
                    <span>•</span>
                    <span>SKU: {product.sku}</span>
                    <span>•</span>
                    <span>{product.category?.name}</span>
                  </div>
                </div>

                {/* Actions */}
                <div className="flex items-center gap-1 flex-shrink-0">
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      navigate(`/produk/${product.id}`);
                    }}
                    className="p-2 rounded-lg hover:bg-neutral-100 text-neutral-400 hover:text-primary-600 transition-colors"
                  >
                    <Pencil className="w-4 h-4" />
                  </button>
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      handleDelete(product);
                    }}
                    className="p-2 rounded-lg hover:bg-danger-50 text-neutral-400 hover:text-danger-500 transition-colors"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </div>
              </Card>
            );
          })}
        </div>
      )}

      {/* Delete confirmation */}
      <ConfirmDialog
        isOpen={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        onConfirm={confirmDelete}
        title="Nonaktifkan Produk"
        message={`Apakah Anda yakin ingin menonaktifkan "${productToDelete?.name}"? Produk ini tidak akan muncul di kasir.`}
        confirmLabel="Nonaktifkan"
        isLoading={isDeleting}
      />
    </div>
  );
}
