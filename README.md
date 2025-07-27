# Go Fiber Next.js E-commerce Platform

Full-stack e-commerce platform ที่ทันสมัย พัฒนาด้วย **Go Fiber** สำหรับ backend API และ **Next.js 15** สำหรับ frontend โปรเจ็กต์นี้ให้บริการระบบ e-commerce ที่สมบูรณ์แบบพร้อมกับ authentication, authorization, และฟีเจอร์ทางการค้าครบครัน รองรับการ deploy ด้วย **Docker** และ **Nginx** reverse proxy

## 🏗️ สถาปัตยกรรมระบบ

```
go-fiber-next-ecommerce/
├── docker-compose.yml               # Docker Compose for Development
├── docker-compose.prod.yml         # Docker Compose for Production
├── .env.example                     # Environment variables template
├── postgresql.conf                  # PostgreSQL configuration
├── README.md                        # Project documentation
├── go-backend/                      # Go Fiber API Backend
│   ├── Dockerfile                   # Docker configuration for backend
│   ├── cmd/                         # Entry points
│   │   ├── api/main.go             # Main API server
│   │   └── migrate/main.go         # Database migration
│   ├── docs/                       # Swagger documentation
│   │   ├── docs.go                 # Generated Swagger docs
│   │   ├── swagger.json            # Swagger JSON
│   │   └── swagger.yaml            # Swagger YAML
│   ├── internal/                   # Internal application code
│   │   ├── adapters/               # External adapters
│   │   │   ├── http/               # HTTP handlers, routes, middleware
│   │   │   │   ├── handlers/       # HTTP request handlers
│   │   │   │   ├── middleware/     # HTTP middleware
│   │   │   │   └── routes/         # API route definitions
│   │   │   └── persistence/        # Database layer
│   │   │       ├── models/         # Database models (GORM)
│   │   │       └── repositories/   # Repository pattern implementations
│   │   ├── config/                 # Configuration files
│   │   │   ├── config.go           # Application configuration
│   │   │   ├── database.go         # Database configuration
│   │   │   └── seeder.go           # Database seeding
│   │   └── core/                   # Business logic (Clean Architecture)
│   │       ├── domain/             # Domain layer
│   │       │   ├── entities/       # Business entities
│   │       │   └── ports/          # Interface definitions
│   │       └── services/           # Business services
│   │           ├── auth_service.go      # Authentication service
│   │           ├── cart_service.go      # Shopping cart service
│   │           ├── category_service.go  # Category management
│   │           ├── order_service.go     # Order processing
│   │           ├── payment_service.go   # Payment processing
│   │           ├── product_service.go   # Product management
│   │           ├── stats_service.go     # Analytics & statistics
│   │           └── user_service.go      # User management
│   └── pkg/                        # Shared utilities
│       └── utils/                  # Utility functions
│           ├── jwt.go              # JWT token handling
│           ├── password.go         # Password hashing
│           └── validator.go        # Input validation
├── next-frontend/                  # Next.js Frontend
│   ├── Dockerfile                  # Production Docker configuration
│   ├── Dockerfile.dev              # Development Docker configuration
│   ├── next.config.ts              # Next.js configuration
│   ├── package.json                # Node.js dependencies
│   ├── tsconfig.json               # TypeScript configuration
│   ├── postcss.config.mjs          # PostCSS configuration
│   ├── components.json             # Shadcn/ui components config
│   ├── public/                     # Static assets
│   │   ├── images/                 # Image assets
│   │   │   ├── arrivals/           # New arrivals images
│   │   │   ├── blog/               # Blog images
│   │   │   ├── cart/               # Shopping cart icons
│   │   │   ├── categories/         # Category images
│   │   │   ├── checkout/           # Checkout process images
│   │   │   ├── hero/               # Hero section images
│   │   │   ├── products/           # Product images
│   │   │   ├── promo/              # Promotional images
│   │   │   └── users/              # User avatar images
│   │   └── icons/                  # SVG icons
│   └── src/
│       ├── middleware.ts           # Next.js middleware
│       ├── app/                    # Next.js App Router
│       │   ├── globals.css         # Global styles
│       │   ├── (admin)/            # Admin panel routes
│       │   ├── (auth)/             # Authentication routes
│       │   │   ├── login/          # Login page
│       │   │   ├── register/       # Registration page
│       │   │   └── forgot-password/ # Password reset
│       │   └── (pages)/            # Main application routes
│       │       ├── shop/           # Product catalog
│       │       ├── cart/           # Shopping cart
│       │       ├── checkout/       # Checkout process
│       │       ├── profile/        # User profile
│       │       └── orders/         # Order history
│       ├── components/             # React components
│       │   ├── admin/              # Admin panel components
│       │   ├── auth/               # Authentication components
│       │   ├── blog/               # Blog components
│       │   ├── cart/               # Shopping cart components
│       │   ├── checkout/           # Checkout components
│       │   ├── home/               # Homepage components
│       │   ├── shop/               # E-commerce components
│       │   ├── shared/             # Shared components
│       │   └── ui/                 # Reusable UI components (Shadcn/ui)
│       ├── context/                # React Context providers
│       │   ├── CartSidebarModalContext.tsx    # Cart sidebar state
│       │   ├── PreviewSliderContext.tsx       # Product preview slider
│       │   └── QuickViewModalContext.tsx      # Quick view modal
│       ├── hooks/                  # Custom React hooks
│       │   ├── useAuth.tsx         # Authentication hook
│       │   ├── useCategory.tsx     # Category management hook
│       │   └── useProduct.tsx      # Product management hook
│       ├── lib/                    # Utility libraries
│       │   └── utils.ts            # Utility functions
│       ├── stores/                 # State management (Zustand)
│       │   ├── cartStore.ts        # Shopping cart state
│       │   ├── counterStore.ts     # General counter state
│       │   └── wishlistStore.ts    # Wishlist state
│       └── types/                  # TypeScript type definitions
│           ├── blogItem.ts         # Blog types
│           ├── category.ts         # Category types
│           ├── Menu.ts             # Navigation menu types
│           ├── product.ts          # Product types
│           └── testimonial.ts      # Testimonial types
└── nginx/                          # Nginx Reverse Proxy
    ├── nginx.conf                  # Main Nginx configuration
    ├── conf.d/                     # Nginx server configurations
    │   ├── default.conf            # Default server config
    │   └── local.conf              # Local development config
    └── logs/                       # Nginx logs
        ├── access.log              # Access logs
        └── error.log               # Error logs
```

