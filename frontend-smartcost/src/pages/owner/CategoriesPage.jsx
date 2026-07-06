import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Plus, Pencil, Trash2, Tag, Package } from "lucide-react";
import { useCategories } from "../../hooks/useCategories";
import { Button } from "../../components/ui/Button";
import { Input } from "../../components/ui/Input";
import { Card } from "../../components/ui/Card";
import { Modal } from "../../components/ui/Modal";
import { ConfirmDialog } from "../../components/ui/ConfirmDialog";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { cn } from "../../lib/utils";

const categorySchema = z.object({
  name: z.string().min(1, "Nama kategori wajib diisi"),
  color: z
    .string()
    .regex(/^#[0-9A-Fa-f]{6}$/, "Format HEX tidak valid")
    .default("#6B7280"),
  description: z.string().optional(),
});

export default function CategoriesPage() {
  const [showModal, setShowModal] = useState(false);
  const [editingCategory, setEditingCategory] = useState(null);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [categoryToDelete, setCategoryToDelete] = useState(null);

  const {
    categories,
    isLoading,
    createCategory,
    updateCategory,
    deleteCategory,
    isCreating,
    isUpdating,
    isDeleting,
  } = useCategories();

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(categorySchema),
    defaultValues: {
      name: "",
      color: "#6B7280",
      description: "",
    },
  });

  const openCreate = () => {
    setEditingCategory(null);
    reset({ name: "", color: "#6B7280", description: "" });
    setShowModal(true);
  };

  const openEdit = (category) => {
    setEditingCategory(category);
    setValue("name", category.name);
    setValue("color", category.color);
    setValue("description", category.description || "");
    setShowModal(true);
  };

  const handleDelete = (category) => {
    setCategoryToDelete(category);
    setShowDeleteConfirm(true);
  };

  const onSubmit = async (data) => {
    try {
      if (editingCategory) {
        await updateCategory({ id: editingCategory.id, data });
      } else {
        await createCategory(data);
      }
      setShowModal(false);
      reset();
    } catch (error) {
      // Error handled by mutation hook
    }
  };

  const confirmDelete = () => {
    if (categoryToDelete) {
      deleteCategory(categoryToDelete.id);
      setShowDeleteConfirm(false);
      setCategoryToDelete(null);
    }
  };

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6">
      <div className="flex items-center justify-between">
        <h2 className="text-lg font-semibold text-neutral-800">
          Kategori Produk
        </h2>
        <Button onClick={openCreate} size="sm">
          <Plus className="w-4 h-4" />
          Tambah
        </Button>
      </div>

      {isLoading ? (
        <ListSkeleton count={4} />
      ) : categories.length === 0 ? (
        <EmptyState
          title="Belum ada kategori"
          description="Buat kategori untuk mengelompokkan produk Anda."
          icon="empty"
          action={
            <Button onClick={openCreate} size="sm">
              <Plus className="w-4 h-4" />
              Tambah Kategori
            </Button>
          }
        />
      ) : (
        <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-3">
          {categories.map((category) => (
            <Card
              key={category.id}
              padding="normal"
              className="flex items-center gap-3"
            >
              <div
                className="w-10 h-10 rounded-[10px] flex items-center justify-center flex-shrink-0"
                style={{ backgroundColor: category.color + "20" }}
              >
                <Tag className="w-5 h-5" style={{ color: category.color }} />
              </div>
              <div className="flex-1 min-w-0">
                <h4 className="text-sm font-semibold text-neutral-800">
                  {category.name}
                </h4>
                <div className="flex items-center gap-2 text-xs text-neutral-500">
                  <Package className="w-3 h-3" />
                  {category.product_count} produk
                </div>
              </div>
              <div className="flex items-center gap-1">
                <button
                  onClick={() => openEdit(category)}
                  className="p-2 rounded-lg hover:bg-neutral-100 text-neutral-400 hover:text-primary-600"
                >
                  <Pencil className="w-4 h-4" />
                </button>
                <button
                  onClick={() => handleDelete(category)}
                  className="p-2 rounded-lg hover:bg-danger-50 text-neutral-400 hover:text-danger-500"
                >
                  <Trash2 className="w-4 h-4" />
                </button>
              </div>
            </Card>
          ))}
        </div>
      )}

      {/* Create/Edit Modal */}
      <Modal
        isOpen={showModal}
        onClose={() => setShowModal(false)}
        title={editingCategory ? "Edit Kategori" : "Tambah Kategori"}
        size="sm"
      >
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <Input
            label="Nama Kategori"
            placeholder="Contoh: Minuman"
            error={errors.name?.message}
            {...register("name")}
          />
          <div>
            <label className="block text-sm font-medium text-neutral-700 mb-1.5">
              Warna
            </label>
            <div className="flex items-center gap-3">
              <input
                type="color"
                {...register("color")}
                className="w-12 h-12 rounded-lg border border-neutral-200 cursor-pointer"
              />
              <Input
                placeholder="#6B7280"
                className="flex-1"
                {...register("color")}
              />
            </div>
            {errors.color && (
              <p className="mt-1 text-sm text-danger-500">
                {errors.color.message}
              </p>
            )}
          </div>
          <Input
            label="Deskripsi (Opsional)"
            placeholder="Deskripsi kategori..."
            {...register("description")}
          />
          <div className="flex gap-3 pt-2">
            <Button
              variant="outline"
              fullWidth
              onClick={() => setShowModal(false)}
            >
              Batal
            </Button>
            <Button fullWidth isLoading={isCreating || isUpdating}>
              {editingCategory ? "Perbarui" : "Simpan"}
            </Button>
          </div>
        </form>
      </Modal>

      {/* Delete confirmation */}
      <ConfirmDialog
        isOpen={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        onConfirm={confirmDelete}
        title="Hapus Kategori"
        message={`Apakah Anda yakin ingin menghapus kategori "${categoryToDelete?.name}"? Semua produk di kategori ini akan kehilangan kategori.`}
        confirmLabel="Hapus"
        isLoading={isDeleting}
      />
    </div>
  );
}
