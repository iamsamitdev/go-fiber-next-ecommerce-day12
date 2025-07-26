# Go Fiber Next.js E-commerce Platform

Full-stack e-commerce platform ที่ทันสมัย พัฒนาด้วย **Go Fiber** สำหรับ backend API และ **Next.js 15** สำหรับ frontend โปรเจ็กต์นี้ให้บริการระบบ e-commerce ที่สมบูรณ์แบบพร้อมกับ authentication, authorization, และฟีเจอร์ทางการค้าครบครัน

## 🏗️ สถาปัตยกรรมระบบ

```
go-fiber-next-ecommerce/
├── go-backend/                      # Go Fiber API Backend
│   ├── cmd/                         # Entry points
│   │   ├── api/main.go             # Main API server
│   │   └── migrate/main.go         # Database migration
│   ├── docs/                       # Swagger documentation
│   ├── internal/                   # Internal application code
│   │   ├── adapters/               # External adapters
│   │   │   ├── http/               # HTTP handlers, routes, middleware
│   │   │   └── persistence/        # Database models & repositories
│   │   ├── config/                 # Configuration files
│   │   └── core/                   # Business logic
│   │       ├── domain/             # Domain entities & ports
│   │       └── services/           # Business services
│   └── pkg/                        # Shared utilities
│       └── utils/                  # JWT, password, validation utilities
└── next-frontend/                  # Next.js Frontend
    ├── public/                     # Static assets & images
    ├── src/
    │   ├── app/                    # Next.js App Router
    │   │   ├── (auth)/            # Authentication routes
    │   │   └── (pages)/           # Main application routes
    │   ├── components/            # React components
    │   │   ├── auth/              # Authentication components
    │   │   ├── shop/              # E-commerce components
    │   │   ├── shared/            # Shared components
    │   │   └── ui/                # Reusable UI components
    │   ├── context/               # React Context providers
    │   ├── hooks/                 # Custom React hooks
    │   ├── lib/                   # Utility libraries
    │   └── types/                 # TypeScript type definitions
    └── ...config files
```

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

### Frontend (Next.js)
- **🎨 Modern UI/UX**: Responsive design with Tailwind CSS v4
- **🔐 Authentication Flow**: Login, register, password reset
- **🛍️ Product Catalog**: Product browsing, search, filtering
- **🛒 Shopping Cart**: Interactive cart with real-time updates
- **💳 Checkout Process**: Complete order placement flow
- **👤 User Dashboard**: Profile management, order history
- **📱 Mobile-First**: Responsive design for all devices
- **⚡ Performance**: Optimized images, fonts, and loading
- **🎭 Animations**: Smooth transitions and micro-interactions
- **🔔 Notifications**: Real-time feedback with toast messages

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

### Frontend
- **Framework**: [Next.js 15](https://nextjs.org/) with App Router
- **Language**: [TypeScript](https://www.typescriptlang.org/)
- **Styling**: [Tailwind CSS v4](https://tailwindcss.com/)
- **UI Components**: [Radix UI](https://www.radix-ui.com/)
- **Forms**: [React Hook Form](https://react-hook-form.com/) + [Zod](https://zod.dev/)
- **HTTP Client**: [Axios](https://axios-http.com/)
- **State Management**: React Context + Custom Hooks
- **Icons**: [Lucide React](https://lucide.dev/)
- **Animations**: [tw-animate-css](https://github.com/ben-rogerson/twin.macro)
- **Sliders**: [Swiper.js](https://swiperjs.com/)

## 🚀 Quick Start

### Prerequisites
- **Go** 1.24+ 
- **Node.js** 18+
- **PostgreSQL** 13+
- **Git**

### 1. Clone Repository
```bash
git clone <repository-url>
cd go-fiber-next-ecommerce
```

### 2. Backend Setup
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

### 3. Frontend Setup
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

### Backend Environment Variables
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

### Frontend Environment Variables
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 📝 Available Scripts

### Backend
```bash
go run cmd/api/main.go           # Start API server
go run cmd/migrate/main.go       # Run migrations
go run cmd/migrate/main.go -seed # Run migrations with seeding
go test ./...                    # Run tests
go build -o bin/api cmd/api/main.go  # Build for production
```

### Frontend
```bash
npm run dev      # Start development server (port 3001)
npm run build    # Build for production
npm run start    # Start production server
npm run lint     # Run ESLint
```

## 🚀 Deployment

### Backend Deployment
1. **Build the application**
   ```bash
   go build -o bin/api cmd/api/main.go
   ```

2. **Run migrations on production**
   ```bash
   ./bin/migrate
   ```

3. **Start the API server**
   ```bash
   ./bin/api
   ```

### Frontend Deployment
1. **Build the application**
   ```bash
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

## 🧪 Testing

### Backend Testing
```bash
cd go-backend
go test ./...                    # Run all tests
go test -v ./internal/...        # Verbose testing
go test -cover ./...             # Test coverage
```

### Frontend Testing
```bash
cd next-frontend
npm test                         # Run tests
npm run test:coverage           # Test coverage
```

## 📚 Documentation

- **API Documentation**: Available at `/swagger/` when backend is running
- **Database Schema**: See [SEEDER_GUIDE.md](go-backend/SEEDER_GUIDE.md)
- **Frontend Components**: Component documentation in respective folders
- **API Endpoints**: Detailed list in [API_ENDPOINTS.md](go-backend/API_ENDPOINTS.md)

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

## 📞 Support

For support, email support@example.com or create an issue in this repository.

---

**Built with ❤️ using Go Fiber + Next.js**