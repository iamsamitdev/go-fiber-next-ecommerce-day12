# 🛒 Fiber E-commerce API

ระบบ E-commerce API ที่สมบูรณ์แบบ พัฒนาด้วย Go Fiber framework โดยใช้หลักการ Clean Architecture รองรับการจัดการสินค้า, ตะกร้าสินค้า, คำสั่งซื้อ, การชำระเงิน และระบบสถิติ พร้อมการ deploy ด้วย **Docker**

## 🏗️ Architecture

โปรเจ็กต์นี้ใช้หลักการ **Clean Architecture** มีโครงสร้างดังนี้:

```
go-backend/
├── Dockerfile                      # Docker configuration
├── cmd/
│   ├── api/                        # Application entry point
│   │   └── main.go
│   └── migrate/                    # Database migration CLI
│       └── main.go
├── internal/
│   ├── adapters/              # External adapters
│   │   ├── http/              # HTTP layer (handlers, middleware, routes)
│   │   │   ├── handlers/      # HTTP request handlers
│   │   │   │   ├── auth_handler.go       # Authentication & Users
│   │   │   │   ├── user_handler.go       # User management (Admin)
│   │   │   │   ├── category_handler.go   # Category management
│   │   │   │   ├── product_handler.go    # Product management
│   │   │   │   ├── cart_handler.go       # Shopping cart
│   │   │   │   ├── order_handler.go      # Order management
│   │   │   │   ├── payment_handler.go    # Payment processing
│   │   │   │   └── stats_handler.go      # Statistics & Analytics
│   │   │   ├── middleware/    # HTTP middleware
│   │   │   │   └── auth_middleware.go    # JWT & Role-based Auth
│   │   │   └── routes/        # Route definitions
│   │   │       └── routes.go
│   │   └── persistence/       # Database layer
│   │       ├── models/        # Database models (GORM)
│   │       │   └── user.go
│   │       └── repositories/  # Data access layer
│   │           ├── user_repository.go
│   │           ├── category_repository.go
│   │           ├── product_repository.go
│   │           ├── cart_repository.go
│   │           ├── order_repository.go
│   │           ├── transaction_repository.go
│   │           ├── role_repository.go
│   │           └── stats_repository.go
│   ├── config/                # Configuration management
│   │   ├── config.go          # App configuration
│   │   ├── database.go        # Database setup & migration
│   │   └── seeder.go          # Database seeding
│   └── core/                  # Business logic core
│       ├── domain/            # Domain entities and interfaces
│       │   ├── entities/      # Business entities
│   │   │   └── user.go
│   │   └── ports/         # Interfaces (ports)
│   │       ├── repositories/  # Repository interfaces
│   │       │   └── user_repository.go
│   │       └── services/      # Service interfaces
│   │           └── auth_service.go
│   └── services/          # Business logic services
│       ├── auth_service.go
│       ├── user_service.go
│       ├── category_service.go
│       ├── product_service.go
│       ├── cart_service.go
│       ├── order_service.go
│       ├── payment_service.go
│       └── stats_service.go
├── pkg/utils/                 # Shared utilities
│   ├── jwt.go                 # JWT utilities
│   ├── password.go            # Password hashing
│   └── validator.go           # Validation utilities
├── docs/                      # API documentation (Swagger)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── .env.example              # Environment variables template
├── API_ENDPOINTS.md          # API endpoints documentation
├── SEEDER_GUIDE.md          # Database seeding guide
├── go.mod                   # Go modules
└── go.sum                   # Go modules checksum
```

## 🚀 Features

### 🔐 Authentication & Authorization
- **User Registration & Login** พร้อม JWT tokens
- **Admin Registration** (เฉพาะ Admin เท่านั้น)
- **Role-based Access Control** (Admin, User)
- **Password Management** (Change, Forgot, Reset)
- **Refresh Token Support**
- **Logout System**

### 🛍️ E-commerce Core Features
- **📦 Category Management** (CRUD operations, Pagination)
- **🛒 Product Management** (CRUD, Search, Filter by category/price)
- **🛍️ Shopping Cart** (Add, Update, Remove, Clear items)
- **📋 Order Management** (Create, View, Cancel, Status tracking)
- **💳 Payment Processing** (Create, Verify, Cancel payments)
- **📊 Statistics & Analytics** (Sales, Products, Users stats)

