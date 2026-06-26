# User Flow & Wireframe Mapping - Smart Cost V1

**Versi:** 1.0
**Status:** Draft Disetujui  
**Target Rilis:** MVP (Minimum Viable Product)

---

## 1. Core User Journeys & Logic Gates (Mermaid Diagrams)

Aplikasi ini memiliki dua peran (_role_) dengan hak akses berbeda: **Pemilik (Owner)** dan **Kasir (Operator)**. Berikut adalah visualisasi alur perjalanan pengguna serta logika percabangannya.

### 1.1 Alur Otentikasi & Gerbang Masuk (Authentication Flow)

Alur ini memastikan pengguna diarahkan ke antarmuka yang tepat sesuai dengan hak akses mereka, serta memblokir akses ilegal jika sesi belum tervalidasi.

```mermaid
graph TD
    A[Mulai] --> B[Halaman Login]
    B --> C[Input Kredensial: Email/Username & Password]
    C --> D{Cek Kredensial?}

    D -- Gagal --> E[Tampilkan Pesan Error: Kredensial Salah]
    E --> B

    D -- Sukses --> F{Periksa Role User?}
    F -- Owner --> G[Diarahkan ke /dashboard<br>Akses Penuh & Fitur Laporan]
    F -- Kasir --> H[Diarahkan ke /kasir<br>Akses Terbatas Teroptimasi Mobile]

    style G fill:#d4edda,stroke:#28a745,stroke-width:2px
    style H fill:#d1ecf1,stroke:#17a2b8,stroke-width:2px
    style E fill:#f8d7da,stroke:#dc3545,stroke-width:2px
```

- **Logic Gates (Aturan Proteksi):**
  - Jika user mencoba mengakses `/dashboard` atau `/kasir` tanpa token sesi (JWT/Session Cookie), sistem wajib melakukan _intercept_ dan melempar kembali ke halaman `/login`.
  - Jika user dengan role _Kasir_ mencoba mengetik URL `/dashboard` atau `/products` secara manual, sistem akan memblokir akses dan menampilkan halaman _403 Forbidden_.

---

### 1.2 Alur Kerja Kasir: Transaksi & Pengkondisian Harga Grosir

Alur ini memetakan bagaimana Kasir memasukkan produk ke keranjang, bagaimana sistem menghitung skema grosir secara otomatis, hingga opsi pembayaran (_Direct Pay_ atau _Hold Bill_).

```mermaid
graph TD
    A[Kasir Buka Halaman /kasir] --> B[Pilih Produk / Scan Barcode Opsional]
    B --> C[Produk Masuk Keranjang Belanja]
    C --> D{Kuantitas Produk >= Batas Minimum Grosir?}

    D -- Ya --> E[Ubah ke Harga Grosir Otomatis]
    D -- Tidak --> F[Gunakan Harga Reguler]

    E --> G[Pilih Metode Penyelesaian]
    F --> G

    G -- Opsi A: Bayar Langsung --> H[Kasir Input Uang Tunai]
    H --> I[Kurangi Stok Permanen]
    I --> J[Transaksi SUCCESS & Rekerd User ID Kasir]
    J --> K[Cetak Struk via Web Thermal API]

    G -- Opsi B: Simpan Pesanan / Hold Bill --> L[Kasir Input Catatan: Nomor Meja/Nama]
    L --> M[Simpan Transaksi dengan Status PENDING]
    M --> N[Stok Dipotong Langsung Saat Itu Juga]
    N --> O[Keranjang Kembali Kosong / Siap untuk Antrean Baru]

    style J fill:#d4edda,stroke:#28a745,stroke-width:2px
    style O fill:#fff3cd,stroke:#ffc107,stroke-width:2px
```

---

### 1.3 Alur Pembatalan / Pengembalian Barang (Return & Void)

Mencegah kerugian dan kecurangan kasir (_fraud_) dengan mencatat log audit lengkap, tanpa menghentikan operasional toko saat pemilik tidak ada di tempat.

```mermaid
graph TD
    A[Kasir Buka Riwayat Transaksi / Hold Bill Aktif] --> B[Pilih Item Produk yang Ingin Dibatalkan]
    B --> C[Klik Tombol Return Item]
    C --> D{Konfirmasi Tindakan?}

    D -- Batalkan --> E[Kembali ke Halaman Sebelumnya]

    D -- Ya, Setuju --> F[Kurangi Total Tagihan Berjalan]
    F --> G[Kembalikan Kuantitas Stok ke Database: +1]
    G --> H[Catat Rekord ke void_logs:<br>Transaksi ID, User ID Kasir, Waktu, & Produk]
    H --> I[Selesai / Update Tampilan Struk Baru]

    style H fill:#f8d7da,stroke:#dc3545,stroke-width:2px
```

---

### 1.4 Alur Pemilik: Pemantauan Dashboard & Notifikasi Stok Minus

