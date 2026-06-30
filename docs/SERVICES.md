# Services API Contract (SERVICES) - Smart Cost V1

**Versi:** 1.0
**Base URL:** `https://api.smartcost.app/api/v1`
**Content-Type:** `application/json`
**Authentication:** JWT Bearer via `httpOnly` Cookie (`access_token`)

---

## Daftar Isi

1. [Global Configuration Network](#1-global-configuration-network)
2. [Authentication Module](#2-authentication-module)
3. [Users Module (Owner Only)](#3-users-module-owner-only)
4. [Products Module](#4-products-module)
5. [Categories Module](#5-categories-module)
6. [Transactions Module (Cashier Only)](#6-transactions-module-cashier-only)
7. [Reports Module (Owner Only)](#7-reports-module-owner-only)
8. [Void Logs Module (Owner Only)](#8-void-logs-module-owner-only)
9. [Appendix: Business Logic & Rules](#9-appendix-business-logic--rules)

---

## 1. Global Configuration Network

### 1.1 Base URLs

| Environment | Base URL                                   |
| :---------- | :----------------------------------------- |
| Development | `http://localhost:8080/api/v1`             |
| Staging     | `https://staging-api.smartcost.app/api/v1` |
| Production  | `https://api.smartcost.app/api/v1`         |

### 1.2 Global Headers

| Header         | Value              | Required | Description                          |
| :------------- | :----------------- | :------- | :----------------------------------- |
| `Content-Type` | `application/json` | Yes      | Request body format                  |
| `Accept`       | `application/json` | Yes      | Response format                      |
| `X-Request-ID` | UUID v4            | No       | Trace ID untuk logging dan debugging |

### 1.3 Pagination & metadatadata Standards

Semua endpoint `GET` yang mengembalikan list WAJIB menggunakan format response berikut:

```json
{
  "code": 200,
  "message": "successfully gether user list",
  "result": {
    "data": [...],
    "metadatadata": {
      "pagination": {
        "current_page": 1,
        "page_size": 20,
        "total_pages": 5,
        "total_items": 98,
        "has_next_page": true,
        "has_prev_page": false
      },
      "sort": {
        "field": "created_at",
        "direction": "desc"
      },
      "filters": {
        "search": "nasi",
        "category_id": "cat_001",
        "stock_status": "all"
      }
    }
  }
}
```

**Query Parameters untuk List Endpoints:**

| Parameter    | Type    | Default      | Description                         |
| :----------- | :------ | :----------- | :---------------------------------- |
| `page`       | integer | 1            | Halaman aktif                       |
| `page_size`  | integer | 20           | Item per halaman (max 100)          |
| `sort_by`    | string  | `created_at` | Field sorting                       |
| `sort_order` | string  | `desc`       | `asc` atau `desc`                   |
| `search`     | string  | -            | Pencarian global (case-insensitive) |

---

## 2. Authentication Module

### 2.1 POST /auth/login

Login untuk Owner dan Kasir.

**Request:**

```json
{
  "email": "andi@warungku.com",
  "password": "password123"
}
```

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "user": {
      "id": "usr_a1b2c3d4",
      "name": "Andi Wijaya",
      "email": "andi@warungku.com",
      "role": "owner",
      "avatar_url": "https://cdn.smartcost.app/avatars/usr_a1b2c3d4.jpg",
      "created_at": "2024-01-15T08:00:00Z"
    },
    "session": {
      "access_token_expires_at": "2024-06-21T08:00:00Z",
      "refresh_token_expires_at": "2024-06-28T08:00:00Z"
    }
  }
}
```

> **Note:** Token diset sebagai `httpOnly` cookie oleh server. Frontend tidak perlu menyimpan token di localStorage.

**Business Logic & Rules:**

- **Rule #1:** Sistem menggunakan JWT dengan algoritma RS256 (asymmetric). Private key disimpan di server, public key digunakan untuk validasi.
- **Rule #2:** Setiap login menghasilkan pasangan Access Token (24 jam) dan Refresh Token (7 hari).
- **Rule #3:** Saat logout, JWT ID (`jti`) dimasukkan ke Redis blacklist dengan TTL sama dengan expiry token, sehingga token tidak bisa dipakai lagi meskipun belum expired.
- **Rule #4:** Password diverifikasi menggunakan bcrypt dengan cost factor 12. Perbandingan dilakukan secara konstan waktu (constant-time comparison) untuk mencegah timing attack.
- **Rule #5:** Jika user `is_active = false`, login ditolak dengan pesan "Akun dinonaktifkan, hubungi owner."
- **Rule #6:** Rate limiting: Maksimal 5 percobaan login gagal per IP dalam 15 menit. Jika terlampaui, IP diblokir selama 30 menit.

---

### 2.2 POST /auth/logout

Logout dan invalidate token.

**Request:** Empty body (token dari cookie)

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": null
}
```

**Business Logic & Rules:**

- **Rule #1:** Server membaca `access_token` dari cookie `httpOnly`, mengekstrak `jti` (JWT ID), dan memasukkannya ke Redis SET `jwt_blacklist` dengan TTL = sisa waktu expiry token.
- **Rule #2:** Cookie `access_token` dan `refresh_token` dihapus dari browser dengan mengatur `Max-Age=0` dan `Expires=Thu, 01 Jan 1970 00:00:00 GMT`.
- **Rule #3:** Jika token sudah tidak valid (expired atau tidak ada di cookie), tetap mengembalikan 200 OK karena tujuan akhir (user tidak lagi terautentikasi) sudah tercapai.

---

### 2.3 POST /auth/refresh

Refresh access token menggunakan refresh token cookie.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "access_token_expires_at": "2024-06-21T16:00:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Endpoint ini membaca `refresh_token` dari cookie terpisah (berbeda dari `access_token`).
- **Rule #2:** Refresh token divalidasi: signature, expiry, dan dicek apakah `jti`-nya ada di Redis blacklist.
- **Rule #3:** Jika refresh token valid, sistem generate Access Token baru dengan `jti` baru dan update cookie `access_token`.
- **Rule #4:** Refresh Token TIDAK di-rotate (tidak diganti). Hanya Access Token yang diperbarui untuk mengurangi kompleksitas.
- **Rule #5:** Jika refresh token invalid/expired, hapus semua cookie dan kembalikan 401 agar user login ulang.

---

### 2.4 GET /auth/me

Ambil data user yang sedang login.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "id": "usr_a1b2c3d4",
  "name": "Andi Wijaya",
  "email": "andi@warungku.com",
  "role": "owner",
  "permissions": [
    "products.read",
    "products.write",
    "reports.read",
    "cashier.operate"
  ]
}
```

**Business Logic & Rules:**

- **Rule #1:** Data diambil dari database berdasarkan `sub` (user_id) di JWT claims, BUKAN dari body request.
- **Rule #2:** Field `permissions` di-generate secara dinamis berdasarkan `role` user:
  - `owner`: `products.read`, `products.write`, `reports.read`, `users.manage`, `cashier.operate`
  - `cashier`: `products.read`, `transactions.operate`, `transactions.return`
- **Rule #3:** Response ini di-cache di frontend (React Query) selama 5 menit untuk mengurangi beban server.

---

## 3. Users Module (Owner Only)

### 3.1 POST /users

Register kasir baru (Owner only).

**Request:**

```json
{
  "name": "Siti Aminah",
  "email": "siti@warungku.com",
  "password": "kasir12345",
  "role": "cashier",
  "phone": "081234567890"
}
```

**Success Response (201):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "usr_e5f6g7h8",
    "name": "Siti Aminah",
    "email": "siti@warungku.com",
    "role": "cashier",
    "phone": "081234567890",
    "is_active": true,
    "created_at": "2024-06-20T10:30:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Endpoint ini dilindungi oleh middleware `RoleMiddleware(["owner"])`. Kasir yang mencoba akses akan mendapat 403 Forbidden.
- **Rule #2:** Field `role` hanya boleh bernilai `cashier`. Jika owner mencoba set `role: "owner"`, sistem akan override menjadi `cashier` karena hanya boleh ada 1 owner per toko (single-tenant model).
- **Rule #3:** Password di-hash menggunakan bcrypt(cost=12) SEBELUM disimpan ke database. Plain text password tidak pernah disimpan atau di-log.
- **Rule #4:** Email wajib unik (UNIQUE constraint di database). Jika duplikat, kembalikan 409 Conflict, BUKAN 500.
- **Rule #5:** Field `phone` opsional (nullable). Jika dikirim, harus lolos validasi format nomor Indonesia (dimulai dengan 08, panjang 10-13 digit).
- **Rule #6:** User baru otomatis `is_active = true`. Owner dapat menonaktifkan nanti via PUT /users/:id.

---

### 3.2 GET /users

List semua user (kasir) dengan pagination.

**Query Parameters:**
| Parameter | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `role` | string | - | Filter: `owner`, `cashier`, `all` |
| `is_active` | boolean | - | Filter status aktif |
| `search` | string | - | Cari nama/email |
| `page` | integer | 1 | Halaman |
| `page_size` | integer | 20 | Item per halaman |
| `sort_by` | string | `created_at` | `name`, `stock`, `base_price`, `created_at` |
| `sort_order` | string | `desc` | `asc`, `desc` |

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": [
      {
        "id": "usr_e5f6g7h8",
        "name": "Siti Aminah",
        "email": "siti@warungku.com",
        "role": "cashier",
        "phone": "081234567890",
        "is_active": true,
        "total_sales": 1545000,
        "transaction_count": 47,
        "created_at": "2024-06-20T10:30:00Z"
      }
    ],
    "metadatadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 20,
        "total_pages": 1,
        "total_items": 3,
        "has_next_page": false,
        "has_prev_page": false
      },
      "sort": { "field": "created_at", "direction": "desc" },
      "filters": { "role": "cashier", "search": "" }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Field `total_sales` dan `transaction_count` dihitung secara real-time dari tabel `transactions` dengan `cashier_id = user.id` dan `status = 'COMPLETED'`.
- **Rule #2:** Search bersifat case-insensitive dan mencocokkan substring di `name` atau `email`. Menggunakan PostgreSQL `ILIKE` atau GIN index untuk performa.
- **Rule #3:** Soft-deleted users (`deleted_at IS NOT NULL`) tidak muncul di list. Hanya owner yang bisa melihat semua user termasuk yang dinonaktifkan via filter `is_active=false`.
- **Rule #4:** Default sort adalah `created_at DESC` (user terbaru di atas).

---

### 3.3 GET /users/:id

Detail user dengan statistik penjualan.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "usr_e5f6g7h8",
    "name": "Siti Aminah",
    "email": "siti@warungku.com",
    "role": "cashier",
    "phone": "081234567890",
    "is_active": true,
    "stats": {
      "total_sales_today": 945000,
      "total_sales_month": 12450000,
      "transaction_count_today": 28,
      "transaction_count_month": 312
    },
    "created_at": "2024-06-20T10:30:00Z",
    "updated_at": "2024-06-20T10:30:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Statistik `today` dihitung dari `transactions` dengan `created_at >= CURRENT_DATE` (00:00 hari ini).
- **Rule #2:** Statistik `month` dihitung dari `transactions` dengan `created_at >= DATE_TRUNC('month', CURRENT_DATE)` (awal bulan ini).
- **Rule #3:** Semua statistik hanya menghitung transaksi dengan `status = 'COMPLETED'`. Transaksi `PENDING` (hold bill) dan `CANCELLED` tidak dihitung.
- **Rule #4:** Jika user tidak ditemukan (404), kembalikan 404 dengan pesan "User tidak ditemukan".

---

## 4. Products Module

### 4.1 POST /products

Tambah produk baru dengan multi-harga (Owner only).

**Request:**

```json
{
  "name": "Teh Kotak",
  "sku": "TK-001",
  "barcode": "8991234567890",
  "category_id": "cat_minuman",
  "base_price": 3000,
  "stock": 50,
  "min_stock_threshold": 10,
  "unit": "pcs",
  "description": "Teh kotak 250ml",
  "is_active": true,
  "price_tiers": [
    {
      "min_qty": 10,
      "price": 2500,
      "label": "Grosir 10pcs"
    },
    {
      "min_qty": 50,
      "price": 2200,
      "label": "Grosir 50pcs"
    }
  ]
}
```

**Success Response (201):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "prd_i9j0k1l2",
    "name": "Teh Kotak",
    "sku": "TK-001",
    "barcode": "8991234567890",
    "category": {
      "id": "cat_minuman",
      "name": "Minuman"
    },
    "base_price": 3000,
    "stock": 50,
    "min_stock_threshold": 10,
    "unit": "pcs",
    "is_active": true,
    "price_tiers": [
      {
        "id": "pt_a1b2",
        "min_qty": 10,
        "price": 2500,
        "label": "Grosir 10pcs"
      },
      {
        "id": "pt_c3d4",
        "min_qty": 50,
        "price": 2200,
        "label": "Grosir 50pcs"
      }
    ],
    "effective_price": 3000,
    "stock_status": "SAFE",
    "created_at": "2024-06-20T10:30:00Z",
    "updated_at": "2024-06-20T10:30:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** `sku` dan `barcode` harus unik secara global (UNIQUE constraint). Jika duplikat, kembalikan 400 dengan detail field yang bermasalah.
- **Rule #2:** `price_tiers` bersifat opsional. Jika tidak dikirim, produk hanya memiliki harga dasar (`base_price`).
- **Rule #3:** Validasi `price_tiers`:
  - `min_qty` harus > 1 (tidak boleh 1, karena harga dasar sudah untuk qty 1).
  - `price` harus > 0.
  - `price` harus < `base_price` (harga grosir harus lebih murah dari harga dasar).
  - `min_qty` harus unik per produk (tidak boleh ada 2 tier dengan min_qty sama).
  - Tier harus diurutkan ascending berdasarkan `min_qty` saat disimpan.
- **Rule #4:** `stock_status` dihitung otomatis oleh database trigger:
  - `stock < 0` -> `MINUS`
  - `0 <= stock <= min_stock_threshold` -> `LOW`
  - `stock > min_stock_threshold` -> `SAFE`
- **Rule #5:** `effective_price` di response adalah harga dasar (karena belum ada konteks kuantitas transaksi). Harga grosir hanya diterapkan saat checkout.
- **Rule #6:** Jika `category_id` dikirim tapi tidak valid (tidak ada di tabel categories), kembalikan 400.
- **Rule #7:** Setelah produk berhasil dibuat, `product_count` di tabel `categories` di-update secara otomatis (via trigger atau application layer).

---

### 4.2 GET /products

List produk dengan filter, search, dan pagination.

**Query Parameters:**
| Parameter | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `search` | string | - | Cari nama, SKU, atau barcode |
| `category_id` | string | - | Filter kategori |
| `stock_status` | string | `all` | `all`, `safe`, `low`, `minus` |
| `is_active` | boolean | true | Filter status aktif |
| `page` | integer | 1 | Halaman |
| `page_size` | integer | 20 | Item per halaman |
| `sort_by` | string | `created_at` | `name`, `stock`, `base_price`, `created_at` |
| `sort_order` | string | `desc` | `asc`, `desc` |

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": [
      {
        "id": "prd_i9j0k1l2",
        "name": "Teh Kotak",
        "sku": "TK-001",
        "barcode": "8991234567890",
        "category": { "id": "cat_minuman", "name": "Minuman" },
        "base_price": 3000,
        "stock": 50,
        "min_stock_threshold": 10,
        "unit": "pcs",
        "is_active": true,
        "stock_status": "SAFE",
        "effective_price": 3000,
        "created_at": "2024-06-20T10:30:00Z"
      }
    ],
    "metadatadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 20,
        "total_pages": 3,
        "total_items": 52,
        "has_next_page": true,
        "has_prev_page": false
      },
      "sort": { "field": "created_at", "direction": "desc" },
      "filters": {
        "search": "",
        "category_id": "",
        "stock_status": "all",
        "is_active": true
      }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Search menggunakan PostgreSQL `ILIKE` dengan wildcard `%search%` di kolom `name`, `sku`, dan `barcode`. Untuk performa > 1000 produk, gunakan GIN index dengan `pg_trgm`.
- **Rule #2:** Filter `stock_status` menggunakan kolom `stock_status` yang sudah di-compute (bukan menghitung di runtime) untuk performa optimal.
- **Rule #3:** Default `is_active = true` artinya produk yang soft delete (`deleted_at IS NOT NULL`) tidak muncul. Owner bisa set `is_active=all` untuk melihat semua produk termasuk yang dihapus.
- **Rule #4:** Response list TIDAK mengembalikan `price_tiers` lengkap untuk menghemat bandwidth. Hanya `base_price` dan `effective_price`. Detail tier ada di `GET /products/:id`.
- **Rule #5:** `effective_price` di list selalu sama dengan `base_price` karena belum ada konteks kuantitas.

---

### 4.3 GET /products/:id

Detail produk.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "prd_i9j0k1l2",
    "name": "Teh Kotak",
    "sku": "TK-001",
    "barcode": "8991234567890",
    "category": { "id": "cat_minuman", "name": "Minuman" },
    "base_price": 3000,
    "stock": 50,
    "min_stock_threshold": 10,
    "unit": "pcs",
    "description": "Teh kotak 250ml",
    "is_active": true,
    "price_tiers": [
      { "id": "pt_a1b2", "min_qty": 10, "price": 2500, "label": "Grosir 10pcs" }
    ],
    "stock_status": "SAFE",
    "sales_stats": {
      "total_sold_today": 12,
      "total_sold_month": 340
    },
    "created_at": "2024-06-20T10:30:00Z",
    "updated_at": "2024-06-20T10:30:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** `sales_stats` dihitung dari tabel `transaction_items` yang di-join dengan `transactions` (status COMPLETED).
  - `total_sold_today`: `SUM(qty)` untuk transaksi hari ini.
  - `total_sold_month`: `SUM(qty)` untuk transaksi bulan ini.
- **Rule #2:** Jika produk soft delete, tetap bisa diakses via endpoint ini (untuk melihat histori), tapi ada flag `is_active: false` dan `deleted_at` di response.
- **Rule #3:** `price_tiers` diurutkan ascending berdasarkan `min_qty` saat di-return.

---

### 4.4 PUT /products/:id

Update produk.

**Request:**

```json
{
  "name": "Teh Kotak 250ml",
  "base_price": 3500,
  "stock": 45,
  "min_stock_threshold": 15,
  "price_tiers": [
    { "id": "pt_a1b2", "min_qty": 10, "price": 2800, "label": "Grosir 10pcs" },
    { "min_qty": 50, "price": 2500, "label": "Grosir 50pcs" }
  ]
}
```

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": null
}
```

**Business Logic & Rules:**

- **Rule #1:** Partial update didukung. Hanya field yang dikirim yang di-update.
- **Rule #2:** Jika `price_tiers` dikirim, sistem melakukan UPSERT:
  - Tier yang punya `id` -> UPDATE tier tersebut.
  - Tier tanpa `id` -> INSERT tier baru.
  - Tier yang ada di DB tapi tidak ada di request -> DELETE tier tersebut.
- **Rule #3:** `sku` dan `barcode` tidak bisa diubah setelah produk dibuat (immutable) untuk menjaga integritas referensi di transaksi historis. Jika ingin ganti, buat produk baru dan nonaktifkan yang lama.
- **Rule #4:** Jika `stock` diubah, `stock_status` di-update otomatis oleh database trigger.
- **Rule #5:** Jika `category_id` diubah, `product_count` di kategori lama dan kategori baru di-update (decrement/increment).

---

### 4.5 DELETE /products/:id

Soft delete produk.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": null
}
```

**Business Logic & Rules:**

- **Rule #1:** SOFT DELETE. `deleted_at` di-set ke `NOW()`, `is_active = false`.
- **Rule #2:** Produk yang sudah soft delete tidak muncul di list default, tapi masih bisa di-query via `GET /products/:id` untuk melihat histori.
- **Rule #3:** `product_count` di kategori terkait di-decrement.
- **Rule #4:** Transaksi historis yang sudah ada tetap valid. Item di transaksi lama tetap merujuk ke produk ini (referential integrity maintained).
- **Rule #5:** Produk yang sudah soft delete tidak bisa ditambahkan ke keranjang kasir (validasi di frontend dan backend).

---

## 5. Categories Module

### 5.1 GET /categories

List semua kategori.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": [
      {
        "id": "cat_makanan",
        "name": "Makanan",
        "product_count": 24,
        "color": "#FF6B6B",
        "created_at": "2024-01-15T08:00:00Z"
      },
      {
        "id": "cat_minuman",
        "name": "Minuman",
        "product_count": 18,
        "color": "#4ECDC4",
        "created_at": "2024-01-15T08:00:00Z"
      }
    ],
    "metadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 50,
        "total_pages": 1,
        "total_items": 2,
        "has_next_page": false,
        "has_prev_page": false
      }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** `product_count` adalah field denormalized yang di-update otomatis saat produk dibuat, dihapus, atau ganti kategori. Ini untuk menghindari COUNT query yang mahal saat render UI kategori.
- **Rule #2:** List diurutkan berdasarkan `created_at ASC` (kategori yang dibuat pertama di atas).
- **Rule #3:** Kategori tidak bisa dihapus jika masih ada produk aktif yang merujuk ke kategori tersebut. Jika owner mencoba hapus, kembalikan 422 "Kategori masih memiliki produk aktif."

---

### 5.2 POST /categories (Owner only)

**Request:**

```json
{
  "name": "Snack",
  "color": "#FFE66D",
  "description": "Makanan ringan"
}
```

**Business Logic & Rules:**

- **Rule #1:** `name` harus unik (UNIQUE constraint). Jika duplikat, kembalikan 409.
- **Rule #2:** `color` harus format HEX valid (regex: `^#[0-9A-Fa-f]{6}$`). Default: `#6B7280`.
- **Rule #3:** `description` opsional, max 500 karakter.
- **Rule #4:** Saat kategori baru dibuat, `product_count` otomatis 0.

---

## 6. Transactions Module (Cashier Only)

### 6.1 POST /transactions

Buat transaksi baru (checkout atau hold bill).

**Request:**

```json
{
  "type": "direct",
  "items": [
    {
      "product_id": "prd_i9j0k1l2",
      "qty": 12,
      "price_at_time": 2500,
      "notes": ""
    },
    {
      "product_id": "prd_m3n4o5p6",
      "qty": 2,
      "price_at_time": 30000,
      "notes": "Extra pedas"
    }
  ],
  "payment": {
    "method": "cash",
    "amount_paid": 100000,
    "change": 20000
  },
  "hold_note": null,
  "discount_amount": 0,
  "tax_amount": 0
}
```

> **Hold Bill:** Set `type: "hold"` dan `hold_note: "Meja 5 / Andi"`. Field `payment` boleh null.

**Success Response (201) - Direct Payment:**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "txn_q7r8s9t0",
    "transaction_code": "TRX-200624-0001",
    "type": "direct",
    "status": "COMPLETED",
    "cashier": {
      "id": "usr_e5f6g7h8",
      "name": "Siti Aminah"
    },
    "items": [
      {
        "id": "tmi_u1v2w3x4",
        "product": {
          "id": "prd_i9j0k1l2",
          "name": "Teh Kotak",
          "sku": "TK-001"
        },
        "qty": 12,
        "unit_price": 2500,
        "subtotal": 30000,
        "notes": ""
      }
    ],
    "summary": {
      "subtotal": 90000,
      "discount_amount": 0,
      "tax_amount": 0,
      "total": 90000
    },
    "payment": {
      "method": "cash",
      "amount_paid": 100000,
      "change": 20000
    },
    "created_at": "2024-06-20T14:30:00Z"
  }
}
```

**Success Response (201) - Hold Bill:**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "txn_q7r8s9t0",
    "transaction_code": "TRX-200624-0001",
    "type": "hold",
    "status": "PENDING",
    "hold_note": "Meja 5 / Andi",
    "cashier": {
      "id": "usr_e5f6g7h8",
      "name": "Siti Aminah"
    },
    "items": [...],
    "summary": { "subtotal": 90000, "total": 90000 },
    "created_at": "2024-06-20T14:30:00Z"
  },
}
```

**Business Logic & Rules:**

#### A. Transaction Code Generation

- **Rule #1:** `transaction_code` di-generate otomatis dengan format: `TRX-YYMMDD-XXXX` (contoh: `TRX-200624-0001`).
- **Rule #2:** Nomor urut `XXXX` reset ke 0001 setiap hari (berdasarkan `created_at` date).
- **Rule #3:** Generasi kode menggunakan database sequence atau atomic counter di Redis untuk mencegah duplikat saat konkurensi tinggi.

#### B. Price Resolution Logic

- **Rule #4:** `price_at_time` di body request adalah harga yang sudah dihitung di frontend. Backend WAJIB melakukan re-kalkulasi untuk mencegah manipulasi.
- **Rule #5:** Re-kalkulasi harga:
  1. Ambil `base_price` dan `price_tiers` produk dari database.
  2. Cari tier dengan `min_qty` terbesar yang <= `qty` item.
  3. Jika ada tier yang cocok, gunakan `tier.price`. Jika tidak, gunakan `base_price`.
  4. Bandingkan harga hasil kalkulasi dengan `price_at_time` dari request. Jika berbeda, kembalikan 409 dengan pesan "Harga tidak valid, silakan refresh halaman."
- **Rule #6:** `subtotal` per item = `unit_price` x `qty`. `subtotal` transaksi = SUM semua item subtotal.
- **Rule #7:** `total` = `subtotal` - `discount_amount` + `tax_amount`.

#### C. Stock Deduction (Critical)

- **Rule #8:** Semua operasi stok menggunakan transaction isolation level `SERIALIZABLE` atau `REPEATABLE READ` dengan `SELECT FOR UPDATE`.
- **Rule #9:** Flow pemotongan stok:
  ```
  BEGIN TRANSACTION;
  SELECT stock FROM products WHERE id = ? FOR UPDATE;
  -- Verifikasi stok cukup (opsional, karena sistem toleran minus)
  UPDATE products SET stock = stock - ? WHERE id = ?;
  COMMIT;
  ```
- **Rule #10:** Sistem TIDAK memblokir transaksi meskipun stok menjadi minus. Ini adalah fitur inti "Stok Minus Toleran".
- **Rule #11:** Jika stok menjadi minus setelah pemotongan, sistem otomatis INSERT record ke tabel `stock_alerts` dengan `alert_type = 'MINUS'`.
- **Rule #12:** Jika stok setelah pemotongan <= `min_stock_threshold` (tapi masih >= 0), INSERT `stock_alerts` dengan `alert_type = 'LOW'`.

#### D. Hold Bill Specific Rules

- **Rule #13:** Hold bill (`type: "hold"`) WAJIB memiliki `hold_note` (max 100 karakter). Jika kosong, kembalikan 400.
- **Rule #14:** Hold bill memiliki `status = "PENDING"` dan tidak memiliki data `payment`.
- **Rule #15:** Stok langsung dipotong saat hold bill dibuat (sama seperti direct). Ini mencegah kasir lain menjual barang yang sudah dipesan.
- **Rule #16:** Hold bill muncul di `GET /transactions/hold` dan bisa dilanjutkan ke pembayaran via `POST /transactions/:id/complete`.

#### E. Payment Validation

- **Rule #17:** Untuk direct payment, `payment.method` harus salah satu dari: `cash`, `qris`, `transfer`.
- **Rule #18:** `amount_paid` harus >= `total`. Jika kurang, kembalikan 422.
- **Rule #19:** `change` harus = `amount_paid` - `total`. Jika tidak cocok, kembalikan 422.
- **Rule #20:** `cashier_id` diambil dari JWT claims (user yang login), BUKAN dari body request. Ini mencegah kasir A mencatat transaksi atas nama kasir B.

---

### 6.2 GET /transactions

List transaksi dengan filter lengkap.

**Query Parameters:**
| Parameter | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `status` | string | - | `completed`, `pending`, `cancelled`, `all` |
| `type` | string | - | `direct`, `hold`, `all` |
| `cashier_id` | string | - | Filter kasir |
| `date_from` | date | - | Filter tanggal mulai (YYYY-MM-DD) |
| `date_to` | date | - | Filter tanggal akhir (YYYY-MM-DD) |
| `search` | string | - | Cari kode transaksi |
| `page` | integer | 1 | Halaman |
| `per_page` | integer | 20 | Item per halaman |
| `sort_by` | string | `created_at` | `created_at`, `total`, `transaction_code` |
| `sort_order` | string | `desc` | `asc`, `desc` |

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": [
      {
        "id": "txn_q7r8s9t0",
        "transaction_code": "TRX-200624-0001",
        "type": "direct",
        "status": "COMPLETED",
        "cashier": { "id": "usr_e5f6g7h8", "name": "Siti Aminah" },
        "item_count": 3,
        "total": 90000,
        "payment_method": "cash",
        "created_at": "2024-06-20T14:30:00Z"
      }
    ],
    "metadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 20,
        "total_pages": 5,
        "total_items": 98,
        "has_next_page": true,
        "has_prev_page": false
      },
      "sort": { "field": "created_at", "direction": "desc" },
      "filters": {
        "status": "all",
        "type": "all",
        "cashier_id": "",
        "date_from": "",
        "date_to": "",
        "search": ""
      }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Filter `date_from` dan `date_to` menggunakan `created_at::date` (PostgreSQL date casting) agar mencakup seluruh hari.
- **Rule #2:** `item_count` dihitung dari COUNT `transaction_items` per transaksi (subquery atau JOIN).
- **Rule #3:** Kasir hanya bisa melihat transaksi miliknya sendiri, KECUALI owner yang bisa melihat semua transaksi. Ini diatur di middleware:
  - Jika role = cashier dan `cashier_id` tidak dikirim, otomatis filter `cashier_id = current_user_id`.
  - Jika role = cashier dan `cashier_id` dikirim dengan ID orang lain, kembalikan 403.
- **Rule #4:** Default sort `created_at DESC` (transaksi terbaru di atas).

---

### 6.3 GET /transactions/:id

Detail transaksi lengkap.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "txn_q7r8s9t0",
    "transaction_code": "TRX-200624-0001",
    "type": "direct",
    "status": "COMPLETED",
    "cashier": { "id": "usr_e5f6g7h8", "name": "Siti Aminah" },
    "items": [
      {
        "id": "tmi_u1v2w3x4",
        "product": {
          "id": "prd_i9j0k1l2",
          "name": "Teh Kotak",
          "sku": "TK-001"
        },
        "qty": 12,
        "unit_price": 2500,
        "subtotal": 30000,
        "notes": ""
      }
    ],
    "summary": {
      "subtotal": 90000,
      "discount_amount": 0,
      "tax_amount": 0,
      "total": 90000
    },
    "payment": {
      "method": "cash",
      "amount_paid": 100000,
      "change": 20000
    },
    "void_logs": [],
    "created_at": "2024-06-20T14:30:00Z",
    "updated_at": "2024-06-20T14:30:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** `void_logs` di-join dari tabel `void_logs` dengan filter `transaction_id = :id`.
- **Rule #2:** Jika transaksi tidak ditemukan, kembalikan 404.
- **Rule #3:** Kasir hanya bisa melihat detail transaksi miliknya sendiri. Owner bisa melihat semua. Jika kasir mencoba akses transaksi kasir lain, kembalikan 403.

---

### 6.4 POST /transactions/:id/return

Return item dari transaksi.

**Request:**

```json
{
  "items": [
    {
      "transaction_item_id": "tmi_u1v2w3x4",
      "qty": 2,
      "reason": "Produk rusak"
    }
  ]
}
```

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "transaction_id": "txn_q7r8s9t0",
    "returned_items": [
      {
        "transaction_item_id": "tmi_u1v2w3x4",
        "product_name": "Teh Kotak",
        "qty_returned": 2,
        "refund_amount": 5000,
        "reason": "Produk rusak"
      }
    ],
    "new_total": 85000,
    "void_log_id": "vld_y5z6a7b8",
    "created_at": "2024-06-20T15:00:00Z"
  }
}
```

**Business Logic & Rules:**

#### A. Return Validation

- **Rule #1:** Transaksi harus memiliki `status = "COMPLETED"`. Transaksi `PENDING` (hold bill) tidak bisa di-return. Jika dicoba, kembalikan 422 "Hold bill belum dibayar, batalkan saja."
- **Rule #2:** `transaction_item_id` harus valid dan milik transaksi ini. Jika tidak, kembalikan 404.
- **Rule #3:** `qty` yang di-return tidak boleh melebihi `qty` asli di `transaction_items`. Jika melebihi, kembalikan 422.
- **Rule #4:** Jika item sudah pernah di-return sebelumnya, total qty return (existing + new) tidak boleh melebihi qty asli.

#### B. Stock Reversal

- **Rule #5:** Stok produk dikembalikan (ditambah) sejumlah `qty_returned`.
- **Rule #6:** Stock reversal menggunakan `SELECT FOR UPDATE` yang sama dengan pemotongan stok untuk mencegah race condition.
- **Rule #7:** Jika stok yang dikembalikan menyebabkan `stock_status` berubah (misal dari MINUS ke LOW, atau LOW ke SAFE), database trigger otomatis update `stock_status`.

#### C. Financial Adjustment

- **Rule #8:** `refund_amount` = `unit_price` (saat transaksi) x `qty_returned`. Menggunakan harga saat transaksi, BUKAN harga sekarang.
- **Rule #9:** `new_total` transaksi = `total` lama - SUM semua `refund_amount`.
- **Rule #10:** `subtotal` transaksi juga dikurangi. `discount_amount` dan `tax_amount` tidak diubah (proportional adjustment out of scope untuk MVP).

#### D. Audit Trail

- **Rule #11:** Setiap return mencatat record ke `void_logs` dengan field:
  - `transaction_id`, `transaction_item_id`, `cashier_id` (kasir yang melakukan return), `product_id`, `qty_returned`, `refund_amount`, `reason`, `created_at`.
- **Rule #12:** `cashier_id` di `void_logs` diambil dari JWT claims user yang melakukan return, BUKAN dari `cashier_id` transaksi asli. Ini memungkinkan kasir shift berbeda melakukan return atas transaksi kasir sebelumnya, tapi tetap tercatat siapa yang melakukan.

---

### 6.5 POST /transactions/:id/complete

Lanjutkan hold bill ke pembayaran.

**Request:**

```json
{
  "payment": {
    "method": "cash",
    "amount_paid": 100000,
    "change": 10000
  }
}
```

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "id": "txn_q7r8s9t0",
    "status": "COMPLETED",
    "payment": { "method": "cash", "amount_paid": 100000, "change": 10000 },
    "completed_at": "2024-06-20T15:15:00Z"
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Endpoint ini hanya untuk transaksi dengan `type = "hold"` dan `status = "PENDING"`. Jika transaksi sudah COMPLETED, kembalikan 422 "Transaksi sudah selesai."
- **Rule #2:** `payment` wajib ada dan mengikuti aturan validasi yang sama dengan `POST /transactions` (Rule #17-#20).
- **Rule #3:** Stok TIDAK dipotong lagi karena sudah dipotong saat hold bill dibuat.
- **Rule #4:** `status` diubah ke `COMPLETED`, `completed_at` di-set ke `NOW()`.
- **Rule #5:** `hold_note` tetap ada di record untuk referensi (tidak dihapus).
- **Rule #6:** Hanya kasir yang membuat hold bill atau owner yang bisa melanjutkan. Kasir lain mendapat 403.

---

### 6.6 GET /transactions/hold

List hold bill aktif.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "resutl": {
    "data": [
      {
        "id": "txn_q7r8s9t0",
        "transaction_code": "TRX-200624-0001",
        "hold_note": "Meja 5 / Andi",
        "cashier": { "id": "usr_e5f6g7h8", "name": "Siti Aminah" },
        "item_count": 3,
        "total": 90000,
        "held_at": "2024-06-20T14:30:00Z",
        "elapsed_minutes": 45
      }
    ],
    "metadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 20,
        "total_pages": 1,
        "total_items": 3,
        "has_next_page": false,
        "has_prev_page": false
      }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Hanya mengembalikan transaksi dengan `type = "hold"` dan `status = "PENDING"`.
- **Rule #2:** `elapsed_minutes` dihitung di runtime: `EXTRACT(EPOCH FROM (NOW() - created_at)) / 60`.
- **Rule #3:** Diurutkan berdasarkan `created_at ASC` (hold bill terlama di atas) agar kasir bisa melayani yang duluan datang.
- **Rule #4:** Kasir hanya melihat hold bill miliknya sendiri. Owner melihat semua hold bill.
- **Rule #5:** Hold bill yang sudah lebih dari 4 jam bisa di-highlight dengan warna merah (warning) di frontend.

---

## 7. Reports Module (Owner Only)

### 7.1 GET /reports/sales

Laporan penjualan dengan aggregasi.

**Query Parameters:**
| Parameter | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `period` | string | `daily` | `daily`, `weekly`, `monthly`, `yearly` |
| `date_from` | date | - | Tanggal mulai |
| `date_to` | date | - | Tanggal akhir |
| `cashier_id` | string | - | Filter per kasir |
| `group_by` | string | `date` | `date`, `cashier`, `product`, `category` |

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": {
      "summary": {
        "total_revenue": 1545000,
        "total_transactions": 47,
        "average_transaction_value": 32872,
        "total_items_sold": 156
      },
      "breakdown": [
        {
          "date": "2024-06-20",
          "revenue": 945000,
          "transaction_count": 28,
          "items_sold": 89
        },
        {
          "date": "2024-06-19",
          "revenue": 600000,
          "transaction_count": 19,
          "items_sold": 67
        }
      ],
      "cashier_performance": [
        {
          "cashier_id": "usr_e5f6g7h8",
          "cashier_name": "Siti Aminah",
          "total_revenue": 945000,
          "transaction_count": 28
        }
      ]
    },
    "metadata": {
      "filters": {
        "period": "daily",
        "date_from": "2024-06-19",
        "date_to": "2024-06-20",
        "cashier_id": ""
      }
    }
  }
}
```