### 👥 User Management
- **User CRUD Operations** (Admin only)
- **Profile Management**
- **User Statistics**
- **Pagination Support**

### 🛡️ Security Features
- **JWT Token-based Authentication**
- **Password Hashing** (bcrypt)
- **Input Validation** (comprehensive)
- **Role-based Route Protection**
- **CORS Support**

### 🗄️ Database Features
- **PostgreSQL Integration**
- **Auto Migration System**
- **Database Seeding** (10 Categories + 20 Products)
- **Admin User Auto-creation**

### � Docker Features
- **Containerized Development** (Hot reload support)
- **Production-ready Docker configuration**
- **Multi-stage Docker builds**
- **Health checks & monitoring**
- **Environment-specific configurations**

### �📚 Documentation & Development
- **Swagger API Documentation** (ภาษาไทย)
- **Hot Reload Development** (Air)
- **Docker Support**
- **Comprehensive API Endpoints**

## 📋 Prerequisites

- **Go** 1.21+
- **PostgreSQL** 15+ (หรือใช้ Docker)
- **Docker & Docker Compose** (แนะนำ)

## 🛠️ Installation & Setup

### Option 1: Docker Setup (แนะนำ)

#### 1. Clone the repository
```bash
git clone <repository-url>
cd go-fiber-next-ecommerce
```

#### 2. Environment Configuration
สร้างไฟล์ environment variables:
```bash
# Copy environment template
cp .env.example .env.development  # For development
cp .env.example .env              # For production

# Edit environment variables in .env.development or .env
```

#### 3. Docker Development Setup
```bash
# Start all services (backend, frontend, database)
docker compose --env-file .env.development up --build -d

# View logs
docker compose --env-file .env.development logs -f go-backend

# Access backend
# API: http://localhost:4001
# Swagger: http://localhost:4001/swagger/
```

#### 4. Docker Production Setup
```bash
# Start production services
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View logs
docker compose -f docker-compose.prod.yml logs -f go-backend
```

#### 5. Stop Services
```bash
# Development
docker compose --env-file .env.development down --volumes

# Production
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Option 2: Manual Setup

#### 1. Clone the repository
```bash
git clone <repository-url>
cd go-fiber-next-ecommerce/go-backend
```

#### 2. Install dependencies
```bash
go mod download
```

#### 3. Install Development Tools
```bash
# Install Air for hot reloading
go install github.com/cosmtrek/air@latest

# Install Swagger generator
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 4. Environment Configuration
สร้างไฟล์ `.env` ตามตัวอย่าง:
```env
# 🌐 Environment
APP_ENV=development
APP_PORT=4001
APP_URL=http://localhost:4001

# 📦 Database (PostgreSQL)
DB_HOST=localhost  # Use 'postgres' for Docker
DB_PORT=5432
DB_NAME=ecommerce_dev
DB_USER=ecommerce_user
DB_PASS=SecureDB#Pass2024!
DB_SSL=disable

# 🔐 JWT Config
JWT_SECRET=fibernextcommerce_development_jwt_secret_key_2024_secure_minimum_32_chars
JWT_EXPIRES_IN=24h

# 🔄 Database Migration
AUTO_MIGRATE=true

# 👑 Admin User Seeding (Optional)
ADMIN_EMAIL=admin@email.com
ADMIN_PASSWORD=AdminSecure#Pass2024!
ADMIN_FIRST_NAME=System
ADMIN_LAST_NAME=Administrator
```

#### 5. Database Setup

#### Option A: Using Docker (แนะนำ)
```bash
# Using Docker Compose from root directory
cd ..  # Go back to root project directory
docker compose --env-file .env.development up -d postgres

# View PostgreSQL logs
docker compose --env-file .env.development logs -f postgres
```

#### Option B: Manual PostgreSQL Setup
สร้าง PostgreSQL database ด้วยข้อมูลใน `.env` file

#### 6. Run the Application

#### Development (with hot reload)
```bash
# Make sure you're in go-backend directory
cd go-backend  # if not already there

# Run with hot reload
air

# หรือรันโดยตรง
go run cmd/api/main.go
```

🚀 **API จะรันที่**: `http://localhost:4001`
📚 **Swagger Docs**: `http://localhost:4001/swagger/`

#### Production
```bash
# Build และรัน
go build -o bin/api cmd/api/main.go
./bin/api
```