## � Docker Deployment

โปรเจ็กต์นี้รองรับการ deploy ด้วย Docker ทั้งสำหรับ development และ production environment

### Architecture Overview
```
┌─────────────────┐    ┌──────────────────┐    ┌────────────────┐
│   Nginx Proxy   │    │  Next.js Frontend│    │   Go Backend   │
│   (Port 8800)   │◄──►│   (Port 3001)    │◄──►│   (Port 4001)  │
└─────────────────┘    └──────────────────┘    └────────────────┘
         │                        │                        │
         │                        │                        │
         └────────────────────────┼────────────────────────┼──────┐
                                  │                        │      │
                             ┌─────────────────┐    ┌─────────────────┐
                             │   PostgreSQL    │    │    Adminer      │
                             │   (Port 5433)   │    │   (Port 8088)   │
                             └─────────────────┘    └─────────────────┘
```

### Services
- **nginx**: Reverse proxy server (production only)
- **next-frontend**: Next.js application
- **go-backend**: Go Fiber API server
- **postgres**: PostgreSQL database
- **adminer**: Database management tool (development only)

### 🚀 Quick Start with Docker

#### Prerequisites
- **Docker** 20.10+
- **Docker Compose** 2.0+

#### 1. Clone Repository
```bash
git clone <repository-url>
cd go-fiber-next-ecommerce
```

#### 2. Environment Setup
```bash
# Copy environment file
cp .env.example .env.development  # For development
cp .env.example .env              # For production

# Edit environment variables
# Update database credentials, JWT secrets, admin user, etc.
```

#### 3. Development Environment

**Build and Start Services:**
```bash
# Build all services
docker compose build

# Start all services in development mode
docker compose --env-file .env.development up --build -d

# View logs
docker compose --env-file .env.development logs -f

# Stop all services
docker compose --env-file .env.development down --volumes
```

