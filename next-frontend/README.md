# Go Fiber Next E-commerce Frontend

โปรเจ็กต์ frontend สำหรับระบบ e-commerce ที่พัฒนาด้วย Next.js 15 และ TypeScript ทำงานร่วมกับ backend ที่พัฒนาด้วย Go Fiber รองรับการ deploy ด้วย **Docker**

## 🚀 ฟีเจอร์หลัก

- **เทคโนโลยีทันสมัย**: พัฒนาด้วย Next.js 15, React 19, TypeScript
- **UI ที่สวยงาม**: ออกแบบด้วย Tailwind CSS v4 และ Radix UI components
- **ระบบยืนยันตัวตน**: Authentication และ Authorization ด้วย JWT
- **การจัดการฟอร์ม**: React Hook Form พร้อม Zod validation
- **การตอบสนองต่อหน้าจอ**: แนวทาง Mobile-first พร้อมแอนิเมชั่นที่ทันสมัย
- **สถาปัตยกรรม Component**: โครงสร้างแบบโมดูลาร์และนำไปใช้ซ้ำได้
- **การปรับปรุงรูปภาพ**: การปรับปรุงรูปภาพอัตโนมัติของ Next.js
- **แบ่งหน้าแล้ว**: แยกโซน user และ admin
- **แอนิเมชั่น**: tw-animate-css สำหรับแอนิเมชั่นที่ลื่นไหล
- **🐳 Docker Support**: รองรับการ deploy ด้วย Docker และ Docker Compose

## 🛠️ เทคโนโลยีที่ใช้