## 🐳 Docker Commands

### Development Environment
```bash
# Start backend with all services (from root directory)
docker compose --env-file .env.development up --build -d

# Start only backend service
docker compose --env-file .env.development up --build go-backend -d

# View backend logs
docker compose --env-file .env.development logs -f go-backend

# Execute commands in backend container
docker compose --env-file .env.development exec go-backend sh

# Restart backend service
docker compose --env-file .env.development restart go-backend

# Stop all services
docker compose --env-file .env.development down --volumes
```

### Production Environment
```bash
# Start production backend
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View production logs
docker compose -f docker-compose.prod.yml logs -f go-backend

# Stop production services
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Backend-only Docker Commands
```bash
# Build backend image only
docker build -t go-ecommerce-backend ./go-backend

# Run backend container standalone
docker run -p 4001:4001 --env-file .env go-ecommerce-backend

# Run with environment variables
docker run -p 4001:4001 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_NAME=ecommerce_dev \
  go-ecommerce-backend
```

🚀 **API จะรันที่**: `http://localhost:3000`

## 🗄️ Database Management

### Migration & Seeding

#### Using Docker
```bash
# Manual migration และ seeding in Docker container
docker compose --env-file .env.development exec go-backend go run cmd/migrate/main.go

# Or run migration when starting container (AUTO_MIGRATE=true in .env)
```

#### Manual Setup
```bash
# Manual migration และ seeding
go run cmd/migrate/main.go

# Auto-migration (development)
AUTO_MIGRATE=true go run cmd/api/main.go

# Disable auto-migration (production)
AUTO_MIGRATE=false go run cmd/api/main.go
```

### 🌱 E-commerce Data Seeding
ระบบจะสร้างข้อมูลตัวอย่างสำหรับ E-commerce อัตโนมัติ:

**ข้อมูลที่จะถูกสร้าง:**
- **Admin User** (ถ้าตั้งค่าใน .env)
- **2 Roles** (admin, user)
- **10 หมวดหมู่สินค้า** (Electronics, Fashion, Home & Garden, Sports & Outdoors, Books & Media, Health & Beauty, Toys & Games, Automotive, Food & Beverages, Office Supplies)
- **20 สินค้าตัวอย่าง** พร้อมรูปภาพและข้อมูลที่สมจริง

**ตัวอย่างสินค้าที่สร้าง:**
- iPhone 15 Pro (฿42,900) - Electronics
- MacBook Air M2 (฿35,900) - Electronics
- เสื้อเชิ้ตผ้าคอตตอน (฿1,290) - Fashion
- กางเกงยีนส์ (฿890) - Fashion
- โซฟาผ้า 3 ที่นั่ง (฿15,900) - Home & Garden
- รองเท้าวิ่ง Nike (฿3,290) - Sports & Outdoors

> 📖 **ดูรายละเอียดเพิ่มเติม:** [SEEDER_GUIDE.md](./SEEDER_GUIDE.md)

**Log Messages ตัวอย่าง:**
```bash
✅ Roles seeded successfully
✅ Admin user created successfully: admin@email.com
✅ Categories seeded successfully (10 categories)
✅ Products seeded successfully (20 products)
```

## 📚 API Documentation

เข้าถึง Swagger UI documentation ได้ที่: 

**Development**: `http://localhost:4001/swagger/`  
**Production**: `http://localhost:4001/swagger/` (หรือตาม domain ที่ตั้งค่า)

### 🛣️ Available Endpoints

#### 🔐 Authentication
- `POST /api/v1/auth/register` - สมัครสมาชิกใหม่
- `POST /api/v1/auth/login` - เข้าสู่ระบบ
- `POST /api/v1/auth/refresh` - รีเฟรช token
- `POST /api/v1/auth/logout` - ออกจากระบบ (Protected)
- `POST /api/v1/auth/change-password` - เปลี่ยนรหัสผ่าน (Protected)
- `POST /api/v1/auth/forgot-password` - ลืมรหัสผ่าน
- `POST /api/v1/auth/reset-password` - รีเซ็ตรหัสผ่าน
- `POST /api/v1/auth/admin/register` - สร้าง Admin ใหม่ (Admin only)