**Access Applications:**
- **Frontend**: http://localhost:3001
- **Backend API**: http://localhost:4001
- **Swagger Docs**: http://localhost:4001/swagger/
- **Adminer (DB Tool)**: http://localhost:8088
- **PostgreSQL**: localhost:5433

#### 4. Production Environment

**Build and Start Services:**
```bash
# Build all services for production
docker compose -f docker-compose.prod.yml build

# Start all services in production mode
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View logs
docker compose -f docker-compose.prod.yml logs -f

# Stop all services
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

**Access Applications:**
- **Nginx Proxy**: http://localhost:8800
- **Frontend**: http://localhost:3001
- **Backend API**: http://localhost:4001
- **Swagger Docs**: http://localhost:4001/swagger/
- **PostgreSQL**: localhost:5433

### 🔧 Docker Commands Reference

#### Development Commands
```bash
# Start development environment
docker compose --env-file .env.development up -d

# Rebuild specific service
docker compose --env-file .env.development up --build go-backend -d

# View service logs
docker compose --env-file .env.development logs go-backend -f

# Execute command in container
docker compose --env-file .env.development exec go-backend sh

# Restart specific service
docker compose --env-file .env.development restart next-frontend

# Remove all containers and volumes
docker compose --env-file .env.development down --volumes --remove-orphans
```

#### Production Commands
```bash
# Start production environment
docker compose -f docker-compose.prod.yml --env-file .env up -d

# Scale services (if needed)
docker compose -f docker-compose.prod.yml --env-file .env up --scale go-backend=2 -d

# Monitor resource usage
docker stats

# Clean up unused resources
docker system prune -a
```

#### Database Commands
```bash
# Backup database
docker compose exec postgres pg_dump -U ecommerce_user ecommerce_dev > backup.sql

# Restore database
docker compose exec -T postgres psql -U ecommerce_user ecommerce_dev < backup.sql

# Access PostgreSQL shell
docker compose exec postgres psql -U ecommerce_user -d ecommerce_dev
```

### 📋 Environment Variables

#### Required Environment Variables
```env
# Application
NODE_ENV=development|production
APP_ENV=development|production
FRONTEND_PORT=3001
APP_URL=http://localhost:4001
NEXT_PUBLIC_API_URL=http://localhost:4001

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=ecommerce_user
DB_PASS=SecureDB#Pass2024!
DB_NAME=ecommerce_dev

# JWT
JWT_SECRET=your_super_secure_jwt_secret_minimum_32_characters
JWT_EXPIRES_IN=24h

# Admin User
ADMIN_EMAIL=admin@your-domain.com
ADMIN_PASSWORD=AdminSecure#Pass2024!
ADMIN_FIRST_NAME=Admin
ADMIN_LAST_NAME=Administrator

# Migration
AUTO_MIGRATE=true  # Set to false in production
```

### 🔍 Troubleshooting

#### Common Issues

**Port Already in Use:**
```bash
# Check what's using the port
netstat -tulpn | grep :3001

# Kill process using port
sudo kill -9 $(sudo lsof -t -i:3001)
```

**Database Connection Issues:**
```bash
# Check database container logs
docker compose logs postgres -f

# Recreate database container
docker compose down postgres
docker volume rm ecommerce_postgres_data_dev
docker compose up postgres -d
```

**Container Build Failures:**
```bash
# Clean Docker cache
docker builder prune -a

# Rebuild without cache
docker compose build --no-cache

