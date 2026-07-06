import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { motion } from "motion/react";
import { Store, Eye, EyeOff, Lock, Mail } from "lucide-react";
import { useAuth } from "../../hooks/useAuth";
import { Input } from "../../components/ui/Input";
import { Button } from "../../components/ui/Button";

const loginSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
});

export default function LoginPage() {
  const { login, isLoggingIn } = useAuth();
  const [showPassword, setShowPassword] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (data) => {
    login(data);
  };

  return (
    <div className="min-h-screen bg-wa-bg flex items-center justify-center p-4">
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.3 }}
        className="w-full max-w-sm"
      >
        {/* Logo */}
        <div className="text-center mb-8">
          <div className="w-16 h-16 bg-primary-500 rounded-[14px] flex items-center justify-center mx-auto mb-4 shadow-[0_4px_12px_rgba(76,175,80,0.3)]">
            <Store className="w-8 h-8 text-white" />
          </div>
          <h1 className="text-2xl font-bold text-neutral-800">SmartCost</h1>
          <p className="text-sm text-neutral-500 mt-1">
            Aplikasi POS untuk UMKM Indonesia
          </p>
        </div>

        {/* Form */}
        <div className="bg-white rounded-[14px] p-6 shadow-[0_1px_3px_rgba(0,0,0,0.05)] border border-neutral-200">
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <Input
              label="Email"
              type="email"
              placeholder="nama@email.com"
              icon={Mail}
              error={errors.email?.message}
              {...register("email")}
            />

            <div className="relative">
              <Input
                label="Password"
                type={showPassword ? "text" : "password"}
                placeholder="••••••••"
                icon={Lock}
                error={errors.password?.message}
                {...register("password")}
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-3 top-[38px] p-1 text-neutral-400 hover:text-neutral-600"
              >
                {showPassword ? (
                  <EyeOff className="w-4 h-4" />
                ) : (
                  <Eye className="w-4 h-4" />
                )}
              </button>
            </div>

            <Button
              type="submit"
              fullWidth
              isLoading={isLoggingIn}
              className="mt-2"
            >
              Masuk
            </Button>
          </form>
        </div>

        {/* Footer */}
        <p className="text-center text-xs text-neutral-400 mt-6">
          © 2026 SmartCost. All rights reserved.
        </p>
      </motion.div>
    </div>
  );
}
