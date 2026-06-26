# Smart Cost - Cloud POS System for UMKM

Smart Cost adalah aplikasi Point of Sale (POS) berbasis cloud yang dirancang khusus untuk UMKM F&B dan retail kecil. Sistem ini menawarkan fleksibilitas operasional tanpa memblokir transaksi saat stok minus, mendukung multi-harga grosir dalam satu produk, dan menyediakan audit trail lengkap untuk setiap aktivitas kasir.

## ✨ Value Proposition

Smart Cost menghilangkan hambatan operasional pada kasir konvensional dengan memungkinkan transaksi meskipun stok tercatat 0 atau minus, sambil tetap menjaga akuntabilitas melalui pelacakan aktivitas per kasir dan notifikasi real-time ke pemilik usaha.

## 📸 Visual Preview

> _Placeholder: Screenshot aplikasi kasir mobile dan dashboard owner akan ditambahkan setelah deployment._

---

## 📑 Documentation Matrix

| Document Name          | Core Focus                                                         | Repository Path                                       |
| :--------------------- | :----------------------------------------------------------------- | :---------------------------------------------------- |
| **PRD.md**             | Product scope, MoSCoW features, user stories & acceptance criteria | [`/docs/PRD.md`](docs//PRD.md)                        |
| **USER_FLOW.md**       | End-to-end user journeys, Mermaid diagrams, wireframe layouts      | [`/docs/USER_FLOW.md`](docs/USER_FLOW.md)             |
| **TECH_SPEC.md**       | System architecture, tech stack rationale, security strategy       | [`/docs/TECH_SPEC.md`](docs/TECH_SPEC.md)             |
| **SERVICES.md**        | API contract, endpoint specs, request/response schemas             | [`/docs/SERVICES.md`](docs/SERVICES.md)               |
| **DATABASE_DESIGN.md** | ERD, data dictionary, indexing & seeding specs                     | [`/docs/DATABASE_DESIGN.md`](docs/DATABASE_DESIGN.md) |
| **DESIGN_GUIDE.md**    | Design tokens, typography, component states                        | [`/docs/DESIGN_GUIDE.md`](docs/DESIGN_GUIDE.md)       |
| **DEPLOYMENT.md**      | Infrastructure map, CI/CD rules, migration procedures              | [`/docs/DEPLOYMENT.md`](docs/DEPLOYMENT.md)           |
| **PROJECT_PLAN.md**    | WBS, milestones, Git workflow, risk management                     | [`/docs/PROJECT_PLAN.md`](docs/PROJECT_PLAN.md)       |
| **QA_CHECKLIST.md**    | Functional verification, test scenarios, cross-browser matrix      | [`/docs/QA_CHECKLIST.md`](docs/QA_CHECKLIST.md)       |

---

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+ (optional, for session caching)

### 1. Clone Repository

```bash
git clone https://github.com/luthfi-indrawan/SmartCost.git
cd SmartCost
```

### 2. Install Backend Dependencies

```bash
cd src/backend-smartcost
go mod init backend-smartcost
go mod tidy
```

### 3. Install Frontend Dependencies

```bash
cd src/frontend-smartcost
npm install
```

### 4. Environment Synchronization

```bash
# Backend
cp src/backend-smartcost/.env.example backend/.env

# Frontend
cp src/frontend-smartcost/.env.example frontend/.env.local
```

### 5. Database Migration

```bash
cd src/backend-smartcost
go run cmd/migrate/main.go up
```

### 6. Run Local Development Server

```bash
# Terminal 1 - Backend API
cd src/backend-smartcost
go run cmd/api/main.go

# Terminal 2 - Frontend Dev Server
cd src/frontend-smartcost
npm run dev
```

- Backend API: `http://localhost:8080`
- Frontend App: `http://localhost:5173`

---

## 🔐 Environment Variables

### Backend (.env)

```env
# Server
APP_ENV=development
SERVER_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=
DB_PASSWORD=
DB_NAME=smart_cost
DB_SSL_MODE=disable

# JWT
JWT_SECRET=
JWT_EXPIRY_HOURS=24

# Redis (Optional)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

### Frontend (.env.local)

```env
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_APP_NAME=SmartCost
```

---

## 📄 License

MIT License - Smart Cost Team
