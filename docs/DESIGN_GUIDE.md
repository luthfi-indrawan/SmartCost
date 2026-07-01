# Design Guide (DESIGN_GUIDE) - Smart Cost V1

**Versi:** 1.1  
**Status:** Regenerated - WA/Tokopedia Style  
**Design Philosophy:** Clean, Simple, Familiar, Mobile-First

---

## 1. Design Philosophy

Smart Cost mengadopsi prinsip desain yang familiar dan mudah digunakan seperti aplikasi chat (WhatsApp) dan e-commerce (Tokopedia) yang sering digunakan oleh target user UMKM di Indonesia.

**Core Principles:**
- **Familiarity:** User tidak perlu belajar ulang. Pattern UI mengikuti aplikasi yang sudah mereka gunakan sehari-hari.
- **Simplicity:** Minimum cognitive load. Setiap screen punya satu tujuan utama.
- **Speed:** Akses cepat ke fitur yang paling sering digunakan (checkout, lihat stok, laporan).
- **Clarity:** Informasi penting (harga, stok, alert) selalu terlihat jelas tanpa perlu tap/scroll.

---

## 2. Design Tokens

### 2.1 Color Palette

#### Primary Colors (WhatsApp Green Inspired)

| Token                 | Hex       | Usage                                      |
| :-------------------- | :-------- | :----------------------------------------- |
| `--color-primary-50`  | `#E8F5E9` | Light backgrounds, hover states            |
| `--color-primary-100` | `#C8E6C9` | Subtle highlights                          |
| `--color-primary-200` | `#A5D6A7` | Borders, dividers                          |
| `--color-primary-300` | `#81C784` | Disabled primary elements                  |
| `--color-primary-400` | `#66BB6A` | Secondary accents                          |
| `--color-primary-500` | `#4CAF50` | **Primary brand color** (WA Green)         |
| `--color-primary-600` | `#43A047` | Primary hover                              |
| `--color-primary-700` | `#388E3C` | Primary active                             |
| `--color-primary-800` | `#2E7D32` | Deep backgrounds                           |
| `--color-primary-900` | `#1B5E20` | Text on light backgrounds                  |

#### Secondary Colors (Tokopedia Orange Inspired)

| Token                   | Hex       | Usage                                      |
| :---------------------- | :-------- | :----------------------------------------- |
| `--color-secondary-50`  | `#FFF3E0` | Light backgrounds, hover states            |
| `--color-secondary-100` | `#FFE0B2` | Subtle highlights                          |
| `--color-secondary-500` | `#FF9800` | **Secondary brand color** (Tokopedia)    |
| `--color-secondary-600` | `#F57C00` | Secondary hover                            |
| `--color-secondary-700` | `#E65100` | Secondary active                           |

#### Semantic Colors

| Token                 | Hex       | Usage                                    |
| :-------------------- | :-------- | :--------------------------------------- |
| `--color-success-500` | `#22C55E` | Success states, completed transactions   |
| `--color-success-50`  | `#F0FDF4` | Success background                       |
| `--color-danger-500`  | `#EF4444` | Errors, stock minus alerts, void actions |
| `--color-danger-50`   | `#FEF2F2` | Danger background                        |
| `--color-warning-500` | `#F59E0B` | Warnings, low stock indicators           |
| `--color-warning-50`  | `#FFFBEB` | Warning background                       |
| `--color-info-500`    | `#3B82F6` | Informational, hold bill badges          |
| `--color-info-50`     | `#EFF6FF` | Info background                          |

#### Neutral Colors (Clean & Minimal)