**Business Logic & Rules:**

#### A. Period & Date Range

- **Rule #1:** Jika `date_from` dan `date_to` tidak dikirim, default adalah periode terakhir sesuai `period`:
  - `daily`: 7 hari terakhir (termasuk hari ini).
  - `weekly`: 4 minggu terakhir.
  - `monthly`: 12 bulan terakhir.
  - `yearly`: 2 tahun terakhir.
- **Rule #2:** `date_from` dan `date_to` harus valid date format (YYYY-MM-DD). Jika tidak, kembalikan 400.
- **Rule #3:** `date_to` tidak boleh lebih kecil dari `date_from`. Jika ya, kembalikan 400.

#### B. Aggregation Logic

- **Rule #4:** Semua agregasi hanya menghitung transaksi dengan `status = 'COMPLETED'`.
- **Rule #5:** `total_revenue` = SUM(`total`) dari semua transaksi dalam range.
- **Rule #6:** `total_transactions` = COUNT transaksi dalam range.
- **Rule #7:** `average_transaction_value` = `total_revenue` / `total_transactions` (dibulatkan ke bawah).
- **Rule #8:** `total_items_sold` = SUM(`qty`) dari `transaction_items` yang di-join dengan transaksi COMPLETED dalam range.

#### C. Group By Logic

