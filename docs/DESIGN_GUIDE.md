# Design Guide (DESIGN_GUIDE) - Smart Cost V1

**Versi:** 1.0  
**Status:** Draft Disetujui  
**Design Tool:** Figma (link akan ditambahkan)

---

## 1. Design Artifact Links

| Artifact                        | URL                                           | Status      |
| :------------------------------ | :-------------------------------------------- | :---------- |
| **Figma - UI Kit & Components** | `https://figma.com/file/smart-cost-ui-kit`    | In Progress |
| **Figma - Cashier Mobile Flow** | `https://figma.com/file/smart-cost-cashier`   | In Progress |
| **Figma - Owner Dashboard**     | `https://figma.com/file/smart-cost-dashboard` | In Progress |
| **Whimsical - User Flow**       | `https://whimsical.com/smart-cost-flow`       | Completed   |

---

## 2. Design Tokens Identity

### 2.1 Color Palette

#### Primary Colors

| Token                 | Hex       | Usage                           |
| :-------------------- | :-------- | :------------------------------ |
| `--color-primary-50`  | `#EEF2FF` | Light backgrounds, hover states |
| `--color-primary-100` | `#E0E7FF` | Subtle highlights               |
| `--color-primary-200` | `#C7D2FE` | Borders, dividers               |
| `--color-primary-300` | `#A5B4FC` | Disabled primary elements       |
| `--color-primary-400` | `#818CF8` | Secondary accents               |
| `--color-primary-500` | `#6366F1` | **Primary brand color**         |
| `--color-primary-600` | `#4F46E5` | Primary hover                   |
| `--color-primary-700` | `#4338CA` | Primary active                  |
| `--color-primary-800` | `#3730A3` | Deep backgrounds                |
| `--color-primary-900` | `#312E81` | Text on light backgrounds       |

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

#### Neutral Colors

| Token                 | Hex       | Usage                               |
| :-------------------- | :-------- | :---------------------------------- |
| `--color-neutral-50`  | `#F9FAFB` | Page background                     |
| `--color-neutral-100` | `#F3F4F6` | Card backgrounds, input backgrounds |
| `--color-neutral-200` | `#E5E7EB` | Borders, dividers                   |
| `--color-neutral-300` | `#D1D5DB` | Disabled borders                    |
| `--color-neutral-400` | `#9CA3AF` | Placeholder text                    |
| `--color-neutral-500` | `#6B7280` | Secondary text                      |
| `--color-neutral-600` | `#4B5563` | Body text                           |
| `--color-neutral-700` | `#374151` | Headings                            |
| `--color-neutral-800` | `#1F2937` | Primary text                        |
| `--color-neutral-900` | `#111827` | Deep text, labels                   |

### 2.2 Typography Hierarchy

**Font Family:** `Inter` (Google Fonts) - weights: 400, 500, 600, 700

| Level           | Size            | Weight | Line Height | Letter Spacing | Usage                      |
| :-------------- | :-------------- | :----- | :---------- | :------------- | :------------------------- |
| **H1**          | 32px (2rem)     | 700    | 1.2         | -0.02em        | Page titles (Dashboard)    |
| **H2**          | 24px (1.5rem)   | 600    | 1.3         | -0.01em        | Section headers            |
| **H3**          | 20px (1.25rem)  | 600    | 1.4         | 0              | Card titles, modal headers |
| **H4**          | 18px (1.125rem) | 600    | 1.4         | 0              | Sub-section headers        |
| **Body**        | 16px (1rem)     | 400    | 1.5         | 0              | Paragraphs, descriptions   |
| **Body Small**  | 14px (0.875rem) | 400    | 1.5         | 0              | Secondary text, metadata   |
| **Caption**     | 12px (0.75rem)  | 500    | 1.4         | 0.01em         | Labels, timestamps, badges |
| **Price**       | 20px (1.25rem)  | 700    | 1.2         | -0.01em        | Price displays, totals     |
| **Price Large** | 28px (1.75rem)  | 700    | 1.1         | -0.02em        | Grand total kasir          |

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

| Token           | Value  | Usage                           |
| :-------------- | :----- | :------------------------------ |
| `--radius-sm`   | 4px    | Inputs, small buttons           |
| `--radius-md`   | 8px    | Cards, modals, standard buttons |
| `--radius-lg`   | 12px   | Large cards, panels             |
| `--radius-xl`   | 16px   | Mobile bottom sheets            |
| `--radius-full` | 9999px | Pills, avatars, badges          |

---

## 3. Component State Specifications

### 3.1 Primary Button

```
+----------------------------------+
|  [Icon]  Label                   |
+----------------------------------+
```

| State        | Background | Text        | Border | Shadow                            | Transform          |
| :----------- | :--------- | :---------- | :----- | :-------------------------------- | :----------------- |
| **Default**  | `#6366F1`  | `#FFFFFF`   | none   | `0 1px 2px rgba(99,102,241,0.2)`  | none               |
| **Hover**    | `#4F46E5`  | `#FFFFFF`   | none   | `0 4px 12px rgba(99,102,241,0.3)` | `translateY(-1px)` |
| **Active**   | `#4338CA`  | `#FFFFFF`   | none   | `0 1px 2px rgba(99,102,241,0.2)`  | `translateY(0)`    |
| **Disabled** | `#A5B4FC`  | `#E0E7FF`   | none   | none                              | none               |
| **Loading**  | `#6366F1`  | transparent | none   | none                              | none               |