| Token                 | Hex       | Usage                               |
| :-------------------- | :-------- | :---------------------------------- |
| `--color-neutral-50`  | `#F9FAFB` | Page background (light gray tint)   |
| `--color-neutral-100` | `#F3F4F6` | Card backgrounds, input backgrounds |
| `--color-neutral-200` | `#E5E7EB` | Borders, dividers                   |
| `--color-neutral-300` | `#D1D5DB` | Disabled borders                    |
| `--color-neutral-400` | `#9CA3AF` | Placeholder text                    |
| `--color-neutral-500` | `#6B7280` | Secondary text                      |
| `--color-neutral-600` | `#4B5563` | Body text                           |
| `--color-neutral-700` | `#374151` | Headings                            |
| `--color-neutral-800` | `#1F2937` | Primary text                        |
| `--color-neutral-900` | `#111827` | Deep text, labels                   |

#### WhatsApp/Tokopedia Specific

| Token                    | Hex       | Usage                                      |
| :----------------------- | :-------- | :----------------------------------------- |
| `--color-wa-bg`          | `#F0F2F5` | Main app background (WA chat bg feel)      |
| `--color-wa-surface`     | `#FFFFFF` | Card/surface background                    |
| `--color-wa-chat-bubble` | `#DCF8C6` | Success/positive indicator (WA bubble)     |
| `--color-tokped-orange`  | `#FF5C00` | CTA highlight, promo badges                |
| `--color-tokped-green`   | `#03AC0E` | Success, available stock                   |

### 2.2 Typography Hierarchy

**Font Family:** `Inter` (Google Fonts) - weights: 400, 500, 600, 700

> **Rationale:** Inter adalah font yang sangat readable di layar kecil (smartphone kasir) dan sudah familiar karena digunakan oleh banyak aplikasi modern termasuk WhatsApp Web.

| Level           | Size            | Weight | Line Height | Letter Spacing | Usage                      |
| :-------------- | :-------------- | :----- | :---------- | :------------- | :------------------------- |
| **H1**          | 28px (1.75rem)  | 700    | 1.2         | -0.02em        | Page titles (Dashboard)    |
| **H2**          | 22px (1.375rem) | 600    | 1.3         | -0.01em        | Section headers            |
| **H3**          | 18px (1.125rem) | 600    | 1.4         | 0              | Card titles, modal headers |
| **H4**          | 16px (1rem)     | 600    | 1.4         | 0              | Sub-section headers        |
| **Body**        | 15px (0.9375rem)| 400    | 1.5         | 0              | Paragraphs, descriptions   |
| **Body Small**  | 13px (0.8125rem)| 400    | 1.5         | 0              | Secondary text, metadata |
| **Caption**     | 11px (0.6875rem)| 500    | 1.4         | 0.01em         | Labels, timestamps, badges |
| **Price**       | 18px (1.125rem) | 700    | 1.2         | -0.01em        | Price displays, totals     |
| **Price Large** | 26px (1.625rem) | 700    | 1.1         | -0.02em        | Grand total kasir          |

> **WA/Tokopedia Style Note:** Ukuran font lebih kecil dan compact untuk menghemat space di layar mobile. Whitespace yang cukup tetap dijaga agar tidak terlalu crowded.

### 2.3 Spacing Scale

| Token        | Value | Usage                        |
| :----------- | :---- | :--------------------------- |
| `--space-1`  | 4px   | Tight gaps, icon padding     |
| `--space-2`  | 8px   | Inline elements, small gaps  |
| `--space-3`  | 12px  | Button padding vertical      |
| `--space-4`  | 16px  | Card padding, standard gap   |
| `--space-5`  | 20px  | Section padding              |
| `--space-6`  | 24px  | Modal padding, form sections |
| `--space-8`  | 32px  | Page padding desktop         |
| `--space-10` | 40px  | Large section gaps           |
| `--space-12` | 48px  | Hero sections                |

### 2.4 Border Radius

| Token           | Value  | Usage                                      |
| :-------------- | :----- | :----------------------------------------- |
| `--radius-sm`   | 6px    | Inputs, small buttons, tags                |
| `--radius-md`   | 10px   | Cards, modals, standard buttons            |
| `--radius-lg`   | 14px   | Large cards, panels, product cards         |
| `--radius-xl`   | 18px   | Mobile bottom sheets, floating cards     |
| `--radius-full` | 9999px | Pills, avatars, badges, FAB                |