# Remove all containers and images
docker compose down --rmi all --volumes --remove-orphans
```

### 🔐 Security Notes

- Change default passwords in production
- Use strong JWT secrets
- Configure SSL/TLS for production deployments
- Regularly update Docker images
- Monitor logs for security issues
- Use environment-specific configuration files

## 🚀 ฟีเจอร์หลัก

### Backend (Go Fiber)
- **🔐 Authentication & Authorization**: JWT-based auth with role-based access control
- **👥 User Management**: User registration, profile management, admin controls
- **🛍️ Product Management**: CRUD operations for products with categories
- **🛒 Shopping Cart**: Cart management with persistent storage
- **📝 Order Management**: Order processing, status tracking
- **💳 Payment Integration**: Payment processing capabilities
- **📊 Dashboard & Analytics**: Sales statistics and reporting
- **📚 API Documentation**: Swagger/OpenAPI documentation
- **🗄️ Database**: PostgreSQL with GORM ORM
- **🔄 Database Migration**: Automated schema management
- **🌱 Data Seeding**: Sample data for development
- **🐳 Docker Support**: Containerized deployment
- **🔧 Health Checks**: Application monitoring endpoints

### Frontend (Next.js)
- **🎨 Modern UI/UX**: Responsive design with Tailwind CSS v4
- **🔐 Authentication Flow**: Login, register, password reset
- **🛍️ Product Catalog**: Product browsing, search, filtering
- **🛒 Shopping Cart**: Interactive cart with real-time updates
- **💳 Checkout Process**: Complete order placement flow
- **👤 User Dashboard**: Profile management, order history
- **👨‍💼 Admin Panel**: Product and order management
- **📱 Mobile-First**: Responsive design for all devices
- **⚡ Performance**: Optimized images, fonts, and loading
- **🎭 Animations**: Smooth transitions and micro-interactions
- **🔔 Notifications**: Real-time feedback with toast messages
- **🐳 Docker Support**: Containerized development and production

## 🛠️ Technology Stack

### Backend
- **Framework**: [Go Fiber v2](https://gofiber.io/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Authentication**: [JWT (golang-jwt)](https://github.com/golang-jwt/jwt)
- **Validation**: [Go Playground Validator](https://github.com/go-playground/validator)
- **Documentation**: [Swagger](https://swagger.io/)
- **Environment**: [Viper](https://github.com/spf13/viper)
- **UUID**: [Google UUID](https://github.com/google/uuid)
- **Containerization**: [Docker](https://www.docker.com/)

### Frontend
- **Framework**: [Next.js 15](https://nextjs.org/) with App Router
- **Language**: [TypeScript](https://www.typescriptlang.org/)
- **Styling**: [Tailwind CSS v4](https://tailwindcss.com/)
- **UI Components**: [Radix UI](https://www.radix-ui.com/) + [Shadcn/ui](https://ui.shadcn.com/)
- **Forms**: [React Hook Form](https://react-hook-form.com/) + [Zod](https://zod.dev/)
- **HTTP Client**: [Axios](https://axios-http.com/)
- **State Management**: [Zustand](https://zustand-demo.pmnd.rs/) + React Context
- **Icons**: [Lucide React](https://lucide.dev/)
- **Animations**: [Framer Motion](https://www.framer.com/motion/)
- **Sliders**: [Swiper.js](https://swiperjs.com/)
- **Containerization**: [Docker](https://www.docker.com/)

### Infrastructure
- **Reverse Proxy**: [Nginx](https://nginx.org/)
- **Container Orchestration**: [Docker Compose](https://docs.docker.com/compose/)
- **Database Admin**: [Adminer](https://www.adminer.org/)
- **Development**: Hot reload with file watching

## 🚀 Quick Start

### Prerequisites
- **Go** 1.24+ 
- **Node.js** 18+
- **PostgreSQL** 13+ (or use Docker)
- **Docker & Docker Compose** (recommended)
- **Git**

### Option 1: Docker Deployment (Recommended)

#### Development Environment
```bash
# 1. Clone repository
git clone <repository-url>
cd go-fiber-next-ecommerce

# 2. Setup environment
cp .env.example .env.development
# Edit .env.development with your configuration

# 3. Start all services
docker compose --env-file .env.development up --build -d

# 4. Access applications
# Frontend: http://localhost:3001
# Backend: http://localhost:4001
# Swagger: http://localhost:4001/swagger/
# Adminer: http://localhost:8088
```

#### Production Environment
```bash
# 1. Setup production environment
cp .env.example .env
# Edit .env with production configuration

# 2. Start production services
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# 3. Access applications
# Nginx Proxy: http://localhost:8800
# Direct Frontend: http://localhost:3001
# Backend API: http://localhost:4001
```

#### Stop Services
```bash
# Development
docker compose --env-file .env.development down --volumes

# Production
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Option 2: Manual Setup

