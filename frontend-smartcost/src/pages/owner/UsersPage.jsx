import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Plus, Pencil, Trash2, User, Phone, Mail, Power } from "lucide-react";
import { useUsers } from "../../hooks/useUsers";
import { Button } from "../../components/ui/Button";
import { Input } from "../../components/ui/Input";
import { Card } from "../../components/ui/Card";
import { Modal } from "../../components/ui/Modal";
import { ConfirmDialog } from "../../components/ui/ConfirmDialog";
import { EmptyState } from "../../components/ui/EmptyState";
import { ListSkeleton } from "../../components/ui/Skeleton";
import { formatRupiah } from "../../lib/constants";
import { cn } from "../../lib/utils";

const userSchema = z.object({
  name: z.string().min(1, "Nama wajib diisi"),
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
  phone: z
    .string()
    .regex(/^08\d{8,11}$/, "Nomor HP tidak valid")
    .optional()
    .or(z.literal("")),
  role: z.literal("cashier"),
});

export default function UsersPage() {
  const [showModal, setShowModal] = useState(false);
  const [editingUser, setEditingUser] = useState(null);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [userToDelete, setUserToDelete] = useState(null);

  const {
    users,
    isLoading,
    createUser,
    updateUser,
    deleteUser,
    isCreating,
    isUpdating,
    isDeleting,
  } = useUsers();

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(userSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      phone: "",
      role: "cashier",
    },
  });

  const openCreate = () => {
    setEditingUser(null);
    reset({ name: "", email: "", password: "", phone: "", role: "cashier" });
    setShowModal(true);
  };

  const openEdit = (user) => {
    setEditingUser(user);
    setValue("name", user.name);
    setValue("email", user.email);
    setValue("phone", user.phone || "");
    setValue("role", "cashier");
    setShowModal(true);
  };

  const handleDelete = (user) => {
    setUserToDelete(user);
    setShowDeleteConfirm(true);
  };

  const onSubmit = async (data) => {
    try {
      if (editingUser) {
        const { password, ...updateData } = data;
        await updateUser({ id: editingUser.id, data: updateData });
      } else {
        await createUser(data);
      }
      setShowModal(false);
      reset();
    } catch (error) {
      // Error handled by mutation hook
    }
  };

  const confirmDelete = () => {
    if (userToDelete) {
      deleteUser(userToDelete.id);
      setShowDeleteConfirm(false);
      setUserToDelete(null);
    }
  };

  return (
    <div className="p-4 space-y-4 pb-24 lg:pb-6">
      <div className="flex items-center justify-between">
        <h2 className="text-lg font-semibold text-neutral-800">Kelola Kasir</h2>
        <Button onClick={openCreate} size="sm">
          <Plus className="w-4 h-4" />
          Tambah Kasir
        </Button>
      </div>

      {isLoading ? (
        <ListSkeleton count={4} />
      ) : users.length === 0 ? (
        <EmptyState
          title="Belum ada kasir"
          description="Daftarkan kasir untuk mulai beroperasi."
          icon="empty"
          action={
            <Button onClick={openCreate} size="sm">
              <Plus className="w-4 h-4" />
              Daftar Kasir
            </Button>
          }
        />
      ) : (
        <div className="space-y-3">
          {users.map((user) => (
            <Card
              key={user.id}
              padding="normal"
              className="flex items-center gap-4"
            >
              <div
                className={cn(
                  "w-12 h-12 rounded-full flex items-center justify-center flex-shrink-0",
                  user.is_active ? "bg-primary-100" : "bg-neutral-100",
                )}
              >
                <User
                  className={cn(
                    "w-6 h-6",
                    user.is_active ? "text-primary-600" : "text-neutral-400",
                  )}
                />
              </div>
              <div className="flex-1 min-w-0">
                <div className="flex items-center gap-2">
                  <h4 className="text-sm font-semibold text-neutral-800">
                    {user.name}
                  </h4>
                  {!user.is_active && (
                    <span className="px-2 py-0.5 bg-neutral-100 text-neutral-500 text-[10px] font-medium rounded-full">
                      Nonaktif
                    </span>
                  )}
                </div>
                <div className="flex items-center gap-3 text-xs text-neutral-500 mt-0.5">
                  <span className="flex items-center gap-1">
                    <Mail className="w-3 h-3" />
                    {user.email}
                  </span>
                  {user.phone && (
                    <span className="flex items-center gap-1">
                      <Phone className="w-3 h-3" />
                      {user.phone}
                    </span>
                  )}
                </div>
                <div className="flex items-center gap-3 mt-1.5">
                  <span className="text-xs font-medium text-primary-600">
                    {formatRupiah(user.total_sales || 0)}
                  </span>
                  <span className="text-xs text-neutral-400">
                    {user.transaction_count || 0} transaksi
                  </span>
                </div>
              </div>
              <div className="flex items-center gap-1">
                <button
                  onClick={() => openEdit(user)}
                  className="p-2 rounded-lg hover:bg-neutral-100 text-neutral-400 hover:text-primary-600"
                >
                  <Pencil className="w-4 h-4" />
                </button>
                <button
                  onClick={() => handleDelete(user)}
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
        title={editingUser ? "Edit Kasir" : "Tambah Kasir"}
        size="sm"
      >
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <Input
            label="Nama Lengkap"
            placeholder="Contoh: Siti Aminah"
            error={errors.name?.message}
            {...register("name")}
          />
          <Input
            label="Email"
            type="email"
            placeholder="siti@warungku.com"
            error={errors.email?.message}
            {...register("email")}
          />
          {!editingUser && (
            <Input
              label="Password"
              type="password"
              placeholder="Minimal 6 karakter"
              error={errors.password?.message}
              {...register("password")}
            />
          )}
          <Input
            label="Nomor HP (Opsional)"
            placeholder="081234567890"
            error={errors.phone?.message}
            {...register("phone")}
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
              {editingUser ? "Perbarui" : "Simpan"}
            </Button>
          </div>
        </form>
      </Modal>

      {/* Delete confirmation */}
      <ConfirmDialog
        isOpen={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        onConfirm={confirmDelete}
        title="Nonaktifkan Kasir"
        message={`Apakah Anda yakin ingin menonaktifkan "${userToDelete?.name}"? Kasir ini tidak akan bisa login lagi.`}
        confirmLabel="Nonaktifkan"
        isLoading={isDeleting}
      />
    </div>
  );
}