> **WA/Tokopedia Style Note:** Border radius lebih rounded (10-14px) untuk kesan friendly dan modern, mirip bubble chat dan card di WA.

---

## 3. Component Specifications

### 3.1 Primary Button (WA Green Style)

```
+--------------------------+
|  [Icon]  Label           |
+--------------------------+
```

| State        | Background | Text        | Border | Shadow                            | Transform          |
| :----------- | :--------- | :---------- | :----- | :-------------------------------- | :----------------- |
| **Default**  | `#4CAF50`  | `#FFFFFF`   | none   | `0 1px 3px rgba(76,175,80,0.3)`  | none               |
| **Hover**    | `#43A047`  | `#FFFFFF`   | none   | `0 4px 12px rgba(76,175,80,0.4)`  | `translateY(-1px)` |
| **Active**   | `#388E3C`  | `#FFFFFF`   | none   | `0 1px 2px rgba(76,175,80,0.3)`  | `translateY(0)`    |
| **Disabled** | `#A5D6A7`  | `#E8F5E9`   | none   | none                              | none               |
| **Loading**  | `#4CAF50`  | transparent | none   | none                              | none               |

**Specs:** Padding `12px 24px`, Font 15px/600, Border-radius `10px`, Transition `all 200ms ease`

### 3.2 Secondary Button (Tokopedia Orange Style)

| State        | Background | Text        | Border | Shadow                             |
| :----------- | :--------- | :---------- | :----- | :--------------------------------- |
| **Default**  | `#FF9800`  | `#FFFFFF`   | none   | `0 1px 3px rgba(255,152,0,0.3)`   |
| **Hover**    | `#F57C00`  | `#FFFFFF`   | none   | `0 4px 12px rgba(255,152,0,0.4)`   |
| **Active**   | `#E65100`  | `#FFFFFF`   | none   | `0 1px 2px rgba(255,152,0,0.3)`   |
| **Disabled** | `#FFE0B2`  | `#FFF3E0`   | none   | none                               |

### 3.3 Ghost Button (WA Style - Minimal)

| State        | Background | Text        | Border              |
| :----------- | :--------- | :---------- | :------------------ |
| **Default**  | transparent| `#4CAF50`   | `1px solid #4CAF50` |
| **Hover**    | `#E8F5E9`  | `#43A047`   | `1px solid #43A047` |
| **Active**   | `#C8E6C9`  | `#388E3C`   | `1px solid #388E3C` |

> **Usage:** Untuk aksi sekunder seperti "Batal", "Kembali", "Hold Bill".

### 3.4 Danger Button (Void/Return)

| State        | Background | Text      | Border |
| :----------- | :--------- | :-------- | :----- |
| **Default**  | `#EF4444`  | `#FFFFFF` | none   |
| **Hover**    | `#DC2626`  | `#FFFFFF` | none   |
| **Active**   | `#B91C1C`  | `#FFFFFF` | none   |
| **Disabled** | `#FCA5A5`  | `#FEE2E2` | none   |

### 3.5 Product Card (Kasir Grid - Tokopedia Style)

```
+------------------------+
|                        |
|   [Product Image/Icon] |
|                        |
+------------------------+
| Product Name           |
| Rp 30.000     [Stock]  |
+------------------------+
```

| State               | Border              | Background | Shadow                             | Stock Badge         |
| :------------------ | :------------------ | :--------- | :--------------------------------- | :------------------ |
| **Default**         | `1px solid #E5E7EB` | `#FFFFFF`  | `0 1px 3px rgba(0,0,0,0.05)`       | Gray `#6B7280`      |
| **Hover**           | `1px solid #4CAF50` | `#E8F5E9`  | `0 4px 12px rgba(76,175,80,0.15)`  | -                   |
| **Active/Selected** | `2px solid #4CAF50` | `#C8E6C9`  | `0 2px 8px rgba(76,175,80,0.2)`    | -                   |
| **Low Stock**       | `1px solid #F59E0B` | `#FFFBEB`  | none                               | Orange `#F59E0B`    |
| **Minus Stock**     | `1px solid #EF4444` | `#FEF2F2`  | none                               | Red `#EF4444` pulse |
| **Disabled**        | `1px solid #E5E7EB` | `#F9FAFB`  | none                               | Gray, opacity 0.5   |

