# QA Checklist (QA_CHECKLIST) - Smart Cost V1

**Versi:** 1.0  
**Status:** Draft Disetujui  
**Tester:** QA Engineer  
**Environment:** Staging (`https://staging.smartcost.app`)

---

## 1. Functional Verification Matrix

### 1.1 Authentication (FR-01)

| ID    | User Story | Test Step                              | Expected Result                                                | Status |
| :---- | :--------- | :------------------------------------- | :------------------------------------------------------------- | :----- |
| FV-01 | US-01      | Owner login dengan kredensial valid    | Redirect ke `/dashboard`, token tersimpan di cookie            | ⬜     |
| FV-02 | US-01      | Kasir login dengan kredensial valid    | Redirect ke `/kasir`, token tersimpan di cookie                | ⬜     |
| FV-03 | US-01      | Login dengan password salah            | Tampil error "Email atau password salah", tetap di `/login`    | ⬜     |
| FV-04 | US-01      | Akses `/dashboard` tanpa token         | Redirect ke `/login`                                           | ⬜     |
| FV-05 | US-01      | Kasir akses `/dashboard` manual        | Tampil halaman 403 Forbidden                                   | ⬜     |
| FV-06 | US-01      | Owner registrasi kasir baru            | Kasir tersimpan di DB dengan role `cashier`, password ter-hash | ⬜     |
| FV-07 | US-01      | Owner registrasi dengan email duplikat | Tampil error 409 "Email sudah terdaftar"                       | ⬜     |
| FV-08 | US-01      | Logout                                 | Cookie token dihapus, redirect ke `/login`                     | ⬜     |
| FV-09 | US-01      | Token expired                          | Auto-redirect ke `/login` atau refresh token                   | ⬜     |

### 1.2 Product Management (FR-02)

| ID    | User Story | Test Step                              | Expected Result                                            | Status |
| :---- | :--------- | :------------------------------------- | :--------------------------------------------------------- | :----- |
| FV-10 | US-02      | Owner tambah produk dengan price tiers | Produk tersimpan, price tiers terkait tercatat             | ⬜     |
| FV-11 | US-02      | Owner edit harga grosir produk         | Harga tier terupdate, transaksi lama tidak berubah         | ⬜     |
| FV-12 | US-02      | Owner hapus produk (soft delete)       | Produk `is_active = false`, tidak muncul di kasir          | ⬜     |
| FV-13 | US-02      | Search produk di halaman owner         | Hasil pencarian sesuai keyword (nama/SKU/barcode)          | ⬜     |
| FV-14 | US-02      | Filter produk by kategori              | Hanya produk kategori terpilih yang muncul                 | ⬜     |
| FV-15 | US-02      | Filter produk by stock status          | Produk dengan status SAFE/LOW/MINUS terfilter dengan benar | ⬜     |
| FV-16 | US-02      | Tambah produk dengan SKU duplikat      | Error 400 "SKU sudah digunakan"                            | ⬜     |
| FV-17 | US-02      | Tambah price tier dengan min_qty <= 1  | Error validasi "min_qty harus lebih besar dari 1"          | ⬜     |

### 1.3 Cashier Checkout (FR-03, FR-04)

| ID    | User Story | Test Step                                   | Expected Result                                       | Status |
| :---- | :--------- | :------------------------------------------ | :---------------------------------------------------- | :----- |
| FV-18 | US-03      | Kasir tambah produk ke keranjang            | Produk muncul di cart dengan qty 1                    | ⬜     |
| FV-19 | US-03      | Kasir tambah qty produk di keranjang        | Qty bertambah, subtotal terupdate                     | ⬜     |
| FV-20 | US-03      | Kasir hapus item dari keranjang             | Item hilang, total berkurang                          | ⬜     |
| FV-21 | US-02      | Kasir tambah produk grosir (qty >= min_qty) | Harga otomatis berubah ke harga grosir                | ⬜     |
| FV-22 | US-02      | Kasir kurangi qty grosir di bawah threshold | Harga kembali ke base_price                           | ⬜     |
| FV-23 | US-03      | Kasir simpan hold bill                      | Transaksi tersimpan status PENDING, stok terpotong    | ⬜     |
| FV-24 | US-03      | Kasir panggil kembali hold bill             | Hold bill muncul di keranjang, bisa dilanjutkan bayar | ⬜     |
| FV-25 | US-03      | Kasir bayar transaksi langsung              | Transaksi status COMPLETED, struk muncul              | ⬜     |
| FV-26 | US-04      | Kasir checkout produk dengan stok 0         | Transaksi berhasil, stok menjadi minus                | ⬜     |
| FV-27 | US-04      | Kasir checkout produk dengan stok minus     | Transaksi berhasil, stok semakin minus                | ⬜     |
| FV-28 | US-04      | Indikator stok minus di halaman kasir       | Badge merah muncul pada produk stok <= 0              | ⬜     |
| FV-29 | US-03      | Kasir input uang tunai dan hitung kembalian | Kembalian dihitung otomatis: amount_paid - total      | ⬜     |
| FV-30 | US-03      | Kasir bayar dengan uang kurang              | Error validasi "Uang tidak cukup"                     | ⬜     |