Sistem memindai stok secara otomatis untuk memberikan peringatan dini kepada pemilik usaha.

```mermaid
graph TD
    A[Owner Login ke /dashboard] --> B{Sistem Memindai Tabel Produk}

    B -- Ada Produk dengan Stok <= Batas Min / Minus (<0) --> C[Munculkan Badge Notifikasi Merah di Navbar]
    C --> D[Tampilkan Daftar Critical Alerts di Halaman Utama Dashboard]

    B -- Stok Aman --> E[Tampilkan Ringkasan Dashboard Normal]

    D --> F[Owner Buka /products untuk Update Stok/Harga]
    E --> G[Owner Buka /reports untuk Cek Grafik & Total Uang Masuk per Kasir]
```

---

## 2. Interface Structures (ASCII Layouts)

Desain antarmuka dirancang responsif agar optimal saat diakses menggunakan tablet toko maupun smartphone kasir.

### 2.1 Tata Letak Halaman Kasir (Responsif / Tablet & Desktop View)

```
+-----------------------------------------------------------------------------+
|  [SMART COST]                                         (User: Siti - KASIR)  |
+------------------------------------+----------------------------------------+
| CARI PRODUK: [ Cari nama/barcode ] | KERANJANG BELANJA      [Pesanan Baru]  |
+------------------------------------+----------------------------------------+
| [ KATEGORI: MAKANAN | MINUMAN ]    | 1. Nasi Goreng Ayam      (x2)  60.000  |
|                                    | 2. Es Teh Manis          (x2)   8.000  |
| +----------------+ +-------------+ | 3. Teh Kotak (Grosir)*  (x10)  25.000  |
| | Nasi Goreng Ap | | Mie Goreng  | |                                        |
| | Rp30.000       | | Rp25.000    | | *Grosir Aktif (Min. 10 pcs)            |
| +----------------+ +-------------+ |                                        |
| +----------------+ +-------------+ |----------------------------------------|
| | Es Teh Manis   | | Teh Kotak  | | TOTAL BELANJA:               Rp93.000  |
| | Rp4.000        | | Rp3.000    | |----------------------------------------|
| +----------------+ +-------------+ | [ TAHAN / HOLD BILL ]  [ BAYAR TUNAI ] |
+------------------------------------+----------------------------------------+
| [Daftar Antrean Hold (3)]          | [Riwayat Transaksi Shift Ini]          |
+------------------------------------+----------------------------------------+
```

### 2.2 Tata Letak Dashboard Pemilik (Owner View - Mobile-Optimized)

```
+---------------------------------------+
| [=] SMART COST              [ (1) 🔔 ]| <-- Notifikasi Stok Menipis/Minus
+---------------------------------------+
| SELAMAT DATANG, ANDI (OWNER)          |
|                                       |
| PENDAPATAN HARI INI:                  |
| Rp 1.545.000                          |
|                                       |
| [ Grafik Pendapatan Bulanan (Canvas) ]|
|  Jan  [██████]                        |
|  Feb  [███████████]                   |
|                                       |
|---------------------------------------|
| PERFORMA KASIR (TOTAL UANG MASUK):    |
| - Siti (Shift Pagi) : Rp  945.000     |
| - Budi (Shift Malam): Rp  600.000     |
|                                       |
|---------------------------------------|
| PENGINGAT STOK (CRITICAL ALERTS):     |
| ⚠️ Es Teh Manis  : Stok MINUS (-3)    |
| ⚠️ Teh Kotak     : Stok Menipis (2)   |
+---------------------------------------+
| [ Dashboard ] [ Produk ] [ Laporan ]  | <-- Bottom Navigation di HP
+---------------------------------------+
```

---

## 3. Celah Logika yang Berhasil Diminimalkan (Loophole Prevention)

1. **Konkurensi Pengurangan Stok:** Stok langsung dikunci saat kasir menekan tombol _Hold Bill_. Ini mencegah kasir lain di perangkat berbeda menjual barang fisik yang sama yang saat itu sedang disajikan di meja pelanggan.
2. **Validasi Stok Minus Kontrol:** Sistem mengizinkan stok menjadi minus di database agar transaksi tidak macet di depan pelanggan, namun _loophole_ kerugian bisnis ditutup dengan langsung memaksa sistem memunculkan peringatan berwarna merah di dashboard pemilik saat itu juga agar segera dilakukan opname fisik.
3. **Jejak Digital Pembatalan:** Tombol _Return_ tidak memerlukan otorisasi PIN fisik pemilik di tempat (karena pemilik sering berada di luar toko), namun celah kecurangan kasir ditutup dengan sistem pencatatan log audit otomatis (`void_logs`) yang mengikat tindakan tersebut pada User ID kasir yang aktif, sehingga pemilik bisa memeriksa kejanggalan kapan saja dari HP mereka.