- **Rule #9:** `group_by = "date"`: Breakdown per tanggal (`created_at::date`).
- **Rule #10:** `group_by = "cashier"`: Breakdown per kasir (cashier_id + name).
- **Rule #11:** `group_by = "product"`: Breakdown per produk (product_id + name + SUM(qty) + SUM(subtotal)).
- **Rule #12:** `group_by = "category"`: Breakdown per kategori (category_id + name + COUNT(DISTINCT products) + SUM(subtotal)).

#### D. Performance

- **Rule #13:** Query laporan diarahkan ke PostgreSQL read replica (jika tersedia) untuk mengurangi beban primary database.
- **Rule #14:** Hasil laporan di-cache di Redis selama 5 menit untuk query yang sama (key: `report:sales:{hash_of_params}`).

---

### 7.2 GET /reports/stock-alerts

Notifikasi stok menipis dan minus.

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "critical_count": 2,
    "low_count": 5,
    "alerts": [
      {
        "product_id": "prd_i9j0k1l2",
        "product_name": "Teh Kotak",
        "current_stock": -3,
        "min_threshold": 10,
        "status": "MINUS",
        "last_updated": "2024-06-20T14:30:00Z"
      },
      {
        "product_id": "prd_m3n4o5p6",
        "product_name": "Es Teh Manis",
        "current_stock": 2,
        "min_threshold": 10,
        "status": "LOW",
        "last_updated": "2024-06-20T12:00:00Z"
      }
    ]
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Data diambil dari tabel `stock_alerts` yang di-join dengan `products`.
- **Rule #2:** Hanya mengembalikan alert dengan `is_resolved = false`.
- **Rule #3:** Diurutkan berdasarkan `alert_type` (MINUS dulu, kemudian LOW) lalu `created_at DESC`.
- **Rule #4:** `critical_count` = COUNT alert dengan `alert_type = 'MINUS'`.
- **Rule #5:** `low_count` = COUNT alert dengan `alert_type = 'LOW'`.
- **Rule #6:** Alert otomatis di-resolve (set `is_resolved = true`) saat owner mengupdate stok produk melebihi `min_stock_threshold` (via trigger di database atau application layer saat PUT /products/:id).
- **Rule #7:** Alert tidak dihapus setelah resolve. Mereka tetap ada dengan `is_resolved = true` dan `resolved_at` terisi untuk histori.

