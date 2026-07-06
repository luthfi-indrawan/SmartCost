import { ProductCard } from "./ProductCard";
import { ProductCardSkeleton } from "../ui/Skeleton";
import { EmptyState } from "../ui/EmptyState";

export function ProductGrid({ products, isLoading, onProductClick }) {
  if (isLoading) {
    return (
      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 p-4">
        {Array.from({ length: 8 }).map((_, i) => (
          <ProductCardSkeleton key={i} />
        ))}
      </div>
    );
  }

  if (!products || products.length === 0) {
    return (
      <div className="p-4">
        <EmptyState
          title="Produk tidak ditemukan"
          description="Coba ubah kata kunci pencarian atau filter kategori."
          icon="search"
        />
      </div>
    );
  }

  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 p-4 pb-24">
      {products.map((product) => (
        <ProductCard
          key={product.id}
          product={product}
          onClick={onProductClick}
        />
      ))}
    </div>
  );
}