#### 👥 User Management (Admin only)
- `GET /api/v1/users` - ดูผู้ใช้ทั้งหมด
- `GET /api/v1/users/{id}` - ดูผู้ใช้ตาม ID
- `PUT /api/v1/users/{id}` - แก้ไขข้อมูลผู้ใช้
- `DELETE /api/v1/users/{id}` - ลบผู้ใช้

#### 📦 Categories
- `GET /api/v1/categories` - ดูหมวดหมู่ทั้งหมด (Public)
- `GET /api/v1/categories/{id}` - ดูหมวดหมู่ตาม ID (Public)
- `POST /api/v1/categories` - สร้างหมวดหมู่ (Admin only)
- `PUT /api/v1/categories/{id}` - แก้ไขหมวดหมู่ (Admin only)
- `DELETE /api/v1/categories/{id}` - ลบหมวดหมู่ (Admin only)

#### 🛒 Products
- `GET /api/v1/products` - ดูสินค้าทั้งหมด (Public)
- `GET /api/v1/products/{id}` - ดูสินค้าตาม ID (Public)
- `GET /api/v1/products/category/{categoryId}` - ดูสินค้าตามหมวดหมู่ (Public)
- `GET /api/v1/products/search` - ค้นหาสินค้า (Public)
- `POST /api/v1/products` - สร้างสินค้า (Admin only)
- `PUT /api/v1/products/{id}` - แก้ไขสินค้า (Admin only)
- `DELETE /api/v1/products/{id}` - ลบสินค้า (Admin only)

#### 🛍️ Shopping Cart (User only)
- `GET /api/v1/cart` - ดูตะกร้าสินค้า
- `POST /api/v1/cart` - เพิ่มสินค้าลงตะกร้า
- `PUT /api/v1/cart/{itemId}` - อัพเดทสินค้าในตะกร้า
- `DELETE /api/v1/cart/{itemId}` - ลบสินค้าจากตะกร้า
- `DELETE /api/v1/cart` - ล้างตะกร้าสินค้า

#### 📋 Orders (User for own orders, Admin for all)
- `POST /api/v1/orders` - สร้างคำสั่งซื้อ
- `GET /api/v1/orders` - ดูคำสั่งซื้อของตัวเอง
- `GET /api/v1/orders/{id}` - ดูคำสั่งซื้อตาม ID
- `PUT /api/v1/orders/{id}/cancel` - ยกเลิกคำสั่งซื้อ
- `GET /api/v1/orders/admin` - ดูคำสั่งซื้อทั้งหมด (Admin only)
- `PUT /api/v1/orders/admin/{id}/status` - อัพเดทสถานะคำสั่งซื้อ (Admin only)

#### 💳 Payments (User only)
- `POST /api/v1/payments` - สร้างการชำระเงิน
- `POST /api/v1/payments/{id}/verify` - ยืนยันการชำระเงิน
- `PUT /api/v1/payments/{id}/cancel` - ยกเลิกการชำระเงิน

#### 📊 Statistics (Admin only)
- `GET /api/v1/stats/sales` - ดูสถิติการขาย
- `GET /api/v1/stats/products` - ดูสถิติสินค้า
- `GET /api/v1/stats/users` - ดูสถิติผู้ใช้

> 📖 **ดูรายละเอียดเพิ่มเติม:** [API_ENDPOINTS.md](./API_ENDPOINTS.md)

## 🔐 Authentication Flow

1. **Register**: สร้างบัญชีผู้ใช้ใหม่ (role = "user")
2. **Login**: เข้าสู่ระบบเพื่อรับ JWT token และ refresh token
3. **Protected Routes**: ใส่ JWT token ใน `Authorization` header เป็น `Bearer <token>`
4. **Admin Routes**: ต้องมี role = "admin"

### 💡 Example Requests

#### 📝 Register User
```bash
curl -X POST http://localhost:4001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "0812345678",
    "address": "123 Main St"
  }'
```

#### 🔑 Login
```bash
curl -X POST http://localhost:4001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### 🛒 Add Product to Cart
```bash
curl -X POST http://localhost:4001/api/v1/cart \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "product_id": "product-uuid",
    "quantity": 2
  }'
```

#### 📋 Create Order
```bash
curl -X POST http://localhost:4001/api/v1/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "shipping_address": "123 Main St, Bangkok",
    "payment_method": "credit_card"
  }'