### 1.4 Return & Void (FR-05)

| ID    | User Story | Test Step                                      | Expected Result                                            | Status |
| :---- | :--------- | :--------------------------------------------- | :--------------------------------------------------------- | :----- |
| FV-31 | US-05      | Kasir return item dari transaksi completed     | Stok produk bertambah, total transaksi berkurang           | ⬜     |
| FV-32 | US-05      | Kasir return item dengan qty melebihi qty asli | Error validasi "Qty return melebihi qty transaksi"         | ⬜     |
| FV-33 | US-05      | Kasir return item tanpa alasan                 | Error validasi "Alasan wajib diisi"                        | ⬜     |
| FV-34 | US-05      | Kasir batalkan hold bill                       | Transaksi status CANCELLED, stok dikembalikan              | ⬜     |
| FV-35 | US-05      | Cek void_logs setelah return                   | Log tercatat dengan cashier_id, transaction_id, product_id | ⬜     |
| FV-36 | US-05      | Owner lihat void_logs di dashboard             | Daftar return terurut by created_at desc                   | ⬜     |

### 1.5 Dashboard & Reports (FR-06, FR-08)

| ID    | User Story | Test Step                             | Expected Result                                              | Status |
| :---- | :--------- | :------------------------------------ | :----------------------------------------------------------- | ------ |
| FV-37 | US-06      | Owner lihat pendapatan hari ini       | Nilai sesuai sum transaksi COMPLETED hari ini                | ⬜     |
| FV-38 | US-06      | Owner lihat grafik pendapatan bulanan | Grafik menampilkan 30 hari terakhir                          | ⬜     |
| FV-39 | US-06      | Owner filter report by kasir          | Hanya transaksi kasir terpilih yang dihitung                 | ⬜     |
| FV-40 | US-06      | Owner filter report by tanggal        | Rentang tanggal sesuai filter                                | ⬜     |
| FV-41 | US-06      | Total uang masuk per kasir akurat     | Sum transaksi COMPLETED per cashier_id                       | ⬜     |
| FV-42 | US-06      | Badge notifikasi stok muncul          | Badge merah dengan jumlah produk bermasalah                  | ⬜     |
| FV-43 | US-06      | Daftar critical alerts akurat         | Produk dengan stok <= threshold atau < 0 terdaftar           | ⬜     |
| FV-44 | US-06      | Alert hilang setelah stok diperbarui  | Produk tidak lagi muncul di critical alerts                  | ⬜     |
| FV-45 | US-06      | Owner update stok produk              | Stok terupdate, alert status berubah jika melebihi threshold | ⬜     |

---

## 2. Test Case Scenarios Hierarchy

### 2.1 Happy Paths

| ID    | Scenario                          | Precondition                        | Steps                                                                                                  | Expected Result                                          |
| :---- | :-------------------------------- | :---------------------------------- | :----------------------------------------------------------------------------------------------------- | :------------------------------------------------------- |
| HP-01 | Kasir melakukan transaksi lengkap | Kasir login, produk tersedia        | 1. Pilih produk<br>2. Tambah qty<br>3. Klik Bayar<br>4. Input uang tunai<br>5. Konfirmasi              | Transaksi COMPLETED, stok terpotong, struk tercetak      |
| HP-02 | Owner melihat laporan harian      | Owner login, ada transaksi hari ini | 1. Buka dashboard<br>2. Lihat card "Pendapatan Hari Ini"                                               | Nilai sesuai total transaksi hari ini                    |
| HP-03 | Kasir menggunakan hold bill       | Kasir login, keranjang tidak kosong | 1. Tambah produk ke keranjang<br>2. Klik Hold Bill<br>3. Input catatan meja<br>4. Simpan               | Transaksi PENDING, stok terpotong, muncul di daftar hold |
| HP-04 | Owner menambah produk grosir      | Owner login                         | 1. Buka halaman produk<br>2. Klik Tambah<br>3. Isi detail + 2 price tiers<br>4. Simpan                 | Produk tersimpan dengan price tiers                      |
| HP-05 | Kasir melakukan return            | Ada transaksi COMPLETED             | 1. Buka riwayat transaksi<br>2. Pilih item<br>3. Klik Return<br>4. Input qty & alasan<br>5. Konfirmasi | Stok bertambah, log tercatat, total berkurang            |

