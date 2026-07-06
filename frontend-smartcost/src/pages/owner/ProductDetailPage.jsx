import { useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { ArrowLeft, Package, Plus, Trash2, Save, Barcode } from "lucide-react";
import { useProducts } from "../../hooks/useProducts";
import { useCategories } from "../../hooks/useCategories";
import { Button } from "../../components/ui/Button";
import { Input } from "../../components/ui/Input";
import { Card } from "../../components/ui/Card";
import { Badge } from "../../components/ui/Badge";
import { Modal } from "../../components/ui/Modal";
import { formatRupiah, getStockBadgeConfig } from "../../lib/constants";
import { cn } from "../../lib/utils";

const priceTierSchema = z.object({
  min_qty: z.coerce.number().min(2, "Min qty harus > 1"),
  price: z.coerce.number().min(1, "Harga harus > 0"),
  label: z.string().min(1, "Label wajib diisi"),
});

const productSchema = z.object({
  name: z.string().min(1, "Nama produk wajib diisi"),
  sku: z.string().min(1, "SKU wajib diisi"),
  barcode: z.string().optional(),
  category_id: z.string().min(1, "Kategori wajib dipilih"),
  base_price: z.coerce.number().min(1, "Harga dasar harus > 0"),
  stock: z.coerce.number().min(0, "Stok minimal 0"),
  min_stock_threshold: z.coerce.number().min(0, "Threshold minimal 0"),
  unit: z.string().min(1, "Satuan wajib diisi"),
  description: z.string().optional(),
  price_tiers: z.array(priceTierSchema).optional(),
});

export default function ProductDetailPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const isNew = id === "new";
  const [tiers, setTiers] = useState([]);

  const { products, isLoading, createProduct, updateProduct } = useProducts();
  const { categories } = useCategories();

  const product = isNew ? null : products.find((p) => p.id === id);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    watch,
  } = useForm({
    resolver: zodResolver(productSchema),
    defaultValues: {
      name: "",
      sku: "",
      barcode: "",
      category_id: "",
      base_price: "",
      stock: "",
      min_stock_threshold: "10",
      unit: "pcs",
      description: "",
      price_tiers: [],
    },
  });

  // Populate form when product loads
  useState(() => {
    if (product && !isNew) {
      setValue("name", product.name);
      setValue("sku", product.sku);
      setValue("barcode", product.barcode || "");
      setValue("category_id", product.category?.id || "");
      setValue("base_price", product.base_price);
      setValue("stock", product.stock);
      setValue("min_stock_threshold", product.min_stock_threshold);
      setValue("unit", product.unit);
      setValue("description", product.description || "");
      setTiers(product.price_tiers || []);
    }
  }, [product]);

  const addTier = () => {
    setTiers([...tiers, { min_qty: "", price: "", label: "" }]);
  };

  const removeTier = (index) => {
    setTiers(tiers.filter((_, i) => i !== index));
  };

  const updateTier = (index, field, value) => {
    const newTiers = [...tiers];
    newTiers[index][field] = value;
    setTiers(newTiers);
  };

  const onSubmit = async (data) => {
    const payload = {
      ...data,
      price_tiers: tiers.filter((t) => t.min_qty && t.price && t.label),
    };

    try {
      if (isNew) {
        await createProduct(payload);
      } else {
        await updateProduct({ id, data: payload });
      }
      navigate("/produk");
    } catch (error) {
      // Error handled by mutation hook
    }
  };

  if (!isNew && isLoading) {
    return (
      <div className="p-4">
        <div className="animate-pulse space-y-4">
          <div className="h-8 bg-neutral-200 rounded w-1/3" />
          <div className="h-64 bg-neutral-200 rounded-[14px]" />
        </div>
      </div>
    );
  }

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6 max-w-2xl">
      {/* Header */}
      <div className="flex items-center gap-3">
        <button
          onClick={() => navigate("/produk")}
          className="p-2 rounded-lg hover:bg-neutral-100"
        >
          <ArrowLeft className="w-5 h-5 text-neutral-600" />
        </button>
        <h1 className="text-xl font-bold text-neutral-800">
          {isNew ? "Tambah Produk" : "Edit Produk"}
        </h1>
      </div>

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
        {/* Basic Info */}
        <Card padding="lg">
          <h3 className="text-sm font-semibold text-neutral-500 uppercase tracking-wide mb-4">
            Informasi Dasar
          </h3>
          <div className="space-y-4">
            <Input
              label="Nama Produk"
              placeholder="Contoh: Teh Kotak 250ml"
              error={errors.name?.message}
              {...register("name")}
            />

            <div className="grid grid-cols-2 gap-3">
              <Input
                label="SKU"
                placeholder="TK-001"
                error={errors.sku?.message}
                {...register("sku")}
              />
              <Input
                label="Barcode (Opsional)"
                placeholder="8991234567890"
                icon={Barcode}
                {...register("barcode")}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-neutral-700 mb-1.5">
                Kategori
              </label>
              <select
                {...register("category_id")}
                className={cn(
                  "w-full h-12 px-4 text-[15px] bg-white border rounded-[10px]",
                  "focus:outline-none focus:border-primary-500 focus:ring-[3px] focus:ring-primary-500/15",
                  errors.category_id && "border-danger-500",
                )}
              >
                <option value="">Pilih Kategori</option>
                {categories.map((cat) => (
                  <option key={cat.id} value={cat.id}>
                    {cat.name}
                  </option>
                ))}
              </select>
              {errors.category_id && (
                <p className="mt-1 text-sm text-danger-500">
                  {errors.category_id.message}
                </p>
              )}
            </div>

            <Input
              label="Deskripsi (Opsional)"
              placeholder="Deskripsi produk..."
              {...register("description")}
            />
          </div>
        </Card>

        {/* Pricing & Stock */}
        <Card padding="lg">
          <h3 className="text-sm font-semibold text-neutral-500 uppercase tracking-wide mb-4">
            Harga & Stok
          </h3>
          <div className="space-y-4">
            <div className="grid grid-cols-2 gap-3">
              <Input
                label="Harga Dasar"
                type="number"
                placeholder="3000"
                error={errors.base_price?.message}
                {...register("base_price")}
              />
              <Input
                label="Stok"
                type="number"
                placeholder="50"
                error={errors.stock?.message}
                {...register("stock")}
              />
            </div>
            <div className="grid grid-cols-2 gap-3">
              <Input
                label="Batas Minimal Stok"
                type="number"
                placeholder="10"
                error={errors.min_stock_threshold?.message}
                {...register("min_stock_threshold")}
              />
              <Input
                label="Satuan"
                placeholder="pcs"
                error={errors.unit?.message}
                {...register("unit")}
              />
            </div>
          </div>
        </Card>

        {/* Price Tiers */}
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-sm font-semibold text-neutral-500 uppercase tracking-wide">
              Harga Grosir
            </h3>
            <Button type="button" variant="ghost" size="sm" onClick={addTier}>
              <Plus className="w-4 h-4" />
              Tambah Tier
            </Button>
          </div>

          {tiers.length === 0 ? (
            <p className="text-sm text-neutral-400 text-center py-4">
              Belum ada harga grosir. Tambahkan untuk mengatur harga berbasis
              kuantitas.
            </p>
          ) : (
            <div className="space-y-3">
              {tiers.map((tier, index) => (
                <div
                  key={index}
                  className="flex items-end gap-2 p-3 bg-neutral-50 rounded-[10px]"
                >
                  <div className="flex-1">
                    <label className="block text-xs font-medium text-neutral-500 mb-1">
                      Min Qty
                    </label>
                    <input
                      type="number"
                      value={tier.min_qty}
                      onChange={(e) =>
                        updateTier(index, "min_qty", e.target.value)
                      }
                      className="w-full h-9 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
                      placeholder="10"
                    />
                  </div>
                  <div className="flex-1">
                    <label className="block text-xs font-medium text-neutral-500 mb-1">
                      Harga
                    </label>
                    <input
                      type="number"
                      value={tier.price}
                      onChange={(e) =>
                        updateTier(index, "price", e.target.value)
                      }
                      className="w-full h-9 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
                      placeholder="2500"
                    />
                  </div>
                  <div className="flex-[1.5]">
                    <label className="block text-xs font-medium text-neutral-500 mb-1">
                      Label
                    </label>
                    <input
                      type="text"
                      value={tier.label}
                      onChange={(e) =>
                        updateTier(index, "label", e.target.value)
                      }
                      className="w-full h-9 px-3 text-sm bg-white border border-neutral-200 rounded-lg focus:outline-none focus:border-primary-500"
                      placeholder="Grosir 10pcs"
                    />
                  </div>
                  <button
                    type="button"
                    onClick={() => removeTier(index)}
                    className="p-2 rounded-lg hover:bg-danger-50 text-neutral-400 hover:text-danger-500 transition-colors"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </div>
              ))}
            </div>
          )}
        </Card>

        {/* Submit */}
        <div className="flex gap-3">
          <Button
            type="button"
            variant="outline"
            fullWidth
            onClick={() => navigate("/produk")}
          >
            Batal
          </Button>
          <Button type="submit" fullWidth>
            <Save className="w-4 h-4" />
            {isNew ? "Simpan Produk" : "Perbarui Produk"}
          </Button>
        </div>
      </form>
    </div>
  );
}