---

## 8. Void Logs Module (Owner Only)

### 8.1 GET /void-logs

Audit trail pembatalan/return.

**Query Parameters:**
| Parameter | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `cashier_id` | string | - | Filter kasir |
| `date_from` | date | - | Filter tanggal |
| `date_to` | date | - | Filter tanggal |
| `page` | integer | 1 | Halaman |
| `per_page` | integer | 20 | Item per halaman |

**Success Response (200):**

```json
{
  "code": 200,
  "message": "successfully",
  "result": {
    "data": [
      {
        "id": "vld_y5z6a7b8",
        "transaction_id": "txn_q7r8s9t0",
        "transaction_code": "TRX-200624-0001",
        "cashier": { "id": "usr_e5f6g7h8", "name": "Siti Aminah" },
        "product": { "id": "prd_i9j0k1l2", "name": "Teh Kotak" },
        "qty_returned": 2,
        "refund_amount": 5000,
        "reason": "Produk rusak",
        "created_at": "2024-06-20T15:00:00Z"
      }
    ],
    "metadata": {
      "pagination": {
        "current_page": 1,
        "per_page": 20,
        "total_pages": 2,
        "total_items": 35,
        "has_next_page": true,
        "has_prev_page": false
      },
      "filters": { "cashier_id": "", "date_from": "", "date_to": "" }
    }
  }
}
```