### 2.2 Edge Cases

| ID    | Scenario                     | Precondition                                | Steps                                        | Expected Result                                                  |
| :---- | :--------------------------- | :------------------------------------------ | :------------------------------------------- | :--------------------------------------------------------------- |
| EC-01 | Transaksi dengan 50+ item    | Keranjang memiliki 50 item berbeda          | 1. Tambah 50 produk<br>2. Checkout           | Transaksi berhasil, performa tetap < 2 detik                     |
| EC-02 | Stok minus ekstrem (-999)    | Produk stok -999                            | 1. Tambah produk ke keranjang<br>2. Checkout | Transaksi berhasil, stok menjadi -1000                           |
| EC-03 | Hold bill lebih dari 24 jam  | Hold bill dibuat kemarin                    | 1. Buka daftar hold<br>2. Panggil kembali    | Hold bill masih bisa dipanggil, timer menunjukkan > 24 jam       |
| EC-04 | Price tier dengan harga sama | Produk memiliki 2 tier dengan harga identik | 1. Tambah produk dengan qty >= min_qty       | Harga yang dipilih adalah tier dengan min_qty terbesar           |
| EC-05 | Transaksi dengan total Rp 0  | Semua item di-return                        | 1. Checkout produk<br>2. Return semua item   | Transaksi total menjadi 0, status tetap COMPLETED                |
| EC-06 | Nama produk 150 karakter     | Produk dengan nama panjang                  | 1. Tampilkan di grid kasir                   | Nama ter-truncate dengan ellipsis, tooltip menampilkan full name |
| EC-07 | 1000+ transaksi dalam sehari | Load test data                              | 1. Buka dashboard owner<br>2. Lihat grafik   | Grafik tetap responsive, tidak lag                               |

### 2.3 Negative Paths

| ID    | Scenario                   | Precondition                          | Steps                                             | Expected Result                                             |
| :---- | :------------------------- | :------------------------------------ | :------------------------------------------------ | :---------------------------------------------------------- |
| NP-01 | SQL Injection di search    | Kasir di halaman produk               | 1. Input `' OR 1=1 --` di search                  | Tidak ada error 500, hasil kosong atau default              |
| NP-02 | XSS di nama produk         | Owner membuat produk                  | 1. Input `<script>alert('xss')</script>` di nama  | Script tidak dieksekusi, nama tersimpan sebagai plain text  |
| NP-03 | Concurrent stock deduction | 2 kasir checkout produk sama (stok 1) | 1. Kasir A dan B checkout bersamaan               | Salah satu berhasil, satu gagal dengan error race condition |
| NP-04 | JWT token manipulation     | User memodifikasi token               | 1. Ubah payload JWT di client<br>2. Kirim request | Server reject dengan 401, tidak ada akses ilegal            |
| NP-05 | Rate limit login           | Brute force attack                    | 1. Kirim 100 request login dalam 10 detik         | IP di-block selama 15 menit, response 429                   |
| NP-06 | Upload file malicious      | Upload avatar                         | 1. Upload file `.exe` sebagai avatar              | Rejected, hanya menerima `.jpg`, `.png`, `.webp`            |
| NP-07 | Access control bypass      | Kasir mencoba API owner               | 1. Kasir kirim POST /products dengan token valid  | Response 403 Forbidden                                      |
| NP-08 | Negative price input       | Input harga produk                    | 1. Input harga -1000                              | Validasi error "Harga harus positif"                        |
| NP-09 | Oversized payload          | Request body > 10MB                   | 1. Kirim request dengan body 50MB                 | Server reject dengan 413 Payload Too Large                  |
| NP-10 | CSRF attack                | User di halaman berbahaya             | 1. Kirim request dari origin lain                 | Rejected oleh CORS middleware                               |

---

## 3. Cross-Browser & Responsiveness Matrix

### 3.1 Browser Compatibility