**Specs:** Border-radius `14px`, Padding `12px`, Image aspect ratio `1:1` dengan border-radius `10px`.

> **WA/Tokopedia Style:** Card lebih rounded, shadow lebih soft, hover state dengan green tint (seperti hover di WA chat list).

### 3.6 Input Field (Clean Minimal Style)

| State        | Border              | Background | Ring                              |
| :----------- | :------------------ | :--------- | :-------------------------------- |
| **Default**  | `1px solid #E5E7EB` | `#FFFFFF`  | none                              |
| **Hover**    | `1px solid #D1D5DB` | `#FFFFFF`  | none                              |
| **Focus**    | `1px solid #4CAF50` | `#FFFFFF`  | `0 0 0 3px rgba(76,175,80,0.15)` |
| **Error**    | `1px solid #EF4444` | `#FEF2F2`  | `0 0 0 3px rgba(239,68,68,0.15)`  |
| **Disabled** | `1px solid #E5E7EB` | `#F3F4F6`  | none                              |

**Specs:** Height `48px`, Padding `12px 16px`, Font 15px/400, Border-radius `10px`

### 3.7 Search Bar (WA Style)

```
+------------------------------------------+
| [Q]  Cari produk...                 [X]  |
+------------------------------------------+
```

| Property     | Value                                    |
| :----------- | :--------------------------------------- |
| Background   | `#F0F2F5` (WA chat list bg)              |
| Border       | none                                     |
| Border-radius| `9999px` (pill shape)                    |
| Height       | `44px`                                   |
| Padding      | `0 16px`                                 |
| Icon         | Search icon `#9CA3AF`, 20px              |
| Placeholder  | `#9CA3AF`, 14px                          |
| Focus        | Background `#FFFFFF`, ring green       |

### 3.8 Stock Alert Badge

| Status    | Background | Text      | Icon           | Animation           |
| :-------- | :--------- | :-------- | :------------- | :------------------ |
| **SAFE**  | `#E8F5E9`  | `#1B5E20` | Check circle   | none                |
| **LOW**   | `#FFFBEB`  | `#92400E` | Alert triangle | none                |
| **MINUS** | `#FEF2F2`  | `#991B1B` | Alert octagon  | `pulse 2s infinite` |

### 3.9 Hold Bill Card (WA Chat Bubble Style)

```
+-----------------------------+
| Meja 5 / Andi    45 menit   |
| 3 items        Rp 90.000    |
| [Lanjutkan]  [Batalkan]     |
+-----------------------------+
```

| State           | Border              | Background | Timer Color |
| :-------------- | :------------------ | :--------- | :---------- |
| **< 30 menit**  | `1px solid #4CAF50` | `#E8F5E9`  | `#4CAF50`   |
| **30-60 menit** | `1px solid #FF9800` | `#FFF3E0`  | `#FF9800`   |
| **> 60 menit**  | `1px solid #EF4444` | `#FEF2F2`  | `#EF4444`   |

**Specs:** Border-radius `14px`, Padding `14px`, Shadow `0 1px 3px rgba(0,0,0,0.08)`.

> **WA Style:** Card menyerupai chat bubble dengan rounded corners yang besar, mirip pesan di WA.

### 3.10 Bottom Navigation (WA Style)

```
+----------------------------------+
| [Beranda]  [Kasir]  [Laporan]   |
+----------------------------------+
```