```

## 🏛️ Project Structure Details

### 🎯 Core Entities
- `User` - ผู้ใช้พร้อม roles และข้อมูลส่วนตัว
- `Category` - หมวดหมู่สินค้า
- `Product` - สินค้าพร้อมรูปภาพและข้อมูลรายละเอียด
- `Cart` & `CartItem` - ตะกร้าสินค้าและรายการสินค้า
- `Order` & `OrderItem` - คำสั่งซื้อและรายการสินค้าที่สั่ง
- `Transaction` - การชำระเงิน
- `Role` - บทบาทผู้ใช้

### ⚙️ Services (Business Logic)
- `AuthService` - Authentication และ Authorization
- `UserService` - การจัดการผู้ใช้
- `CategoryService` - การจัดการหมวดหมู่
- `ProductService` - การจัดการสินค้า
- `CartService` - การจัดการตะกร้าสินค้า
- `OrderService` - การจัดการคำสั่งซื้อ
- `PaymentService` - การจัดการการชำระเงิน
- `StatsService` - การจัดการสถิติ

### 🗄️ Repositories (Data Access)
- `UserRepository` - การเข้าถึงข้อมูลผู้ใช้
- `CategoryRepository` - การเข้าถึงข้อมูลหมวดหมู่
- `ProductRepository` - การเข้าถึงข้อมูลสินค้า
- `CartRepository` - การเข้าถึงข้อมูลตะกร้า
- `OrderRepository` - การเข้าถึงข้อมูลคำสั่งซื้อ
- `TransactionRepository` - การเข้าถึงข้อมูลการชำระเงิน
- `RoleRepository` - การเข้าถึงข้อมูลบทบาท
- `StatsRepository` - การเข้าถึงข้อมูลสถิติ

### 🔧 Utilities
- `utils.ValidateStruct` - ตรวจสอบความถูกต้องของ struct
- `utils.HashPassword` - เข้ารหัสรหัสผ่าน
- `utils.CheckPassword` - ตรวจสอบรหัสผ่าน
- `utils.GenerateJWT` - สร้าง JWT token
- `utils.ParseJWT` - แปลง JWT token

### 🛡️ Middleware
- `AuthRequired()` - ตรวจสอบ JWT token
- `AdminRequired()` - ตรวจสอบ admin role
- `RoleRequired(roles...)` - ตรวจสอบ roles ที่กำหนด

### 🐳 Docker Configuration
- `Dockerfile` - Multi-stage build สำหรับ production
- `docker-compose.yml` - Development environment
- `docker-compose.prod.yml` - Production environment
- Health checks และ monitoring
- Hot reload support ใน development

## 🔧 Development Tools

### 🔥 Hot Reload
โปรเจ็กต์ใช้ [Air](https://github.com/cosmtrek/air) สำหรับ hot reloading ระหว่างการพัฒนา

#### การใช้งาน Hot Reload
```bash
# ติดตั้ง Air
go install github.com/cosmtrek/air@latest

# รันด้วย hot reload
air

# หรือใช้ Docker ที่มี hot reload support
docker compose --env-file .env.development up --build go-backend -d
```

### 📚 Swagger Documentation
สร้าง/อัพเดท Swagger documentation:
```bash
swag init -g cmd/api/main.go -o docs
```

### 🛠️ Make Commands (ถ้ามี Makefile)
```bash
make build         # Build แอปพลิเคชัน
make run           # รันแอปพลิเคชัน
make dev           # รันแบบ development mode
make test          # รัน tests
make migrate       # รัน database migration
make docker-build  # Build Docker image
make docker-run    # Run Docker container
```

## 🛠️ Tech Stack

### Backend Framework
- **Fiber v2** - Fast HTTP web framework
- **GORM** - ORM สำหรับ Go
- **PostgreSQL** - ฐานข้อมูลเชิงสัมพันธ์

### Authentication & Security
- **JWT** - JSON Web Tokens
- **bcrypt** - Password hashing
- **Validator** - Input validation

### Development Tools
- **Air** - Hot reload สำหรับ Go
- **Swagger** - API documentation
- **Docker** - Containerization

## 🐳 Docker Support

### Docker Compose (แนะนำ)
ใช้ Docker Compose files ที่มีให้ในโฟลเดอร์ root:

#### Development Environment
```bash
# Start all services (backend, frontend, database)
docker compose --env-file .env.development up --build -d

# Start only backend + database
docker compose --env-file .env.development up --build go-backend postgres -d