#### 1. Clone Repository
```bash
git clone <repository-url>
cd go-fiber-next-ecommerce
```

#### 2. Backend Setup
```bash
cd go-backend

# Install dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# Run database migrations
go run cmd/migrate/main.go

# Seed sample data (optional)
go run cmd/migrate/main.go -seed

# Start the API server
go run cmd/api/main.go
```

Backend จะรันที่: `http://localhost:8080`
Swagger Docs: `http://localhost:8080/swagger/`

#### 3. Frontend Setup
```bash
cd next-frontend

# Install dependencies
npm install

# Set up environment variables
echo "NEXT_PUBLIC_API_URL=http://localhost:8080" > .env.local

# Start development server
npm run dev
```

Frontend จะรันที่: `http://localhost:3001`

## 📡 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh token
- `POST /api/v1/auth/forgot-password` - Forgot password
- `POST /api/v1/auth/logout` - User logout

### Products & Categories
- `GET /api/v1/products` - Get all products
- `GET /api/v1/products/:id` - Get product by ID
- `GET /api/v1/categories` - Get all categories
- `POST /api/v1/products` - Create product (Admin)
- `PUT /api/v1/products/:id` - Update product (Admin)

### Shopping Cart & Orders
- `GET /api/v1/cart` - Get user cart
- `POST /api/v1/cart/add` - Add item to cart
- `POST /api/v1/orders` - Create order
- `GET /api/v1/orders` - Get user orders

### Admin Dashboard
- `GET /api/v1/admin/stats` - Dashboard statistics
- `GET /api/v1/admin/users` - Manage users
- `GET /api/v1/admin/orders` - Manage orders

*ดู [API_ENDPOINTS.md](go-backend/API_ENDPOINTS.md) สำหรับรายละเอียดครบถ้วน*

## 🗄️ Database Schema

### Core Tables
- **users** - User accounts & profiles
- **categories** - Product categories
- **products** - Product catalog
- **cart_items** - Shopping cart items
- **orders** - Order records
- **order_items** - Order line items
- **payments** - Payment transactions

### Authentication
- JWT tokens with refresh mechanism
- Role-based access control (user, admin)
- Password hashing with bcrypt

## 🔧 Configuration

### Environment Variables Template

#### Root Level (.env.development / .env)
```env
# 🌐 Application Configuration
NODE_ENV=development
APP_ENV=development
FRONTEND_PORT=3001
APP_URL=http://localhost:4001
NEXT_PUBLIC_API_URL=http://localhost:4001

# Database Configuration
DB_HOST=postgres  # Use 'localhost' for manual setup
DB_PORT=5432
DB_USER=ecommerce_user
DB_PASS=SecureDB#Pass2024!
DB_NAME=ecommerce_dev

# JWT Configuration
JWT_SECRET=fibernextcommerce_development_jwt_secret_key_2024_secure_minimum_32_chars
JWT_EXPIRES_IN=24h

# Migration Configuration
AUTO_MIGRATE=true  # Set to false in production

# Admin User Configuration
ADMIN_EMAIL=admin@your-domain.com
ADMIN_PASSWORD=AdminSecure#Pass2024!
ADMIN_FIRST_NAME=Admin
ADMIN_LAST_NAME=Administrator
```

#### Backend Environment Variables
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=ecommerce_db

# JWT
JWT_SECRET=your_jwt_secret
JWT_EXPIRES_IN=24h
JWT_REFRESH_SECRET=your_refresh_secret
JWT_REFRESH_EXPIRES_IN=7d

# Server
PORT=8080
APP_ENV=development
```

#### Frontend Environment Variables
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### Docker Environment Variables

#### Development (.env.development)
```env
NODE_ENV=development
APP_ENV=development
DB_HOST=postgres
AUTO_MIGRATE=true
```

#### Production (.env)
```env
NODE_ENV=production
APP_ENV=production
DB_HOST=postgres
AUTO_MIGRATE=false
# Add SSL configurations
# Add production secrets
```

## 📝 Available Scripts

### Docker Scripts

#### Development
```bash
# Build all services
docker compose build

# Start all services
docker compose --env-file .env.development up --build -d