| Property        | Value                                    |
| :-------------- | :--------------------------------------- |
| Background      | `#FFFFFF`                                |
| Border-top      | `1px solid #E5E7EB`                      |
| Height          | `60px`                                   |
| Active Icon     | `#4CAF50`, 24px                          |
| Active Label    | `#4CAF50`, 11px, 500                     |
| Inactive Icon   | `#9CA3AF`, 24px                          |
| Inactive Label  | `#9CA3AF`, 11px, 400                     |
| Badge (alert)   | `#EF4444` dot, 8px, top-right            |

### 3.11 Floating Action Button (FAB - WA Style)

```
    +------+
    |  +   |
    +------+
```

| Property     | Value                                    |
| :----------- | :--------------------------------------- |
| Background   | `#4CAF50`                                |
| Icon         | `#FFFFFF`, 24px                          |
| Size         | `56px x 56px`                            |
| Border-radius| `9999px` (circle)                        |
| Shadow       | `0 4px 12px rgba(76,175,80,0.4)`        |
| Hover        | `#43A047`, shadow lebih besar            |

> **Usage:** FAB untuk aksi utama di halaman kasir (tambah item cepat, bayar).

---

## 4. Layout Patterns

### 4.1 Cashier Page Layout (Mobile-First)

```
Mobile (< 768px):
+------------------+
| [=] SMART COST   |  <- Header minimal
+------------------+
| [Q] Cari...      |  <- Pill search bar
+------------------+
| [Semua][Makanan] |  <- Horizontal scroll tabs
+------------------+
|                  |
|  Product Grid    |
|  (2 columns)     |
|  [Card] [Card]   |
|  [Card] [Card]   |
|                  |
+------------------+
| [Cart FAB]       |  <- Floating cart button
+------------------+

Tablet (768px+):
+------------------+----------+
| [=] SMART COST   | KERANJANG|
+------------------+          |
| [Q] Cari...      | Items... |
+------------------+          |
| [Tabs]           | Total    |
+------------------+ [Bayar]  |
|                  |          |
|  Product Grid    |          |
|  (3-4 columns)   |          |
|                  |          |
+------------------+----------+
```

### 4.2 Owner Dashboard Layout (Mobile-First)

```
Mobile (< 768px):
+------------------+
| [=] SMART COST   |
| Selamat datang,  |
| Andi!            |
+------------------+
| PENDAPATAN HARI  |
| Rp 1.545.000     |  <- Large price display
+------------------+
| [Grafik Bulanan] |
+------------------+
| PERFORMA KASIR   |
| - Siti: Rp945K   |
| - Budi: Rp600K   |
+------------------+
| STOK ALERT (2)   |  <- Red badge
| Teh Kotak: -3    |
| Es Teh: 2        |
+------------------+
| [Beranda][Produk]|  <- Bottom nav
| [Laporan][Akun]  |
+------------------+

Desktop (1024px+):
+------+---------------------------+
|      | PENDAPATAN    | GRAFIK    |
| Side | Rp 1.545.000  | [Chart]   |
| bar  |---------------------------|
| Nav  | PERFORMA KASIR            |
|      |---------------------------|
|      | STOK ALERT   | TRANSAKSI  |
|      | [List]       | [Recent]   |
+------+---------------------------+
```

### 4.3 Transaction Detail / Receipt Style

```
+------------------+
| STRUK PENJUALAN  |
| TRX-200624-0001  |
| 20 Juni 2024     |
| Kasir: Siti      |
+------------------+
| 1. Nasi Goreng   |
|    x2 @30.000    |
|    60.000        |
| 2. Es Teh        |
|    x2 @4.000     |
|    8.000         |
+------------------+
| Subtotal  68.000 |
| Total     68.000 |
| Tunai    100.000 |
| Kembali   32.000 |
+------------------+
| [Cetak Struk]    |
+------------------+
```

> **Style:** Mirip struk kasir fisik dengan font monospace untuk angka, border dashed untuk pemisah section.

---

## 5. Responsive Breakpoints