**Business Logic & Rules:**

- **Rule #1:** Endpoint ini hanya bisa diakses oleh owner. Kasir mendapat 403.
- **Rule #2:** `cashier` di response adalah kasir yang MELAKUKAN return (dari `void_logs.cashier_id`), bukan kasir yang membuat transaksi asli.
- **Rule #3:** Filter `date_from` dan `date_to` menggunakan `created_at` di tabel `void_logs`.
- **Rule #4:** Default sort `created_at DESC` (return terbaru di atas).
- **Rule #5:** `transaction_code` di-join dari tabel `transactions` untuk kemudahan identifikasi.
- **Rule #6:** Data ini tidak bisa diubah atau dihapus (immutable audit trail). Tidak ada endpoint PUT atau DELETE untuk void logs.

---

## 9. Appendix: Business Logic & Rules

### 9.1 Ringkasan Aturan Global

|  #  | Aturan                                       | Modul       | Keterangan                     |
| :-: | -------------------------------------------- | ----------- | ------------------------------ |
|  1  | JWT di `httpOnly` cookie                     | Auth        | Mitigasi XSS attack            |
|  2  | bcrypt(cost=12) untuk password               | Auth        | Hashing aman                   |
|  3  | Rate limit 5x login gagal per 15 menit       | Auth        | Mencegah brute force           |
|  4  | `cashier_id` diambil dari JWT, bukan body    | Transaksi   | Audit trail akurat             |
|  5  | `sku` dan `barcode` immutable setelah dibuat | Produk      | Integritas referensi           |
|  6  | Harga grosir < harga dasar                   | Produk      | Logika bisnis grosir           |
|  7  | `min_qty` tier harus > 1                     | Produk      | Harga dasar untuk qty 1        |
|  8  | Stok dipotong saat hold bill dibuat          | Transaksi   | Mencegah double sell           |
|  9  | Transaksi tidak diblokir meski stok minus    | Transaksi   | Fitur toleransi stok           |
| 10  | `SELECT FOR UPDATE` untuk stok               | Transaksi   | Mencegah race condition        |
| 11  | Return hanya untuk transaksi COMPLETED       | Return/Void | Hold bill tidak bisa di-return |
| 12  | Refund pakai harga saat transaksi            | Return/Void | Bukan harga sekarang           |
| 13  | Void logs immutable                          | Audit       | Tidak bisa edit/delete         |
| 14  | Soft delete untuk users & products           | Global      | Data historis tetap ada        |
| 15  | Owner tidak bisa hapus dirinya sendiri       | Users       | Mencegah lockout               |