- **เฟรมเวิร์ก**: [Next.js 15](https://nextjs.org/) พร้อม App Router
- **ภาษาโปรแกรม**: [TypeScript](https://www.typescriptlang.org/)
- **การจัดแต่งสไตล์**: [Tailwind CSS v4](https://tailwindcss.com/)
- **UI Components**: [Radix UI](https://www.radix-ui.com/)
- **ไอคอน**: [Lucide React](https://lucide.dev/)
- **แอนิเมชั่น**: [tw-animate-css](https://github.com/ben-rogerson/twin.macro)
- **สไลด์เซอร์**: [Swiper.js](https://swiperjs.com/)
- **ฟอร์ม**: [React Hook Form](https://react-hook-form.com/)
- **การตรวจสอบ**: [Zod](https://zod.dev/)
- **HTTP Client**: [Axios](https://axios-http.com/)
- **การจัดการ Cookies**: [js-cookie](https://github.com/js-cookie/js-cookie)
- **การแจ้งเตือน**: [React Hot Toast](https://react-hot-toast.com/)

## 📁 โครงสร้างโปรเจ็กต์

```
next-frontend/
├── Dockerfile                      # Production Docker configuration
├── Dockerfile.dev                 # Development Docker configuration  
├── public/                         # ไฟล์สแตติก
│   ├── images/                      # รูปภาพสินค้าและ UI
│   │   ├── arrivals/               # รูปภาพสินค้าใหม่
│   │   ├── blog/                   # รูปภาพบล็อก
│   │   ├── cart/                   # รูปภาพตะกร้าสินค้า
│   │   ├── categories/             # รูปภาพหมวดหมู่
│   │   ├── checkout/               # ไอคอนชำระเงิน
│   │   ├── products/               # รูปภาพสินค้า
│   │   ├── sellers/                # รูปภาพขายดี
│   │   └── users/                  # รูปภาพผู้ใช้
│   └── *.svg                       # ไอคอน SVG
├── src/
│   ├── app/                        # Next.js App Router
│   │   ├── (auth)/                 # กลุ่มเส้นทางสำหรับ Authentication
│   │   ├── (pages)/                # กลุ่มเส้นทางหลักของเว็บไซต์
│   │   ├── globals.css             # สไตล์ทั่วไป
│   │   └── favicon.ico             # ไอคอนไซต์
│   ├── components/                 # React components
│   │   ├── auth/                   # คอมโพเนนต์การยืนยันตัวตน
│   │   ├── blog/                   # คอมโพเนนต์บล็อก
│   │   ├── cart/                   # คอมโพเนนต์ตะกร้าสินค้า
│   │   ├── checkout/               # คอมโพเนนต์ชำระเงิน
│   │   ├── home/                   # คอมโพเนนต์หน้าแรก
│   │   ├── myaccount/              # คอมโพเนนต์บัญชีผู้ใช้
│   │   ├── orders/                 # คอมโพเนนต์คำสั่งซื้อ
│   │   ├── shop/                   # คอมโพเนนต์ร้านค้า
│   │   ├── shared/                 # คอมโพเนนต์ที่ใช้ร่วมกัน
│   │   ├── ui/                     # คอมโพเนนต์ UI ที่ใช้ซ้ำได้
│   │   └── wishlist/               # คอมโพเนนต์รายการโปรด
│   ├── context/                    # React Context providers
│   │   ├── CartSidebarModalContext.tsx
│   │   ├── PreviewSliderContext.tsx
│   │   └── QuickViewModalContext.tsx
│   ├── hooks/                      # Custom React hooks
│   │   └── useAuth.tsx             # Hook สำหรับการยืนยันตัวตน
│   ├── lib/                        # ไลบรารี่เครื่องมือ
│   │   └── utils.ts                # ฟังก์ชั่นเครื่องมือ
│   └── types/                      # นิยาม TypeScript types
│       ├── blogItem.ts             # ประเภทข้อมูลบล็อก
│       ├── category.ts             # ประเภทข้อมูลหมวดหมู่
│       ├── Menu.ts                 # ประเภทข้อมูลเมนู
│       ├── product.ts              # ประเภทข้อมูลสินค้า
│       └── testimonial.ts          # ประเภทข้อมูลคำแนะนำ
├── components.json                 # การตั้งค่า shadcn/ui
├── next.config.ts                  # การตั้งค่า Next.js
├── package.json                    # Dependencies และ scripts
├── postcss.config.mjs             # การตั้งค่า PostCSS
├── tailwind.config.js             # การตั้งค่า Tailwind CSS
└── tsconfig.json                   # การตั้งค่า TypeScript
```

## 🚀 การเริ่มต้นใช้งาน

### ข้อกำหนดเบื้องต้น

- Node.js 18+ และ npm/yarn/pnpm
- Go Fiber Backend API (หรือใช้ Docker)
- Docker & Docker Compose (แนะนำ)
- Git

### Option 1: Docker Setup (แนะนำ)

#### 1. Clone และ Setup Environment
```bash
# Clone repository
git clone <repository-url>
cd go-fiber-next-ecommerce

# Setup environment variables
cp .env.example .env.development  # For development
cp .env.example .env              # For production
```

#### 2. Docker Development Setup
```bash
# Start all services (frontend, backend, database)
docker compose --env-file .env.development up --build -d

# Start only frontend + backend + database
docker compose --env-file .env.development up --build next-frontend go-backend postgres -d

# View frontend logs
docker compose --env-file .env.development logs -f next-frontend

# Access applications:
# Frontend: http://localhost:3001
# Backend: http://localhost:4001
```

#### 3. Docker Production Setup
```bash
# Start production services
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# Access applications:
# Nginx Proxy: http://localhost:8800
# Frontend: http://localhost:3001
# Backend: http://localhost:4001
```

#### 4. Stop Services
```bash
# Development
docker compose --env-file .env.development down --volumes

# Production
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Option 2: Manual Setup

#### 1. การติดตั้ง

**เข้าไปยังโฟลเดอร์ frontend**
```bash
cd next-frontend
```

2. **ติดตั้ง dependencies**
```bash
npm install
# หรือ
yarn install
# หรือ
pnpm install
```

3. **ตั้งค่าตัวแปร environment**
```bash
# สร้างไฟล์ .env.local
echo "NEXT_PUBLIC_API_URL=http://localhost:4001" > .env.local
```

4. **รันเซิร์ฟเวอร์สำหรับพัฒนา**
```bash
npm run dev
# หรือ
yarn dev
# หรือ
pnpm dev
```

5. **เปิดเบราว์เซอร์**
ไปที่ [http://localhost:3001](http://localhost:3001) เพื่อดูแอปพลิเคชั่น

## 🐳 Docker Commands

### Development Environment
```bash
# Start frontend with all services (from root directory)
docker compose --env-file .env.development up --build -d

# Start only frontend service  
docker compose --env-file .env.development up --build next-frontend -d

# View frontend logs
docker compose --env-file .env.development logs -f next-frontend

# Execute commands in frontend container
docker compose --env-file .env.development exec next-frontend sh

# Restart frontend service
docker compose --env-file .env.development restart next-frontend

# Stop all services
docker compose --env-file .env.development down --volumes
```

### Production Environment
```bash
# Start production frontend
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# View production logs
docker compose -f docker-compose.prod.yml logs -f next-frontend

# Stop production services
docker compose -f docker-compose.prod.yml --env-file .env down --volumes
```

### Frontend-only Docker Commands
```bash
# Build development image
docker build -f Dockerfile.dev -t next-ecommerce-frontend-dev .

# Build production image
docker build -f Dockerfile -t next-ecommerce-frontend-prod .

# Run development container standalone
docker run -p 3001:3001 --env-file .env.local next-ecommerce-frontend-dev

# Run production container standalone
docker run -p 3001:3001 -e NEXT_PUBLIC_API_URL=http://localhost:4001 next-ecommerce-frontend-prod
```

### Docker Development Features
- **Hot Reload**: โค้ดจะ reload อัตโนมัติเมื่อมีการแก้ไข
- **Volume Mounting**: โค้ดใน host จะ sync กับ container  
- **Fast Refresh**: Next.js fast refresh ทำงานใน Docker
- **Environment Variables**: รองรับ environment-specific configuration

## 📝 คำสั่งที่มีให้ใช้

### Manual Commands
- `npm run dev` - เริ่มเซิร์ฟเวอร์พัฒนาด้วย Turbopack บนพอร์ต 3001
- `npm run build` - สร้างบิลด์สำหรับ production ที่ปรับปรุงแล้ว
- `npm run start` - เริ่มเซิร์ฟเวอร์ production บนพอร์ต 3001
- `npm run lint` - รัน ESLint เพื่อตรวจสอบคุณภาพโค้ด

### Docker Commands
- `docker compose --env-file .env.development up next-frontend -d` - เริ่ม development container
- `docker compose -f docker-compose.prod.yml up next-frontend -d` - เริ่ม production container
- `docker compose logs -f next-frontend` - ดู logs
- `docker compose exec next-frontend npm run build` - Build ใน container

## 🏗️ สถาปัตยกรรม

### โครงสร้าง Component

- **Auth Components** (`src/components/auth/`): คอมโพเนนต์การยืนยันตัวตน
- **E-commerce Components**: คอมโพเนนต์สำหรับ shop, cart, checkout, orders
- **Blog Components**: คอมโพเนนต์สำหรับระบบบล็อก
- **Shared Components** (`src/components/shared/`): คอมโพเนนต์ที่ใช้ร่วมกันทั่วไซต์
- **UI Components** (`src/components/ui/`): คอมโพเนนต์ UI พื้นฐาน

### ระบบ Layout

โปรเจ็กต์ใช้ Next.js App Router พร้อมโครงสร้างเลย์เอาต์แบบจัดกลุ่ม:
- `(auth)/` - เส้นทางสำหรับการยืนยันตัวตน (signin, signup)
- `(pages)/` - เส้นทางหลักของเว็บไซต์ (home, shop, blog, etc.)

### ระบบ Authentication

- ใช้ JWT tokens เก็บใน cookies
- Custom `useAuth` hook สำหรับจัดการสถานะ authentication
- Form validation ด้วย React Hook Form + Zod
- Auto redirect ตามสิทธิ์ผู้ใช้

### Context Management

- `CartSidebarModalContext` - จัดการ modal ตะกร้าสินค้า
- `PreviewSliderContext` - จัดการ slider preview
- `QuickViewModalContext` - จัดการ modal quick view สินค้า

## 🎨 ฟีเจอร์ UI/UX

- **ดีไซน์ทันสมัย**: อินเตอร์เฟซที่สะอาดและเรียบง่าย
- **เลย์เอาต์ตอบสนอง**: ทำงานได้กับทุกขนาดอุปกรณ์
- **องค์ประกอบแบบโต้ตอบ**: แอนิเมชั่นและทรานซิชั่นที่ลื่นไหล
- **การแจ้งเตือน**: React Hot Toast สำหรับ notifications
- **Slider**: Swiper.js สำหรับแกลเลอรี่และสไลด์
- **Range Slider**: สำหรับการกรองราคาสินค้า

## 🔧 การตั้งค่า

### Environment Variables

#### Development (.env.local)
```bash
NEXT_PUBLIC_API_URL=http://localhost:4001  # URL ของ Go Fiber backend

# For Docker development
NEXT_PUBLIC_API_URL=http://go-backend:4001  # ใช้ service name ใน Docker network
```

#### Production (.env)  
```bash
NEXT_PUBLIC_API_URL=https://your-api-domain.com  # Production API URL
NODE_ENV=production
```

#### Docker Environment Variables
```bash
# In docker-compose.yml
NODE_ENV=development
NEXT_PUBLIC_API_URL=http://go-backend:4001
PORT=3001
HOSTNAME=0.0.0.0

# For hot reload in Docker
WATCHPACK_POLLING=true
CHOKIDAR_USEPOLLING=true
```

### Tailwind CSS
โปรเจ็กต์ใช้ Tailwind CSS v4 พร้อมการตั้งค่าที่กำหนดเองใน `tailwind.config.js`

### TypeScript
รองรับ TypeScript แบบเต็มรูปแบบพร้อมการตรวจสอบประเภทแบบเข้มงวด

### Next.js
- App Router สำหรับ routing
- การปรับปรุงรูปภาพและฟอนต์
- Turbopack สำหรับ fast refresh

## 📱 การออกแบบตอบสนอง

แอปพลิเคชั่นสร้างด้วยแนวทาง mobile-first:
- **มือถือ**: 320px - 768px
- **แท็บเล็ต**: 768px - 1024px
- **เดสก์ท็อป**: 1024px+

## 🔗 การเชื่อมต่อ API

แอปพลิเคชั่นเชื่อมต่อกับ Go Fiber backend ผ่าน:
- RESTful API endpoints
- JWT authentication
- Axios HTTP client  
- Auto token refresh

### API Configuration
```typescript
// lib/api.ts
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:4001';

// In Docker environment
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://go-backend:4001';
```

## 🚀 การ Deploy

### Option 1: Docker Deployment (แนะนำ)

#### Development Deployment
```bash
# Start development environment
docker compose --env-file .env.development up --build -d

# Access frontend: http://localhost:3001
```

#### Production Deployment  
```bash
# Start production environment
docker compose -f docker-compose.prod.yml --env-file .env up --build -d

# Access via Nginx proxy: http://localhost:8800
# Direct access: http://localhost:3001
```

#### Standalone Docker Deployment
```bash
# Build production image
docker build -f Dockerfile -t next-ecommerce-frontend .

# Run container
docker run -d -p 3001:3001 \
  -e NEXT_PUBLIC_API_URL=https://your-api-domain.com \
  next-ecommerce-frontend
```

### Option 2: Manual Deployment

1. **สร้าง production build**
   ```bash
   npm run build
   ```

2. **เริ่ม production server**
   ```bash
   npm run start
   ```

3. **หรือ deploy ไปยัง Vercel**
   ```bash
   vercel --prod
   ```

### Option 3: Cloud Deployment

#### Vercel Deployment
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy to Vercel
vercel --prod

# Set environment variables in Vercel dashboard
NEXT_PUBLIC_API_URL=https://your-api-domain.com
```

#### VPS/Server Deployment
```bash
# 1. Clone on server
git clone <repository-url>
cd go-fiber-next-ecommerce

# 2. Set up environment
cp .env.example .env
# Edit .env with production values

# 3. Deploy with Docker
docker compose -f docker-compose.prod.yml --env-file .env up -d

# 4. Access via Nginx proxy
curl http://your-server-ip:8800
```

### Environment-specific Deployment

#### Development
- Hot reload enabled
- Source maps included
- Debug mode on
- Volume mounting for live code changes

#### Production  
- Optimized build
- Static file serving
- Compressed assets
- Health checks enabled

## 🧪 Testing

### Manual Testing
```bash
# Run tests (if configured)
npm test

# Run tests with coverage
npm run test:coverage

# Run E2E tests
npm run test:e2e
```

### Docker Testing
```bash
# Run tests in Docker container
docker compose --env-file .env.development exec next-frontend npm test

# Build test in Docker
docker compose --env-file .env.development exec next-frontend npm run build
```

### Frontend Testing
```bash
# Test frontend accessibility
curl http://localhost:3001

# Test API connectivity (requires backend)
curl http://localhost:3001/api/health
```

## �️ Docker Configuration

### Dockerfile.dev (Development)
- Hot reload support
- Volume mounting for live code changes
- Development dependencies included
- Source maps enabled

### Dockerfile (Production)
- Multi-stage build for optimization
- Static file serving with Next.js
- Minimal image size
- Health checks included

### Docker Compose Integration
- Network isolation with backend services
- Environment variable injection
- Volume persistence for development
- Service dependencies management

## 🔧 Troubleshooting

### Common Docker Issues

#### Port Already in Use
```bash
# Check what's using port 3001
netstat -tulpn | grep :3001

# Kill process
sudo kill -9 $(sudo lsof -t -i:3001)
```

#### Container Build Failures
```bash
# Clean Docker cache
docker builder prune -a

# Rebuild without cache
docker compose build --no-cache next-frontend

# Remove container and rebuild
docker compose down next-frontend
docker compose up --build next-frontend -d
```

#### API Connection Issues
```bash
# Check API connectivity from container
docker compose exec next-frontend curl http://go-backend:4001/health

# Check environment variables
docker compose exec next-frontend env | grep NEXT_PUBLIC
```

## 🤝 การมีส่วนร่วม

1. Fork โปรเจ็กต์
2. สร้างบรานช์ฟีเจอร์ของคุณ (`git checkout -b feature/amazing-feature`)
3. Commit การเปลี่ยนแปลงของคุณ (`git commit -m 'Add some amazing feature'`)
4. Push ไปยังบรานช์ (`git push origin feature/amazing-feature`)
5. เปิด Pull Request

## 📄 ใบอนุญาต

โปรเจ็กต์นี้ได้รับใบอนุญาตภายใต้ MIT License

## 🙏 การขอบคุณ

- [Next.js](https://nextjs.org/) สำหรับ React framework ที่ยอดเยี่ยม
- [Tailwind CSS](https://tailwindcss.com/) สำหรับ utility-first CSS framework
- [Radix UI](https://www.radix-ui.com/) สำหรับ UI components ที่เข้าถึงได้
- [Go Fiber](https://gofiber.io/) สำหรับ backend API framework

---

สร้างด้วย ❤️ โดยใช้ Next.js, TypeScript และ Go Fiber