# View logs (all services)
docker compose --env-file .env.development logs -f

# View logs (specific service)
docker compose --env-file .env.development logs -f go-backend

# Stop all services
docker compose --env-file .env.development down --volumes

# Restart specific service
docker compose --env-file .env.development restart next-frontend

# Execute command in container
docker compose --env-file .env.development exec go-backend sh
```

#### Production
```bash
# Build all services
docker compose -f docker-compose.prod.yml build

# Start all services
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View logs
docker compose -f docker-compose.prod.yml logs -f

# Stop all services
docker compose -f docker-compose.prod.yml --env-file .env down --volumes

# Scale services
docker compose -f docker-compose.prod.yml --env-file .env up --scale go-backend=2 -d
```

### Backend (Manual)
```bash
cd go-backend

go run cmd/api/main.go           # Start API server
go run cmd/migrate/main.go       # Run migrations
go run cmd/migrate/main.go -seed # Run migrations with seeding
go test ./...                    # Run tests
go build -o bin/api cmd/api/main.go  # Build for production

# Development with hot reload
go install github.com/cosmtrek/air@latest
air                              # Start with hot reload
```

### Frontend (Manual)
```bash
cd next-frontend

npm run dev      # Start development server (port 3001)
npm run build    # Build for production
npm run start    # Start production server
npm run lint     # Run ESLint
npm run type-check  # TypeScript check
```

### Database Management
```bash
# Docker database commands
docker compose exec postgres pg_dump -U ecommerce_user ecommerce_dev > backup.sql
docker compose exec -T postgres psql -U ecommerce_user ecommerce_dev < backup.sql
docker compose exec postgres psql -U ecommerce_user -d ecommerce_dev

# Manual database commands
psql -h localhost -p 5432 -U ecommerce_user -d ecommerce_dev
pg_dump -h localhost -p 5432 -U ecommerce_user ecommerce_dev > backup.sql
```

## 🚀 Deployment

### Docker Deployment (Recommended)

#### Development Deployment
```bash
# Build and start development environment
docker compose --env-file .env.development up --build -d

# Access services:
# - Frontend: http://localhost:3001
# - Backend: http://localhost:4001  
# - Database Admin: http://localhost:8088
```

#### Production Deployment
```bash
# Build and start production environment
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# Access services:
# - Nginx Proxy: http://localhost:8800
# - Frontend: http://localhost:3001
# - Backend: http://localhost:4001
```

#### Production with SSL (Recommended for public deployment)
1. **Update Nginx configuration for SSL**
2. **Use proper domain names**
3. **Add SSL certificates**
4. **Update environment variables**

### Manual Deployment

#### Backend Deployment
1. **Build the application**
   ```bash
   cd go-backend
   go build -o bin/api cmd/api/main.go
   ```

2. **Setup production database**
   ```bash
   # Create production database
   createdb -h your-db-host -U postgres ecommerce_prod
   
   # Run migrations
   ./bin/migrate
   ```

3. **Start the API server**
   ```bash
   ./bin/api
   ```

#### Frontend Deployment
1. **Build the application**
   ```bash
   cd next-frontend
   npm run build
   ```

2. **Deploy to Vercel** (recommended)
   ```bash
   npx vercel --prod
   ```

3. **Or run on VPS**
   ```bash
   npm run start
   ```

### Cloud Deployment Options

#### Docker Hub & Production Server
```bash
# 1. Build and push images
docker build -t your-username/go-ecommerce-backend ./go-backend
docker build -t your-username/next-ecommerce-frontend ./next-frontend
docker push your-username/go-ecommerce-backend
docker push your-username/next-ecommerce-frontend

# 2. On production server
docker-compose -f docker-compose.prod.yml pull
docker-compose -f docker-compose.prod.yml --env-file .env up -d
```

#### AWS/DigitalOcean Deployment
1. **Setup VPS with Docker**
2. **Clone repository**
3. **Configure environment variables**
4. **Run production docker-compose**
5. **Setup domain and SSL**

### Monitoring and Logs
```bash
# View application logs
docker compose logs -f

# Monitor resource usage
docker stats

