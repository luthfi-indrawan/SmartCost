# Product Requirements Document (PRD) - Smart Cost V1

**Versi:** 1.0  
**Status:** Draft Disetujui  
**Target Rilis:** MVP (Minimum Viable Product)

---

## 1. Problem Statement & Goals

### 1.1 Problem Statement

Banyak pemilik usaha mikro, kecil, dan menengah (UMKM) di bidang F&B (restoran/warung) dan retail kecil masih menggunakan pencatatan manual atau sistem kasir konvensional yang kaku. Masalah utama yang dihadapi meliputi:

- **Fragmentasi Data:** Kesulitan memantau performa penjualan harian secara _real-time_ karena data tidak terpusat.
- **Operasional Terhambat:** Ketidakcocokan stok barang akibat sistem kasir yang mengunci/menolak transaksi saat stok di sistem tercatat 0, padahal barang fisik tersedia nyata di toko.
- **Celah Keamanan & Akuntabilitas:** Kurangnya transparansi operasional kasir, sehingga sulit melacak siapa yang bertanggung jawab jika terjadi selisih uang atau pembatalan transaksi (_void/return_).
- **Manajemen Harga Kaku:** Proses pengelolaan harga grosir yang rumit karena memaksa pemilik membuat entri produk baru untuk barang yang sama, yang akhirnya merusak akurasi data stok tunggal.

### 1.2 Product Goals

Smart Cost dikembangkan sebagai aplikasi POS berbasis cloud yang fleksibel, responsif, dan mudah digunakan untuk menyelesaikan masalah operasional UMKM dengan target:

- Menyediakan sistem kasir yang tidak menghambat operasional lapangan (mendukung pemotongan stok langsung dan toleransi stok minus dengan penyesuaian kemudian).
- Meningkatkan akuntabilitas lewat pelacakan aktivitas kasir secara spesifik (_audit trail_).
- Menyederhanakan manajemen harga lewat fitur multi-harga (grosir berbasis kuantitas) dalam satu produk tunggal.

---

## 2. User Personas

### 2.1 Pemilik Toko (Owner) - "Andi, 34 Tahun"

- **Profil:** Pemilik warung makan dan toko retail kecil modern. Jarang berada di toko secara penuh karena harus mengurus suplai barang dan ekspansi bisnis.
- **Kebutuhan:** Melihat laporan pendapatan harian/bulanan dari HP secara _real-time_, memantau produk yang stoknya menipis, dan mengetahui kasir mana yang bertanggung jawab atas setiap transaksi atau tindakan _return_.
- **Frustrasi:** Karyawan sering keliru memberikan kembalian uang tunai, dan sistem lama terlalu rumit untuk menginput skema harga grosir.

### 2.2 Kasir / Staf Toko - "Siti, 22 Tahun"

- **Profil:** Karyawan paruh waktu yang bertugas melayani pelanggan langsung di meja kasir. Menggunakan smartphone pribadi atau tablet toko untuk operasional harian.
- **Kebutuhan:** Proses _checkout_ yang cepat, bisa menahan pesanan (_hold bill_) untuk pelanggan restoran yang makan di tempat, dan bisa mencetak struk langsung ke printer thermal tanpa alur yang rumit.
- **Frustrasi:** Aplikasi kasir sering macet atau menolak transaksi saat stok di sistem terbaca habis, padahal pelanggan sudah mengantre panjang.

---

## 3. Functional Requirements (MoSCoW)

Berikut adalah ruang lingkup fitur Smart Cost V1 yang diklasifikasikan menggunakan metode MoSCoW:

| Kategori        | ID Kebutuhan | Deskripsi Fitur / Modul           | Logika Bisnis & Spesifikasi                                                                                                                                                               |
| :-------------- | :----------- | :-------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Must-Have**   | FR-01        | Multi-User Authentication         | Login terpisah untuk Pemilik dan Kasir. Setiap aktivitas (penjualan, void, return) wajib merekam User ID kasir yang bertugas untuk keperluan audit.                                       |
|                 | FR-02        | Manajemen Produk & Multi-Harga    | Input produk manual dengan opsi kategori fleksibel (Barcode opsional via kamera web/HP). Mendukung harga grosir otomatis dalam satu produk berdasarkan batas minimum kuantitas penjualan. |
|                 | FR-03        | Kasir Fleksibel (Checkout & Hold) | Keranjang belanja dengan opsi "Bayar Langsung" atau "Simpan Pesanan (Hold Bill)" untuk manajemen antrean/meja restoran. Stok langsung terpotong saat pesanan dibuat/dihold.               |
|                 | FR-04        | Sistem Stok Minus Toleran         | Sistem tidak memblokir penjualan jika stok bernilai 0 atau minus di database. Indikator minus akan diberi tanda khusus untuk disesuaikan kemudian (opname) oleh Pemilik.                  |
|                 | FR-05        | Manajemen Return & Void           | Kasir dapat melakukan pembatalan item atau pengembalian barang. Sistem otomatis mengembalikan jumlah stok (+1) ke database dan mencatat log audit pendapatan.                             |
|                 | FR-06        | Dashboard Laporan Pendapatan      | Grafik pendapatan bulanan dan harian yang dapat diakses Pemilik, dilengkapi pelacakan total uang masuk yang melekat pada masing-masing user kasir.                                        |
| **Should-Have** | FR-07        | Cetak Struk Thermal Web API       | Integrasi cetak struk langsung dari browser ke printer thermal menggunakan Web Bluetooth API atau Web USB API.                                                                            |
|                 | FR-08        | Notifikasi Stok Menipis In-App    | Pemberitahuan visual (badge/toast merah) di dashboard internal web Pemilik ketika stok produk berada di bawah ambang batas minimal yang diatur.                                           |
| **Could-Have**  | FR-09        | Template Data UMKM Default        | Opsi bagi pengguna baru untuk memuat master data produk umum (seperti es teh, air mineral, rokok) saat registrasi awal agar tidak perlu input manual dari nol.                            |
| **Won't-Have**  | FR-10        | Offline-First & Hutang Piutang    | Sistem murni berjalan online dan tidak menyediakan fitur pencatatan hutang/bon pelanggan untuk menjaga kesederhanaan MVP awal.                                                            |

---

## 4. User Stories & Acceptance Criteria

### US-01: Manajemen Kasir (Audit Trail)

**Sebagai** _Pemilik Toko_, **saya ingin** setiap kasir memiliki akun login sendiri **agar** saya bisa melacak total pemasukan dan tanggung jawab transaksi berdasarkan aktivitas shift mereka secara transparan.

**Acceptance Criteria:**

- AC1: Sistem menyediakan endpoint registrasi kasir dengan role `cashier` yang hanya dapat diakses oleh user dengan role `owner`.
- AC2: Setiap transaksi yang tercatat di database menyimpan foreign key `cashier_id` yang tidak nullable.
- AC3: Dashboard owner menampilkan total uang masuk per kasir dengan filter rentang tanggal.
- AC4: Sistem menolak transaksi tanpa `cashier_id` yang valid (HTTP 401).

### US-02: Manajemen Harga Grosir

**Sebagai** _Pemilik Toko_, **saya ingin** mengatur harga grosir (misal: beli >= 10 pcs harga otomatis turun) pada satu produk yang sama **agar** saya tidak perlu membuat dua produk terpisah yang merusak kalkulasi stok tunggal.

**Acceptance Criteria:**

- AC1: Tabel `product_prices` mendukung multiple price tiers dengan kolom `min_qty`, `price`, dan `is_wholesale`.
- AC2: Saat produk ditambahkan ke keranjang, sistem otomatis memilih harga tier berdasarkan total kuantitas item tersebut.
- AC3: Jika kuantitas berubah (naik/turun), harga di keranjang otomatis disesuaikan ke tier yang sesuai tanpa refresh halaman.
- AC4: Harga grosir hanya berlaku untuk kuantitas dalam satu transaksi (tidak akumulasi antar transaksi).

### US-03: Proses Checkout dengan Hold Bill

**Sebagai** _Kasir_, **saya ingin** bisa menyimpan pesanan pelanggan (_Hold Bill_) dengan catatan nomor meja atau nama **agar** saya bisa melayani pelanggan lain tanpa menghapus pesanan yang sedang berjalan.