### 9.2 State Machine Transaksi

```
[PENDING] --(POST /:id/complete)--> [COMPLETED]
   |                                      |
   |                                      |
   |--(tidak ada cancel)-->              |--(POST /:id/return)--> [COMPLETED]
                                         |    (stok dikembalikan,
                                         |     total dikurangi,
                                         |     void log tercatat)
```

> **Catatan:** Tidak ada state `CANCELLED` untuk transaksi yang sudah dibuat. Hold bill yang dibatalkan di-handle dengan menghapus record (hard delete) karena belum ada pembayaran. Transaksi COMPLETED tidak bisa di-cancel, hanya bisa di-return item per item.

### 9.3 Price Tier Resolution Algorithm

```
function resolvePrice(product, qty):
    applicable_tiers = filter(product.price_tiers, tier.min_qty <= qty)

    if applicable_tiers is empty:
        return product.base_price

    best_tier = max(applicable_tiers, by=min_qty)  // tier dengan min_qty terbesar
    return best_tier.price
```

**Contoh:**

- Produk "Teh Kotak" dengan base_price = 3000
- Tier 1: min_qty=10, price=2500
- Tier 2: min_qty=50, price=2200

| Qty   | Harga yang Dipakai | Keterangan  |
| ----- | ------------------ | ----------- |
| 1-9   | 3000               | Harga dasar |
| 10-49 | 2500               | Tier 1      |
| 50+   | 2200               | Tier 2      |