**Specs:** Padding `12px 24px`, Font 16px/600, Border-radius `8px`, Transition `all 200ms ease`

### 3.2 Danger Button (Void/Return)

| State        | Background | Text      | Border |
| :----------- | :--------- | :-------- | :----- |
| **Default**  | `#EF4444`  | `#FFFFFF` | none   |
| **Hover**    | `#DC2626`  | `#FFFFFF` | none   |
| **Active**   | `#B91C1C`  | `#FFFFFF` | none   |
| **Disabled** | `#FCA5A5`  | `#FEE2E2` | none   |

### 3.3 Product Card (Kasir Grid)

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
| **Hover**           | `1px solid #6366F1` | `#EEF2FF`  | `0 4px 12px rgba(99,102,241,0.15)` | -                   |
| **Active/Selected** | `2px solid #6366F1` | `#E0E7FF`  | `0 2px 8px rgba(99,102,241,0.2)`   | -                   |
| **Low Stock**       | `1px solid #F59E0B` | `#FFFBEB`  | none                               | Orange `#F59E0B`    |
| **Minus Stock**     | `1px solid #EF4444` | `#FEF2F2`  | none                               | Red `#EF4444` pulse |
| **Disabled**        | `1px solid #E5E7EB` | `#F9FAFB`  | none                               | Gray, opacity 0.5   |

### 3.4 Input Field

| State        | Border              | Background | Ring                              |
| :----------- | :------------------ | :--------- | :-------------------------------- |
| **Default**  | `1px solid #D1D5DB` | `#FFFFFF`  | none                              |
| **Hover**    | `1px solid #9CA3AF` | `#FFFFFF`  | none                              |
| **Focus**    | `1px solid #6366F1` | `#FFFFFF`  | `0 0 0 3px rgba(99,102,241,0.15)` |
| **Error**    | `1px solid #EF4444` | `#FEF2F2`  | `0 0 0 3px rgba(239,68,68,0.15)`  |
| **Disabled** | `1px solid #E5E7EB` | `#F3F4F6`  | none                              |

**Specs:** Height `44px`, Padding `12px 16px`, Font 16px/400, Border-radius `8px`

### 3.5 Stock Alert Badge

| Status    | Background | Text      | Icon           | Animation           |
| :-------- | :--------- | :-------- | :------------- | :------------------ |
| **SAFE**  | `#F0FDF4`  | `#166534` | Check circle   | none                |
| **LOW**   | `#FFFBEB`  | `#92400E` | Alert triangle | none                |
| **MINUS** | `#FEF2F2`  | `#991B1B` | Alert octagon  | `pulse 2s infinite` |

### 3.6 Hold Bill Card

```
+-----------------------------+
| 🕐 45 menit    [Meja 5/Andi] |
| 3 items        Rp 90.000    |
| [Lanjutkan]  [Batalkan]     |
+-----------------------------+
```

| State           | Border              | Background | Timer Color |
| :-------------- | :------------------ | :--------- | :---------- |
| **< 30 menit**  | `1px solid #3B82F6` | `#EFF6FF`  | `#3B82F6`   |
| **30-60 menit** | `1px solid #F59E0B` | `#FFFBEB`  | `#F59E0B`   |
| **> 60 menit**  | `1px solid #EF4444` | `#FEF2F2`  | `#EF4444`   |

---

## 4. Responsive Breakpoints

| Breakpoint | Width           | Target Device        | Layout                                |
| :--------- | :-------------- | :------------------- | :------------------------------------ |
| **xs**     | < 480px         | Smartphone portrait  | Single column, bottom nav             |
| **sm**     | 480px - 767px   | Smartphone landscape | Single column, bottom nav             |
| **md**     | 768px - 1023px  | Tablet               | Two column (products + cart)          |
| **lg**     | 1024px - 1279px | Small desktop        | Two column, sidebar nav               |
| **xl**     | >= 1280px       | Desktop              | Three column, sidebar + main + detail |

### 4.1 Cashier Page Layout Responsive

```
Mobile (< 768px):
+------------------+
| Search Bar       |
+------------------+
| Category Tabs    |
+------------------+
|                  |
|  Product Grid    |
|  (2 columns)     |
|                  |
+------------------+
| [Cart FAB]       |
+------------------+

Tablet (768px+):
+------------------+----------+
| Search Bar       | CART     |
+------------------+          |
| Category Tabs    | Items... |
+------------------+          |
|                  | Total    |
|  Product Grid    | [Pay]    |
|  (3-4 columns)   |          |
|                  |          |
+------------------+----------+
```

### 4.2 Owner Dashboard Layout Responsive

```
Mobile (< 768px):
+------------------+
| Header + Alert   |
+------------------+
| Revenue Card     |
+------------------+
| Chart (swipe)    |
+------------------+
| Cashier List     |
+------------------+
| Stock Alerts     |
+------------------+
| Bottom Nav       |
+------------------+

Desktop (1024px+):
+------+---------------------------+
|      | Revenue    | Chart        |
| Side |---------------------------|
| bar  | Cashier Performance       |
| Nav  |---------------------------|
|      | Stock Alerts | Recent Txn |
+------+---------------------------+
```