| Browser          | Version | OS           | Kasir Page | Dashboard | Auth | Status |
| :--------------- | :------ | :----------- | :--------- | :-------- | :--- | :----- |
| Chrome           | 120+    | Windows 11   | ⬜         | ⬜        | ⬜   |        |
| Chrome           | 120+    | Android 14   | ⬜         | ⬜        | ⬜   |        |
| Firefox          | 121+    | Windows 11   | ⬜         | ⬜        | ⬜   |        |
| Firefox          | 121+    | Android 14   | ⬜         | ⬜        | ⬜   |        |
| Safari           | 17+     | macOS Sonoma | ⬜         | ⬜        | ⬜   |        |
| Safari           | 17+     | iOS 17       | ⬜         | ⬜        | ⬜   |        |
| Edge             | 120+    | Windows 11   | ⬜         | ⬜        | ⬜   |        |
| Samsung Internet | 23+     | Android 14   | ⬜         | ⬜        | ⬜   |        |

### 3.2 Device Breakpoints

| Breakpoint      | Device    | Kasir Layout                 | Dashboard Layout             | Touch Targets | Status |
| :-------------- | :-------- | :--------------------------- | :--------------------------- | :------------ | :----- |
| **xs** (320px)  | iPhone SE | Single column, bottom nav    | Single column, stacked cards | >= 44x44px    | ⬜     |
| **sm** (375px)  | iPhone 14 | Single column, bottom nav    | Single column, stacked cards | >= 44x44px    | ⬜     |
| **md** (768px)  | iPad Mini | Two column (products + cart) | Two column, sidebar          | >= 44x44px    | ⬜     |
| **lg** (1024px) | iPad Pro  | Two column, sidebar nav      | Three column, sidebar        | >= 40x40px    | ⬜     |
| **xl** (1440px) | Desktop   | Three column, sidebar        | Full dashboard, sidebar      | Mouse hover   | ⬜     |

### 3.3 Performance Benchmarks

| Metric                         | Target  | Tool            | Status |
| :----------------------------- | :------ | :-------------- | :----- |
| First Contentful Paint (FCP)   | < 1.5s  | Lighthouse      | ⬜     |
| Largest Contentful Paint (LCP) | < 2.5s  | Lighthouse      | ⬜     |
| Time to Interactive (TTI)      | < 3.5s  | Lighthouse      | ⬜     |
| Cumulative Layout Shift (CLS)  | < 0.1   | Lighthouse      | ⬜     |
| API Response Time (p95)        | < 200ms | k6 / Artillery  | ⬜     |
| Product Search Response        | < 300ms | Chrome DevTools | ⬜     |
| Concurrent Users Support       | 50+     | k6 Load Test    | ⬜     |

### 3.4 Accessibility Checklist

| Check                                      | Standard         | Status |
| :----------------------------------------- | :--------------- | :----- |
| Semantik HTML (header, nav, main, section) | WCAG 2.1 AA      | ⬜     |
| ARIA labels pada interactive elements      | WCAG 2.1 AA      | ⬜     |
| Keyboard navigation (Tab, Enter, Escape)   | WCAG 2.1 AA      | ⬜     |
| Color contrast ratio >= 4.5:1              | WCAG 2.1 AA      | ⬜     |
| Focus indicators visible                   | WCAG 2.1 AA      | ⬜     |
| Alt text pada images                       | WCAG 2.1 AA      | ⬜     |
| Screen reader compatible                   | NVDA / VoiceOver | ⬜     |

---

## 4. Regression Testing

### 4.1 Pre-Release Regression Suite

- [ ] Semua happy path scenarios (HP-01 s/d HP-05) passing
- [ ] Semua edge case scenarios (EC-01 s/d EC-07) passing
- [ ] Semua negative path scenarios (NP-01 s/d NP-10) handled correctly
- [ ] Cross-browser matrix (8 browsers) passing
- [ ] Responsiveness matrix (5 breakpoints) passing
- [ ] Performance benchmarks (6 metrics) memenuhi target
- [ ] Accessibility checklist (7 items) passing
- [ ] Security audit (OWASP Top 10) passing
- [ ] Database migration up/down berfungsi
- [ ] Backup dan restore procedure tested

### 4.2 Sign-Off Criteria

| Stakeholder   | Sign-Off Required         | Date |
| :------------ | :------------------------ | :--- |
| Product Owner | Fitur sesuai PRD          | ⬜   |
| Tech Lead     | Arsitektur & code quality | ⬜   |
| QA Lead       | Semua test case passing   | ⬜   |
| Security Lead | Security audit clear      | ⬜   |
| DevOps        | Deployment pipeline green | ⬜   |