### 9.4 Stock Status Computation

```sql
CREATE OR REPLACE FUNCTION update_stock_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.stock < 0 THEN
        NEW.stock_status := 'MINUS';
        -- Insert stock alert MINUS
        INSERT INTO stock_alerts (product_id, alert_type, stock_at_alert)
        VALUES (NEW.id, 'MINUS', NEW.stock)
        ON CONFLICT DO NOTHING;
    ELSIF NEW.stock <= NEW.min_stock_threshold THEN
        NEW.stock_status := 'LOW';
        -- Insert stock alert LOW
        INSERT INTO stock_alerts (product_id, alert_type, stock_at_alert)
        VALUES (NEW.id, 'LOW', NEW.stock)
        ON CONFLICT DO NOTHING;
    ELSE
        NEW.stock_status := 'SAFE';
        -- Resolve existing alerts
        UPDATE stock_alerts
        SET is_resolved = true, resolved_at = NOW()
        WHERE product_id = NEW.id AND is_resolved = false;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

### 9.5 Role-Based Access Control (RBAC) Matrix

| Endpoint Pattern                  | Owner | Cashier | Method        | Deskripsi Akses                       |
| :-------------------------------- | :---- | :------ | :------------ | :------------------------------------ |
| `POST /auth/login`                | Yes   | Yes     | Public        | Semua user bisa login                 |
| `POST /auth/logout`               | Yes   | Yes     | Authenticated | Semua user bisa logout                |
| `POST /auth/refresh`              | Yes   | Yes     | Authenticated | Semua user bisa refresh token         |
| `GET /auth/me`                    | Yes   | Yes     | Authenticated | Semua user ambil data diri            |
| `POST /users`                     | Yes   | No      | Owner Only    | Register kasir baru                   |
| `GET /users`                      | Yes   | No      | Owner Only    | List semua user                       |
| `GET /users/:id`                  | Yes   | No      | Owner Only    | Detail user + stats                   |
| `PUT /users/:id`                  | Yes   | No      | Owner Only    | Update data user                      |
| `DELETE /users/:id`               | Yes   | No      | Owner Only    | Nonaktifkan user                      |
| `POST /products`                  | Yes   | No      | Owner Only    | Tambah produk baru                    |
| `GET /products`                   | Yes   | Yes     | Authenticated | List produk (kasir hanya lihat aktif) |
| `GET /products/:id`               | Yes   | Yes     | Authenticated | Detail produk                         |
| `PUT /products/:id`               | Yes   | No      | Owner Only    | Update produk                         |
| `DELETE /products/:id`            | Yes   | No      | Owner Only    | Soft delete produk                    |
| `GET /categories`                 | Yes   | Yes     | Authenticated | List kategori                         |
| `POST /categories`                | Yes   | No      | Owner Only    | Tambah kategori                       |
| `POST /transactions`              | No    | Yes     | Cashier Only  | Buat transaksi/checkout               |
| `GET /transactions`               | Yes   | Yes     | Authenticated | Kasir: milik sendiri. Owner: semua.   |
| `GET /transactions/:id`           | Yes   | Yes     | Authenticated | Kasir: milik sendiri. Owner: semua.   |
| `POST /transactions/:id/return`   | No    | Yes     | Cashier Only  | Return item transaksi                 |
| `POST /transactions/:id/complete` | No    | Yes     | Cashier Only  | Lanjutkan hold bill                   |
| `GET /transactions/hold`          | No    | Yes     | Cashier Only  | List hold bill aktif                  |
| `GET /reports/sales`              | Yes   | No      | Owner Only    | Laporan penjualan                     |
| `GET /reports/stock-alerts`       | Yes   | No      | Owner Only    | Alert stok menipis/minus              |
| `GET /void-logs`                  | Yes   | No      | Owner Only    | Audit trail return                    |

### 9.6 Error Response Standard

Semua error response mengikuti format berikut:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE_UPPERCASE",
    "message": "Pesan error yang human-readable",
    "details": [{ "field": "field_name", "message": "Pesan spesifik field" }],
    "timestamp": "2024-06-20T10:30:00Z",
    "request_id": "req_abc123xyz"
  }
}
```

