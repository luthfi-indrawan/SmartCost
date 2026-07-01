# Database Design Document (DATABASE_DESIGN) - Smart Cost V1

**Versi:** 1.1  
**Status:** Regenerated & Synced with PRD  
**DBMS:** PostgreSQL 15  
**ORM:** GORM v2 (Go) / Raw SQL dengan `pgx`

---

## 1. Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    USERS ||--o{ TRANSACTIONS : creates
    USERS ||--o{ VOID_LOGS : performs
    CATEGORIES ||--o{ PRODUCTS : categorizes
    PRODUCTS ||--o{ PRODUCT_PRICES : has
    PRODUCTS ||--o{ TRANSACTION_ITEMS : contains
    PRODUCTS ||--o{ VOID_LOGS : referenced
    PRODUCTS ||--o{ STOCK_ALERTS : triggers
    TRANSACTIONS ||--o{ TRANSACTION_ITEMS : contains
    TRANSACTIONS ||--o{ VOID_LOGS : referenced

    USERS {
        uuid id PK
        varchar name
        varchar email UK
        varchar password_hash
        enum role "owner, cashier"
        varchar phone
        boolean is_active
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    CATEGORIES {
        uuid id PK
        varchar name UK
        varchar color
        text description
        int product_count
        timestamp created_at
        timestamp updated_at
    }

    PRODUCTS {
        uuid id PK
        varchar name
        varchar sku UK
        varchar barcode UK
        uuid category_id FK
        int base_price
        int stock
        int min_stock_threshold
        varchar unit
        text description
        boolean is_active
        enum stock_status "SAFE, LOW, MINUS"
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    PRODUCT_PRICES {
        uuid id PK
        uuid product_id FK
        int min_qty
        int price
        varchar label
        timestamp created_at
        timestamp updated_at
    }

    TRANSACTIONS {
        uuid id PK
        varchar transaction_code UK
        enum type "direct, hold"
        enum status "PENDING, COMPLETED, CANCELLED"
        uuid cashier_id FK
        text hold_note
        int subtotal
        int discount_amount
        int tax_amount
        int total
        enum payment_method "cash, qris, transfer"
        int amount_paid
        int change_amount
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }

    TRANSACTION_ITEMS {
        uuid id PK
        uuid transaction_id FK
        uuid product_id FK
        int qty
        int unit_price
        int subtotal
        text notes
        timestamp created_at
    }

    VOID_LOGS {
        uuid id PK
        uuid transaction_id FK
        uuid transaction_item_id FK
        uuid cashier_id FK
        uuid product_id FK
        int qty_returned
        int refund_amount
        text reason
        timestamp created_at
        timestamp updated_at
    }

    STOCK_ALERTS {
        uuid id PK
        uuid product_id FK
        enum alert_type "LOW, MINUS"
        int stock_at_alert
        boolean is_resolved
        timestamp resolved_at
        timestamp created_at
        timestamp updated_at
    }
```

### ERD Notes

- **USERS → TRANSACTIONS:** One-to-Many. Satu kasir bisa membuat banyak transaksi.
- **USERS → VOID_LOGS:** One-to-Many. Satu kasir bisa melakukan banyak return/void.
- **CATEGORIES → PRODUCTS:** One-to-Many. Satu kategori bisa memiliki banyak produk. `ON DELETE SET NULL`.
- **PRODUCTS → PRODUCT_PRICES:** One-to-Many. Satu produk bisa memiliki banyak price tiers. `ON DELETE CASCADE`.
- **PRODUCTS → TRANSACTION_ITEMS:** One-to-Many. Satu produk bisa muncul di banyak transaksi. `ON DELETE RESTRICT`.
- **PRODUCTS → VOID_LOGS:** One-to-Many. Satu produk bisa di-return banyak kali. `ON DELETE RESTRICT`.
- **PRODUCTS → STOCK_ALERTS:** One-to-Many. Satu produk bisa memiliki banyak alert histori. `ON DELETE CASCADE`.
- **TRANSACTIONS → TRANSACTION_ITEMS:** One-to-Many. Satu transaksi bisa memiliki banyak item. `ON DELETE CASCADE`.
- **TRANSACTIONS → VOID_LOGS:** One-to-Many. Satu transaksi bisa memiliki banyak return log. `ON DELETE CASCADE`.

---

## 2. Data Dictionary Matrix

### 2.1 Table: `users`

| Column Name     | Data Type                | Constraints                   | Description           |
| :-------------- | :----------------------- | :---------------------------- | :-------------------- |
| `id`            | UUID                     | PK, DEFAULT gen_random_uuid() | Primary key user      |
| `name`          | VARCHAR(100)             | NOT NULL                      | Nama lengkap user     |
| `email`         | VARCHAR(150)             | NOT NULL, UNIQUE              | Email login           |
| `password_hash` | VARCHAR(255)             | NOT NULL                      | Hash bcrypt password  |
| `role`          | ENUM('owner', 'cashier') | NOT NULL, DEFAULT 'cashier'   | Peran akses           |
| `phone`         | VARCHAR(20)              | NULL                          | Nomor telepon         |
| `is_active`     | BOOLEAN                  | NOT NULL, DEFAULT true        | Status aktif          |
| `created_at`    | TIMESTAMP                | NOT NULL, DEFAULT NOW()       | Waktu registrasi      |
| `updated_at`    | TIMESTAMP                | NOT NULL, DEFAULT NOW()       | Waktu update terakhir |
| `deleted_at`    | TIMESTAMP                | NULL                          | Soft delete (GORM)    |

### 2.2 Table: `categories`

| Column Name     | Data Type   | Constraints                   | Description                  |
| :-------------- | :---------- | :---------------------------- | :--------------------------- |
| `id`            | UUID        | PK, DEFAULT gen_random_uuid() | Primary key kategori         |
| `name`          | VARCHAR(50) | NOT NULL, UNIQUE              | Nama kategori                |
| `color`         | VARCHAR(7)  | NOT NULL, DEFAULT '#6B7280'   | Kode warna HEX               |
| `description`   | TEXT        | NULL                          | Deskripsi kategori           |
| `product_count` | INTEGER     | NOT NULL, DEFAULT 0           | Jumlah produk (denormalized) |
| `created_at`    | TIMESTAMP   | NOT NULL, DEFAULT NOW()       | Waktu pembuatan              |
| `updated_at`    | TIMESTAMP   | NOT NULL, DEFAULT NOW()       | Waktu update                 |

### 2.3 Table: `products`

| Column Name           | Data Type                    | Constraints                             | Description             |
| :-------------------- | :--------------------------- | :-------------------------------------- | :---------------------- |
| `id`                  | UUID                         | PK, DEFAULT gen_random_uuid()           | Primary key produk      |
| `name`                | VARCHAR(150)                 | NOT NULL                                | Nama produk             |
| `sku`                 | VARCHAR(50)                  | NOT NULL, UNIQUE                        | Stock Keeping Unit      |
| `barcode`             | VARCHAR(50)                  | UNIQUE, NULL                            | Kode barcode            |
| `category_id`         | UUID                         | FK -> categories.id, ON DELETE SET NULL | Kategori produk         |
| `base_price`          | INTEGER                      | NOT NULL, DEFAULT 0                     | Harga dasar (rupiah)    |
| `stock`               | INTEGER                      | NOT NULL, DEFAULT 0                     | Stok aktual             |
| `min_stock_threshold` | INTEGER                      | NOT NULL, DEFAULT 0                     | Batas minimal stok      |
| `unit`                | VARCHAR(20)                  | NOT NULL, DEFAULT 'pcs'                 | Satuan (pcs, kg, liter) |
| `description`         | TEXT                         | NULL                                    | Deskripsi produk        |
| `is_active`           | BOOLEAN                      | NOT NULL, DEFAULT true                  | Status aktif            |
| `stock_status`        | ENUM('SAFE', 'LOW', 'MINUS') | NOT NULL, DEFAULT 'SAFE'                | Status stok (computed)  |
| `created_at`          | TIMESTAMP                    | NOT NULL, DEFAULT NOW()                 | Waktu pembuatan         |
| `updated_at`          | TIMESTAMP                    | NOT NULL, DEFAULT NOW()                 | Waktu update            |
| `deleted_at`          | TIMESTAMP                    | NULL                                    | Soft delete             |

### 2.4 Table: `product_prices`

| Column Name  | Data Type   | Constraints                          | Description                       |
| :----------- | :---------- | :----------------------------------- | :-------------------------------- |
| `id`         | UUID        | PK, DEFAULT gen_random_uuid()        | Primary key tier                  |
| `product_id` | UUID        | FK -> products.id, ON DELETE CASCADE | Produk terkait                    |
| `min_qty`    | INTEGER     | NOT NULL, CHECK (min_qty > 1)        | Minimal kuantitas                 |
| `price`      | INTEGER     | NOT NULL, CHECK (price > 0)          | Harga tier (rupiah)               |
| `label`      | VARCHAR(50) | NOT NULL                             | Label tier (e.g., "Grosir 10pcs") |
| `created_at` | TIMESTAMP   | NOT NULL, DEFAULT NOW()              | Waktu pembuatan                   |
| `updated_at` | TIMESTAMP   | NOT NULL, DEFAULT NOW()              | Waktu update                      |

### 2.5 Table: `transactions`

| Column Name        | Data Type                                 | Constraints                   | Description                   |
| :----------------- | :---------------------------------------- | :---------------------------- | :---------------------------- |
| `id`               | UUID                                      | PK, DEFAULT gen_random_uuid() | Primary key transaksi         |
| `transaction_code` | VARCHAR(20)                               | NOT NULL, UNIQUE              | Kode unik (TRX-YYMMDD-XXXX)   |
| `type`             | ENUM('direct', 'hold')                    | NOT NULL                      | Tipe transaksi                |
| `status`           | ENUM('PENDING', 'COMPLETED', 'CANCELLED') | NOT NULL, DEFAULT 'PENDING'   | Status transaksi              |
| `cashier_id`       | UUID                                      | FK -> users.id, NOT NULL      | Kasir yang membuat            |
| `hold_note`        | VARCHAR(100)                              | NULL                          | Catatan hold bill             |
| `subtotal`         | INTEGER                                   | NOT NULL, DEFAULT 0           | Subtotal sebelum diskon/pajak |
| `discount_amount`  | INTEGER                                   | NOT NULL, DEFAULT 0           | Total diskon                  |
| `tax_amount`       | INTEGER                                   | NOT NULL, DEFAULT 0           | Total pajak                   |
| `total`            | INTEGER                                   | NOT NULL, DEFAULT 0           | Total akhir                   |
| `payment_method`   | ENUM('cash', 'qris', 'transfer')          | NULL                          | Metode pembayaran             |
| `amount_paid`      | INTEGER                                   | NULL                          | Jumlah dibayar                |
| `change_amount`    | INTEGER                                   | NULL                          | Kembalian                     |
| `completed_at`     | TIMESTAMP                                 | NULL                          | Waktu penyelesaian            |
| `created_at`       | TIMESTAMP                                 | NOT NULL, DEFAULT NOW()       | Waktu pembuatan               |
| `updated_at`       | TIMESTAMP                                 | NOT NULL, DEFAULT NOW()       | Waktu update                  |

### 2.6 Table: `transaction_items`

| Column Name      | Data Type | Constraints                              | Description                 |
| :--------------- | :-------- | :--------------------------------------- | :-------------------------- |
| `id`             | UUID      | PK, DEFAULT gen_random_uuid()            | Primary key item            |
| `transaction_id` | UUID      | FK -> transactions.id, ON DELETE CASCADE | Transaksi terkait           |
| `product_id`     | UUID      | FK -> products.id, ON DELETE RESTRICT    | Produk terkait              |
| `qty`            | INTEGER   | NOT NULL, CHECK (qty > 0)                | Kuantitas                   |
| `unit_price`     | INTEGER   | NOT NULL                                 | Harga satuan saat transaksi |
| `subtotal`       | INTEGER   | NOT NULL                                 | Subtotal item               |
| `notes`          | TEXT      | NULL                                     | Catatan item                |
| `created_at`     | TIMESTAMP | NOT NULL, DEFAULT NOW()                  | Waktu pembuatan             |

### 2.7 Table: `void_logs`

| Column Name           | Data Type | Constraints                                   | Description          |
| :-------------------- | :-------- | :-------------------------------------------- | :------------------- |
| `id`                  | UUID      | PK, DEFAULT gen_random_uuid()                 | Primary key log      |
| `transaction_id`      | UUID      | FK -> transactions.id, ON DELETE CASCADE      | Transaksi terkait    |
| `transaction_item_id` | UUID      | FK -> transaction_items.id, ON DELETE CASCADE | Item terkait         |
| `cashier_id`          | UUID      | FK -> users.id, NOT NULL                      | Kasir yang melakukan |
| `product_id`          | UUID      | FK -> products.id, ON DELETE RESTRICT         | Produk terkait       |
| `qty_returned`        | INTEGER   | NOT NULL, CHECK (qty_returned > 0)            | Jumlah dikembalikan  |
| `refund_amount`       | INTEGER   | NOT NULL                                      | Nominal refund       |
| `reason`              | TEXT      | NOT NULL                                      | Alasan return        |
| `created_at`          | TIMESTAMP | NOT NULL, DEFAULT NOW()                       | Waktu return         |
| `updated_at`          | TIMESTAMP | NOT NULL, DEFAULT NOW()                       | Waktu update         |

### 2.8 Table: `stock_alerts`

| Column Name      | Data Type            | Constraints                          | Description            |
| :--------------- | :------------------- | :----------------------------------- | :--------------------- |
| `id`             | UUID                 | PK, DEFAULT gen_random_uuid()        | Primary key alert      |
| `product_id`     | UUID                 | FK -> products.id, ON DELETE CASCADE | Produk terkait         |
| `alert_type`     | ENUM('LOW', 'MINUS') | NOT NULL                             | Tipe alert             |
| `stock_at_alert` | INTEGER              | NOT NULL                             | Stok saat alert dibuat |
| `is_resolved`    | BOOLEAN              | NOT NULL, DEFAULT false              | Status resolved        |
| `resolved_at`    | TIMESTAMP            | NULL                                 | Waktu resolve          |
| `created_at`     | TIMESTAMP            | NOT NULL, DEFAULT NOW()              | Waktu alert            |
| `updated_at`     | TIMESTAMP            | NOT NULL, DEFAULT NOW()              | Waktu update           |

---

## 3. Database Triggers

### 3.1 Stock Status Auto-Update Trigger

```sql
CREATE OR REPLACE FUNCTION update_stock_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.stock < 0 THEN
        NEW.stock_status := 'MINUS';
        -- Insert stock alert MINUS (hanya jika belum ada alert aktif)
        INSERT INTO stock_alerts (product_id, alert_type, stock_at_alert)
        SELECT NEW.id, 'MINUS', NEW.stock
        WHERE NOT EXISTS (
            SELECT 1 FROM stock_alerts 
            WHERE product_id = NEW.id 
            AND alert_type = 'MINUS' 
            AND is_resolved = false
        );
    ELSIF NEW.stock <= NEW.min_stock_threshold THEN
        NEW.stock_status := 'LOW';
        -- Insert stock alert LOW (hanya jika belum ada alert aktif)
        INSERT INTO stock_alerts (product_id, alert_type, stock_at_alert)
        SELECT NEW.id, 'LOW', NEW.stock
        WHERE NOT EXISTS (
            SELECT 1 FROM stock_alerts 
            WHERE product_id = NEW.id 
            AND alert_type = 'LOW' 
            AND is_resolved = false
        );
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

CREATE TRIGGER trigger_update_stock_status
BEFORE INSERT OR UPDATE OF stock ON products
FOR EACH ROW
EXECUTE FUNCTION update_stock_status();
```

### 3.2 Product Count Auto-Update Trigger

```sql
CREATE OR REPLACE FUNCTION update_category_product_count()
RETURNS TRIGGER AS $$
BEGIN
    -- On INSERT
    IF TG_OP = 'INSERT' THEN
        IF NEW.category_id IS NOT NULL AND NEW.deleted_at IS NULL THEN
            UPDATE categories SET product_count = product_count + 1 WHERE id = NEW.category_id;
        END IF;
        RETURN NEW;

    -- On UPDATE
    ELSIF TG_OP = 'UPDATE' THEN
        -- Category changed
        IF NEW.category_id IS DISTINCT FROM OLD.category_id THEN
            -- Decrement old category
            IF OLD.category_id IS NOT NULL THEN
                UPDATE categories SET product_count = product_count - 1 WHERE id = OLD.category_id;
            END IF;
            -- Increment new category
            IF NEW.category_id IS NOT NULL AND NEW.deleted_at IS NULL THEN
                UPDATE categories SET product_count = product_count + 1 WHERE id = NEW.category_id;
            END IF;
        END IF;

        -- Soft delete / restore
        IF NEW.deleted_at IS DISTINCT FROM OLD.deleted_at THEN
            IF NEW.deleted_at IS NOT NULL AND NEW.category_id IS NOT NULL THEN
                UPDATE categories SET product_count = product_count - 1 WHERE id = NEW.category_id;
            ELSIF NEW.deleted_at IS NULL AND NEW.category_id IS NOT NULL THEN
                UPDATE categories SET product_count = product_count + 1 WHERE id = NEW.category_id;
            END IF;
        END IF;
        RETURN NEW;

    -- On DELETE (hard delete - should not happen with soft delete)
    ELSIF TG_OP = 'DELETE' THEN
        IF OLD.category_id IS NOT NULL THEN
            UPDATE categories SET product_count = product_count - 1 WHERE id = OLD.category_id;
        END IF;
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_product_count
AFTER INSERT OR UPDATE OR DELETE ON products
FOR EACH ROW
EXECUTE FUNCTION update_category_product_count();
```

### 3.3 Transaction Code Generation Function

```sql
CREATE OR REPLACE FUNCTION generate_transaction_code()
RETURNS TRIGGER AS $$
DECLARE
    date_part VARCHAR(6);
    seq_num INTEGER;
    new_code VARCHAR(20);
BEGIN
    date_part := TO_CHAR(NEW.created_at, 'YYMMDD');

    -- Get next sequence number for today
    SELECT COALESCE(MAX(CAST(SUBSTRING(transaction_code FROM 12) AS INTEGER)), 0) + 1
    INTO seq_num
    FROM transactions
    WHERE transaction_code LIKE 'TRX-' || date_part || '-%';

    new_code := 'TRX-' || date_part || '-' || LPAD(seq_num::TEXT, 4, '0');
    NEW.transaction_code := new_code;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_generate_transaction_code
BEFORE INSERT ON transactions
FOR EACH ROW
EXECUTE FUNCTION generate_transaction_code();
```

---

## 4. Performance Optimization Plan

### 4.1 Indexing Strategy

| Table               | Index Name                  | Columns                    | Type           | Purpose                 |
| :------------------ | :-------------------------- | :------------------------- | :------------- | :---------------------- |
| `users`             | `idx_users_email`           | `email`                    | B-Tree, UNIQUE | Fast login lookup       |
| `users`             | `idx_users_role`            | `role`                     | B-Tree         | Filter owner/cashier    |
| `users`             | `idx_users_deleted_at`      | `deleted_at`               | B-Tree         | Soft delete filter      |
| `categories`        | `idx_categories_name`       | `name`                     | B-Tree, UNIQUE | Name uniqueness         |
| `products`          | `idx_products_sku`          | `sku`                      | B-Tree, UNIQUE | SKU uniqueness          |
| `products`          | `idx_products_barcode`      | `barcode`                  | B-Tree, UNIQUE | Barcode scan            |
| `products`          | `idx_products_category`     | `category_id`              | B-Tree         | Category filter         |
| `products`          | `idx_products_stock_status` | `stock_status`             | B-Tree         | Dashboard alerts        |
| `products`          | `idx_products_active`       | `is_active`, `deleted_at`  | B-Tree         | Active product filter   |
| `products`          | `idx_products_search`       | `name`, `sku`, `barcode`   | GIN (pg_trgm)  | Fuzzy search kasir      |
| `product_prices`    | `idx_prices_product_qty`    | `product_id`, `min_qty`    | B-Tree         | Price tier lookup       |
| `transactions`      | `idx_txn_code`              | `transaction_code`         | B-Tree, UNIQUE | Code lookup             |
| `transactions`      | `idx_txn_cashier`           | `cashier_id`               | B-Tree         | Cashier performance     |
| `transactions`      | `idx_txn_status`          | `status`                   | B-Tree         | Hold bill filter        |
| `transactions`      | `idx_txn_type_status`       | `type`, `status`           | B-Tree         | Hold bill list          |
| `transactions`      | `idx_txn_created`           | `created_at`               | B-Tree         | Date range reports      |
| `transactions`      | `idx_txn_date_cashier`      | `created_at`, `cashier_id` | B-Tree         | Composite report filter |
| `transaction_items` | `idx_items_transaction`     | `transaction_id`           | B-Tree         | Transaction detail      |
| `transaction_items` | `idx_items_product`         | `product_id`               | B-Tree         | Product sales stats     |
| `void_logs`         | `idx_void_cashier`          | `cashier_id`               | B-Tree         | Audit filter            |
| `void_logs`         | `idx_void_transaction`      | `transaction_id`           | B-Tree         | Transaction void lookup |
| `void_logs`         | `idx_void_created`          | `created_at`               | B-Tree         | Date range audit        |
| `stock_alerts`      | `idx_alert_product`         | `product_id`               | B-Tree         | Product alert lookup    |
| `stock_alerts`      | `idx_alert_resolved`        | `is_resolved`              | B-Tree         | Unresolved alerts       |
| `stock_alerts`      | `idx_alert_type_created`    | `alert_type`, `created_at` | B-Tree         | Alert sorting           |

### 4.2 Optimization Policies

1. **Computed Column `stock_status`:** Kolom `stock_status` di tabel `products` di-update via database trigger setiap kali `stock` berubah, menghindari perhitungan runtime.

2. **Transaction Isolation:** Semua operasi pemotongan stok menggunakan `REPEATABLE READ` isolation level dengan `SELECT FOR UPDATE` untuk mencegah race condition:
   ```sql
   BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
   SELECT stock FROM products WHERE id = ? FOR UPDATE;
   UPDATE products SET stock = stock - ? WHERE id = ?;
   COMMIT;
   ```

3. **Connection Pooling:** `pgx` pool dengan max 25 connections, min 5 connections, max conn lifetime 1 jam.

4. **Read Replicas:** Untuk reporting (`/reports/*`), query diarahkan ke read replica PostgreSQL untuk mengurangi beban primary.

5. **Partitioning (Future):** Tabel `transactions` dan `transaction_items` bisa di-partition berdasarkan `created_at` (monthly) jika data sudah > 1 juta records.

---

## 5. Data Seeding Specifications

### 5.1 Seed Script Structure

```go
// cmd/seed/main.go
package main

func main() {
    // 1. Seed Owner Account
    owner := User{
        Name:     "Andi Wijaya",
        Email:    "andi@warungku.com",
        Password: bcryptHash("owner12345"),
        Role:     "owner",
        IsActive: true,
    }

    // 2. Seed Categories
    categories := []Category{
        {Name: "Makanan", Color: "#FF6B6B"},
        {Name: "Minuman", Color: "#4ECDC4"},
        {Name: "Snack", Color: "#FFE66D"},
        {Name: "Rokok", Color: "#95E1D3"},
    }

    // 3. Seed Products (UMKM Template - FR-09)
    products := []Product{
        {
            Name: "Es Teh Manis", SKU: "ETM-001", Barcode: "8991234567001",
            BasePrice: 4000, Stock: 50, MinStockThreshold: 10, Unit: "pcs",
            CategoryID: catMinuman, IsActive: true,
        },
        {
            Name: "Teh Kotak 250ml", SKU: "TK-001", Barcode: "8991234567002",
            BasePrice: 3000, Stock: 100, MinStockThreshold: 20, Unit: "pcs",
            CategoryID: catMinuman, IsActive: true,
            PriceTiers: []ProductPrice{
                {MinQty: 10, Price: 2500, Label: "Grosir 10pcs"},
                {MinQty: 50, Price: 2200, Label: "Grosir 50pcs"},
            },
        },
        {
            Name: "Nasi Goreng Ayam", SKU: "NGA-001", Barcode: "8991234567003",
            BasePrice: 30000, Stock: 20, MinStockThreshold: 5, Unit: "porsi",
            CategoryID: catMakanan, IsActive: true,
        },
        {
            Name: "Air Mineral 600ml", SKU: "AM-001", Barcode: "8991234567004",
            BasePrice: 3500, Stock: 80, MinStockThreshold: 15, Unit: "pcs",
            CategoryID: catMinuman, IsActive: true,
        },
        {
            Name: "Gudang Garam", SKU: "GG-001", Barcode: "8991234567005",
            BasePrice: 25000, Stock: 30, MinStockThreshold: 5, Unit: "bungkus",
            CategoryID: catRokok, IsActive: true,
        },
    }

    // 4. Seed Cashier Account
    cashier := User{
        Name:     "Siti Aminah",
        Email:    "siti@warungku.com",
        Password: bcryptHash("kasir12345"),
        Role:     "cashier",
        IsActive: true,
    }
}
```

### 5.2 Mock Transaction Data (Development Only)

```go
// Generate 100 mock transactions for dashboard testing
for i := 0; i < 100; i++ {
    tx := Transaction{
        TransactionCode: fmt.Sprintf("TRX-%06d", i+1),
        Type:           "direct",
        Status:         "COMPLETED",
        CashierID:      cashier.ID,
        Subtotal:       rand.Intn(50000) + 10000,
        Total:          0, // calculated
        PaymentMethod:  "cash",
        AmountPaid:     0, // calculated
        ChangeAmount:   0, // calculated
        CreatedAt:      time.Now().AddDate(0, 0, -rand.Intn(30)),
    }
    tx.Total = tx.Subtotal
    tx.AmountPaid = tx.Total + rand.Intn(10000)
    tx.ChangeAmount = tx.AmountPaid - tx.Total
}
```

---

## 6. Migration Strategy

### 6.1 Migration Files Structure

```
migrations/
├── 001_create_users_table.sql
├── 002_create_categories_table.sql
├── 003_create_products_table.sql
├── 004_create_product_prices_table.sql
├── 005_create_transactions_table.sql
├── 006_create_transaction_items_table.sql
├── 007_create_void_logs_table.sql
├── 008_create_stock_alerts_table.sql
├── 009_create_triggers.sql
├── 010_create_indexes.sql
└── 011_seed_data.sql
```

### 6.2 Migration Tool

Menggunakan `golang-migrate/migrate` atau `pressly/goose` untuk version-controlled database migrations.

```bash
# Install goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# Create new migration
goose create add_user_phone_column sql

# Run migrations
goose postgres "postgres://user:pass@localhost/smartcost?sslmode=disable" up

# Rollback
goose postgres "postgres://user:pass@localhost/smartcost?sslmode=disable" down
```

---

**Dokumen ini merupakan spesifikasi database lengkap untuk Smart Cost V1.**
**Versi:** 1.1
**Terakhir Diperbarui:** Juli 2026