**Acceptance Criteria:**

- AC1: Tombol "Hold Bill" tersedia di halaman kasir saat keranjang tidak kosong.
- AC2: Kasir wajib mengisi field `hold_note` (max 100 karakter) sebelum menyimpan hold bill.
- AC3: Stok produk dalam hold bill langsung terpotong saat tombol ditekan (status `PENDING`).
- AC4: Hold bill muncul di panel "Daftar Antrean" dengan timestamp dan nama kasir.
- AC5: Kasir dapat memanggil kembali hold bill ke keranjang aktif untuk dilanjutkan ke pembayaran.

### US-04: Fleksibilitas Operasional (Stok Minus)

**Sebagai** _Kasir_, **saya ingin** sistem tetap mengizinkan transaksi meskipun stok di sistem tertulis 0 atau minus **karena** di dunia nyata barang fisik tersebut siap dijual dan pelayanan pelanggan tidak boleh terhenti.

**Acceptance Criteria:**

- AC1: Sistem tidak memvalidasi stok > 0 saat produk ditambahkan ke keranjang.
- AC2: Sistem tidak memblokir checkout meskipun stok akhir menjadi minus.
- AC3: Produk dengan stok <= 0 ditandai dengan indikator visual (badge merah/oranye) di halaman kasir.
- AC4: Setiap penjualan yang menyebabkan stok minus tercatat di `stock_alerts` untuk ditampilkan di dashboard owner.

### US-05: Pembatalan Transaksi (Return/Void)

**Sebagai** _Kasir_, **saya ingin** bisa melakukan _return_ pesanan yang batal **agar** nominal tagihan berkurang secara akurat dan stok barang otomatis kembali bertambah di database.

**Acceptance Criteria:**

- AC1: Kasir dapat memilih item dari transaksi aktif atau hold bill untuk di-return.
- AC2: Sistem menampilkan modal konfirmasi sebelum eksekusi return.
- AC3: Setelah return, stok produk bertambah sesuai kuantitas yang di-return.
- AC4: Sistem mencatat log ke tabel `void_logs` dengan field: `transaction_id`, `cashier_id`, `product_id`, `qty_returned`, `reason`, `created_at`.
- AC5: Total tagihan transaksi berkurang sesuai harga item yang di-return pada saat transaksi dibuat.

### US-06: Pemantauan Stok Menipis

**Sebagai** _Pemilik Toko_, **saya ingin** melihat indikator warna merah di dashboard web saat login **sebagai** pengingat bahwa stok barang tertentu sudah menipis dan harus segera dibeli kembali.

**Acceptance Criteria:**

- AC1: Dashboard owner menampilkan section "Critical Alerts" jika ada produk dengan `stock <= min_stock_threshold` atau `stock < 0`.
- AC2: Badge notifikasi merah muncul di navbar dengan jumlah produk bermasalah.
- AC3: Daftar alert menampilkan: nama produk, stok aktual, threshold, dan status (`LOW` atau `MINUS`).
- AC4: Alert otomatis hilang setelah stok diperbarui melebihi threshold.

---

## 5. Non-Functional Requirements

- **Performa Kecepatan:** Aplikasi web harus responsif dan dimuat (_page load_) kurang dari 2 detik pada koneksi internet 4G standar. Proses pencarian produk di halaman kasir harus instan (kurang dari 300ms).
- **Desain Responsif:** Antarmuka (UI/UX) halaman kasir harus dioptimalkan secara khusus untuk layar perangkat mobile (smartphone) dan tablet tanpa mengorbankan fungsionalitas utama kasir.
- **Keamanan Data:** Seluruh komunikasi data wajib berjalan di atas protokol terenkripsi HTTPS. Password pengguna harus di-hash menggunakan algoritma aman (seperti `bcrypt`) sebelum disimpan ke dalam database.
- **Ketersediaan Data & Konkurensi:** Mengingat sistem ini 100% online, arsitektur database harus dirancang mantap untuk menangani konkurensi tinggi agar pemotongan stok tetap konsisten dan aman dari _race condition_ saat diakses oleh beberapa kasir secara bersamaan.