# Health checks
curl http://localhost:4001/health
curl http://localhost:3001/api/health
```

## 🧪 Testing

### Backend Testing
```bash
cd go-backend

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./internal/...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -v ./internal/core/services/

# Run tests in Docker
docker compose exec go-backend go test ./...
```

### Frontend Testing
```bash
cd next-frontend

# Run tests
npm test

# Run tests with coverage
npm run test:coverage

# Run tests in watch mode
npm run test:watch

# Run E2E tests (if configured)
npm run test:e2e
```

### Integration Testing with Docker
```bash
# Start test environment
docker compose --env-file .env.development up -d

# Run API tests
curl -X GET http://localhost:4001/api/v1/health
curl -X GET http://localhost:4001/api/v1/categories

# Run frontend tests
curl -X GET http://localhost:3001/
```

## 📚 Documentation

### API Documentation
- **Swagger UI**: Available at `/swagger/` when backend is running
- **OpenAPI Spec**: Available at `/swagger/doc.json`
- **Postman Collection**: Available in `docs/` folder

### Project Documentation
- **Database Schema**: See [SEEDER_GUIDE.md](go-backend/SEEDER_GUIDE.md)
- **API Endpoints**: Detailed list in [API_ENDPOINTS.md](go-backend/API_ENDPOINTS.md)
- **Frontend Components**: Component documentation in respective folders
- **Docker Setup**: This README.md file

### Code Documentation
```bash
# Generate Go documentation
cd go-backend
godoc -http=:6060

# Generate frontend documentation
cd next-frontend
npm run docs  # If configured
```

### Database Schema Documentation
```sql
-- Core tables overview
Users (authentication & profiles)
├── Categories (product categories)
├── Products (product catalog)
├── CartItems (shopping cart)
├── Orders (order records)
├── OrderItems (order details)
└── Payments (payment transactions)
```

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/) - Express-inspired web framework for Go
- [Next.js](https://nextjs.org/) - The React framework for production
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [PostgreSQL](https://www.postgresql.org/) - Advanced open source database
- [Docker](https://www.docker.com/) - Containerization platform
- [Nginx](https://nginx.org/) - High-performance web server
- [Swagger](https://swagger.io/) - API documentation
- [Shadcn/ui](https://ui.shadcn.com/) - Beautiful and accessible React components

## 📊 Project Statistics

### Backend Stats
- **Go Modules**: 15+ dependencies
- **API Endpoints**: 25+ REST endpoints
- **Database Tables**: 8 core tables
- **Middleware**: Authentication, CORS, Rate limiting
- **Services**: 8 business service layers

### Frontend Stats
- **React Components**: 50+ reusable components
- **Pages**: 15+ application pages
- **Custom Hooks**: 5+ React hooks
- **Context Providers**: 3+ global state management
- **UI Components**: Shadcn/ui + custom components

### Docker Configuration
- **Services**: 5 containerized services
- **Environments**: Development & Production
- **Networks**: Custom bridge network
- **Volumes**: Persistent data storage
- **Health Checks**: All services monitored

## 🔮 Roadmap

### Upcoming Features
- [ ] **Multi-language Support** (i18n)
- [ ] **Email Notifications** (Order confirmations, etc.)
- [ ] **Advanced Search** (Elasticsearch integration)
- [ ] **Inventory Management** (Stock tracking)
- [ ] **Coupon System** (Discount codes)
- [ ] **Review & Rating System** (Product reviews)
- [ ] **Wishlist Sharing** (Social features)
- [ ] **Advanced Analytics** (Sales reporting)
- [ ] **Mobile App** (React Native)
- [ ] **Microservices** (Service decomposition)

### Technical Improvements
- [ ] **Redis Caching** (Performance optimization)
- [ ] **Message Queue** (Background jobs)
- [ ] **CDN Integration** (Static asset delivery)
- [ ] **Monitoring** (Prometheus & Grafana)
- [ ] **CI/CD Pipeline** (GitHub Actions)
- [ ] **Load Balancing** (Multiple backend instances)
- [ ] **SSL/TLS** (HTTPS support)
- [ ] **Unit Tests** (Comprehensive test coverage)

## 📞 Support

For support, email support@example.com or create an issue in this repository.

---

**Built with ❤️ using Go Fiber + Next.js**