# View logs
docker compose --env-file .env.development logs -f go-backend

# Stop services
docker compose --env-file .env.development down --volumes
```

#### Production Environment
```bash
# Start production services
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View logs
docker compose -f docker-compose.prod.yml logs -f go-backend

# Stop services
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Standalone Docker
```bash
# Build backend image
docker build -t go-ecommerce-backend .

# Run with PostgreSQL connection
docker run -p 4001:4001 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_NAME=ecommerce_dev \
  -e DB_USER=ecommerce_user \
  -e DB_PASS=SecureDB#Pass2024! \
  go-ecommerce-backend
```

### Docker Development Features
- **Hot Reload**: โค้ดจะ reload อัตโนมัติเมื่อมีการแก้ไข
- **Volume Mounting**: โค้ดใน host จะ sync กับ container
- **Health Checks**: ตรวจสอบสถานะแอปพลิเคชัน
- **Network Isolation**: Services แยกออกจากกันอย่างปลอดภัย

## 🧪 Testing

### Manual Testing
```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Docker Testing
```bash
# Run tests in Docker container
docker compose --env-file .env.development exec go-backend go test ./...

# Run tests with coverage in Docker
docker compose --env-file .env.development exec go-backend go test -cover ./...
```

### API Testing
```bash
# Test health endpoint
curl http://localhost:4001/health

# Test API endpoints (requires running backend)
curl http://localhost:4001/api/v1/categories
curl http://localhost:4001/api/v1/products
```

## 📝 API Response Format

### ✅ Success Response
```json
{
  "success": true,
  "message": "ดึงข้อมูลสินค้าสำเร็จ",
  "data": {
    "id": "uuid",
    "name": "iPhone 15 Pro",
    "price": 42900,
    "stock": 10,
    "category_id": "uuid"
  },
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 20,
    "total_pages": 2
  }
}
```

### ❌ Error Response
```json
{
  "success": false,
  "message": "ไม่พบสินค้า",
  "error": "Product not found"
}
```

## 🚀 Deployment

### Docker Deployment (แนะนำ)

#### Production Build
```bash
# Build production image
docker build -t go-ecommerce-backend-prod .

# Or use docker-compose
docker compose -f docker-compose.prod.yml build go-backend
```

#### Environment Variables for Production
```bash
# Set production environment
export APP_ENV=production
export AUTO_MIGRATE=false
export JWT_SECRET=<strong-production-secret>
export DB_HOST=<production-db-host>
export DB_USER=<production-db-user>
export DB_PASS=<production-db-password>
```

#### Deploy with Docker Compose
```bash
# Start production environment
docker compose -f docker-compose.prod.yml --env-file .env up -d

# Check logs
docker compose -f docker-compose.prod.yml logs -f go-backend
```

### Manual Deployment

#### Production Build
```bash
# Build for production
go build -o bin/api cmd/api/main.go

# Run production binary
./bin/api
```

#### Environment Variables for Manual Setup
สำหรับ production ให้ตั้งค่า:
- `APP_ENV=production`
- `AUTO_MIGRATE=false`
- `JWT_SECRET=<strong-secret>`
- Database credentials

### Cloud Deployment Options

#### VPS/Server Deployment
```bash
# 1. Clone repository on server
git clone <repository-url>
cd go-fiber-next-ecommerce

# 2. Set up environment variables
cp .env.example .env
# Edit .env with production values

# 3. Start with Docker Compose
docker compose -f docker-compose.prod.yml --env-file .env up -d

# 4. Check logs
docker compose -f docker-compose.prod.yml logs -f
```

#### Container Registry Deployment
```bash
# 1. Build and tag image
docker build -t your-registry/go-ecommerce-backend:latest ./go-backend

# 2. Push to registry
docker push your-registry/go-ecommerce-backend:latest

# 3. Deploy on target server
docker run -d -p 4001:4001 \
  --env-file .env \
  your-registry/go-ecommerce-backend:latest
```

## 📄 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📞 Support

หากมีปัญหาหรือข้อสงสัย สามารถติดต่อได้ที่:
- Email: support@example.com
- GitHub Issues: [Create an issue](https://github.com/your-repo/issues)

---

⭐ **ถ้าโปรเจ็กต์นี้มีประโยชน์ กรุณาให้ Star ด้วยนะครับ!**