| Breakpoint | Width           | Target Device        | Layout                                |
| :--------- | :-------------- | :------------------- | :------------------------------------ |
| **xs**     | < 480px         | Smartphone portrait  | Single column, bottom nav             |
| **sm**     | 480px - 767px   | Smartphone landscape | Single column, bottom nav             |
| **md**     | 768px - 1023px  | Tablet               | Two column (products + cart)          |
| **lg**     | 1024px - 1279px | Small desktop        | Two column, sidebar nav               |
| **xl**     | >= 1280px       | Desktop              | Three column, sidebar + main + detail |

---

## 6. Animation & Micro-interactions

### 6.1 Page Transitions

| Transition | Duration | Easing     | Usage                      |
| :--------- | :------- | :--------- | :------------------------- |
| **Slide In** | 300ms   | ease-out   | Modal, bottom sheet        |
| **Fade**     | 200ms   | ease-in-out| Toast, alert, overlay      |
| **Scale**    | 150ms   | ease-out   | Button press, card select  |

### 6.2 Feedback Patterns

| Action           | Feedback                              | Duration |
| :--------------- | :------------------------------------ | :------- |
| **Add to cart**  | Card scales down 0.95 then back       | 150ms    |
| **Remove item**  | Item slides out left + fade           | 200ms    |
| **Checkout**     | Screen transitions to receipt         | 300ms    |
| **Error**        | Input shakes + red border pulse       | 300ms    |
| **Success**      | Green checkmark scale + toast slide   | 300ms    |
| **Hold Bill**    | Card slides up from bottom            | 300ms    |

### 6.3 Loading States

| Component    | Loading Style                          |
| :----------- | :------------------------------------- |
| **Button**   | Spinner 16px white, replace label      |
| **Card**     | Skeleton pulse with rounded corners    |
| **List**     | Skeleton rows, 3-5 items               |
| **Page**     | Full screen spinner with logo        |
| **Image**    | Blur placeholder + fade in on load   |

---

## 7. Iconography

**Icon Library:** `lucide-react` (consistent, lightweight, modern)

| Icon Name        | Usage                          | Size  |
| :--------------- | :----------------------------- | :---- |
| `Home`           | Nav - Beranda                  | 24px  |
| `ShoppingCart`   | Nav - Kasir / Cart             | 24px  |
| `BarChart3`      | Nav - Laporan                  | 24px  |
| `User`           | Nav - Akun / Profile           | 24px  |
| `Search`         | Search bar                     | 20px  |
| `Plus`           | Add item / FAB                 | 24px  |
| `Minus`          | Decrease qty                   | 16px  |
| `Trash2`         | Delete / Void                  | 20px  |
| `Printer`        | Print receipt                  | 20px  |
| `Bell`           | Notifications / Alerts         | 24px  |
| `Package`        | Products                       | 24px  |
| `Tags`           | Categories                     | 20px  |
| `ArrowLeft`      | Back navigation                | 24px  |
| `MoreVertical`   | Context menu                   | 20px  |
| `CheckCircle`    | Success / Safe stock           | 20px  |
| `AlertTriangle`  | Warning / Low stock            | 20px  |
| `AlertOctagon`   | Danger / Minus stock           | 20px  |
| `Clock`          | Hold bill timer                | 16px  |
| `Receipt`        | Transaction history            | 20px  |

---

## 8. Accessibility Guidelines

- **Minimum Touch Target:** 44px x 44px untuk semua interactive elements.
- **Color Contrast:** Minimum 4.5:1 untuk body text, 3:1 untuk large text.
- **Focus States:** Semua interactive elements memiliki visible focus ring (`ring-2 ring-primary-500 ring-offset-2`).
- **Screen Reader:** Semua icon buttons memiliki `aria-label`, semua images memiliki `alt` text.
- **Reduced Motion:** Respect `prefers-reduced-motion` untuk users yang sensitive terhadap animasi.

---

**Dokumen ini merupakan panduan desain lengkap untuk Smart Cost V1.**
**Versi:** 1.1
**Terakhir Diperbarui:** Juli 2026