**Daftar Error Codes:**

| Code                      | HTTP Status | Konteks                     |
| ------------------------- | ----------- | --------------------------- |
| `INVALID_CREDENTIALS`     | 401         | Email/password salah        |
| `UNAUTHORIZED`            | 401         | Token tidak ada/invalid     |
| `FORBIDDEN`               | 403         | Role tidak cukup            |
| `NOT_FOUND`               | 404         | Resource tidak ditemukan    |
| `VALIDATION_ERROR`        | 400         | Input tidak valid           |
| `EMAIL_EXISTS`            | 409         | Email duplikat              |
| `SKU_EXISTS`              | 409         | SKU duplikat                |
| `RACE_CONDITION`          | 409         | Stok berubah saat transaksi |
| `BUSINESS_RULE_VIOLATION` | 422         | Melanggar aturan bisnis     |
| `RATE_LIMIT_EXCEEDED`     | 429         | Terlalu banyak request      |
| `INTERNAL_ERROR`          | 500         | Kesalahan server            |

### 9.7 Data Flow: Checkout Lengkap

```
+---------------------------------------------------------------------+
|                         KASIR MENEKAN BAYAR                         |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 1. VALIDASI KERANJANG                                              |
|    - Items tidak kosong                                            |
|    - Semua product_id valid dan is_active = true                   |
|    - Semua qty > 0                                                 |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 2. RESOLUSI HARGA (per item)                                       |
|    - Ambil base_price + price_tiers dari DB                        |
|    - Cari tier dengan min_qty terbesar yang <= qty                 |
|    - Bandingkan dengan price_at_time dari request                  |
|    - Jika mismatch -> 409 RACE_CONDITION                            |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 3. KALKULASI TOTAL                                                 |
|    - subtotal_item = unit_price x qty                              |
|    - subtotal_tx = SUM(subtotal_item)                              |
|    - total = subtotal_tx - discount + tax                          |
|    - Validasi payment: amount_paid >= total                        |
|    - Validasi change: amount_paid - total                         |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 4. GENERATE KODE TRANSAKSI                                         |
|    - Format: TRX-YYMMDD-XXXX                                       |
|    - Nomor urut atomic (sequence/Redis)                          |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 5. DATABASE TRANSACTION (ACID)                                     |
|    BEGIN SERIALIZABLE;                                             |
|    |                                                               |
|    |-- INSERT INTO transactions (...)                              |
|    |   -> Dapat transaction_id                                       |
|    |                                                               |
|    |-- FOR EACH item:                                              |
|    |   |-- SELECT stock FROM products WHERE id = ? FOR UPDATE      |
|    |   |-- INSERT INTO transaction_items (...)                      |
|    |   |-- UPDATE products SET stock = stock - qty WHERE id = ?   |
|    |                                                               |
|    |-- IF stock menjadi MINUS/LOW:                                  |
|    |   |-- INSERT INTO stock_alerts (...)                          |
|    |                                                               |
|    |-- COMMIT;                                                     |
|                                                                    |
|    Jika COMMIT gagal (race condition):                            |
|    |-- ROLLBACK -> 409 RACE_CONDITION                              |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 6. RESPONSE                                                        |
|    - Return detail transaksi lengkap                               |
|    - Trigger print struk (opsional, via Web API)                   |
+---------------------------------------------------------------------+
```

### 9.8 Data Flow: Return Item Lengkap

```
+---------------------------------------------------------------------+
|                    KASIR MENEKAN RETURN ITEM                        |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 1. VALIDASI TRANSAKSI                                              |
|    - Transaksi harus COMPLETED (bukan PENDING)                     |
|    - Kasir harus punya akses ke transaksi ini                      |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 2. VALIDASI ITEM                                                   |
|    - transaction_item_id valid dan milik transaksi ini             |
|    - qty_returned > 0                                              |
|    - qty_returned <= qty_asli - qty_sudah_return                   |
|    - reason tidak boleh kosong                                     |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 3. KALKULASI REFUND                                                |
|    - refund_amount = unit_price (saat transaksi) x qty_returned    |
|    - new_total = total_lama - refund_amount                        |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 4. DATABASE TRANSACTION (ACID)                                     |
|    BEGIN SERIALIZABLE;                                             |
|    |                                                               |
|    |-- SELECT stock FROM products WHERE id = ? FOR UPDATE          |
|    |                                                               |
|    |-- INSERT INTO void_logs (...)                                 |
|    |   -> cashier_id = user yang login (bukan pembuat transaksi)  |
|    |                                                               |
|    |-- UPDATE products SET stock = stock + qty_returned            |
|    |                                                               |
|    |-- UPDATE transactions SET                                     |
|    |   total = total - refund_amount,                             |
|    |   subtotal = subtotal - refund_amount                        |
|    |                                                               |
|    |-- COMMIT;                                                     |
+---------------------------------------------------------------------+
                                    |
                                    v
+---------------------------------------------------------------------+
| 5. RESPONSE                                                        |
|    - Return detail return + new_total                              |
|    - Update tampilan struk di frontend                             |
+---------------------------------------------------------------------+
```

---

**Dokumen ini merupakan kontrak API lengkap untuk Smart Cost V1.**
**Versi:** 1.0
**Terakhir Diperbarui:** Juni 2026
