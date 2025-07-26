# E-commerce API Endpoints

## Authentication

### Public Routes
- `POST /api/v1/auth/register` - ลงทะเบียนผู้ใช้
- `POST /api/v1/auth/admin/register` - ลงทะเบียน Admin
- `POST /api/v1/auth/login` - เข้าสู่ระบบ
- `POST /api/v1/auth/refresh` - รีเฟรช token
- `POST /api/v1/auth/forgot-password` - ลืมรหัสผ่าน
- `POST /api/v1/auth/reset-password` - รีเซ็ตรหัสผ่าน

### Protected Routes
- `POST /api/v1/auth/logout` - ออกจากระบบ
- `POST /api/v1/auth/change-password` - เปลี่ยนรหัสผ่าน

## Users (Admin Only)
- `GET /api/v1/users` - ดูรายการผู้ใช้ทั้งหมด
- `GET /api/v1/users/:id` - ดูข้อมูลผู้ใช้
- `PUT /api/v1/users/:id` - แก้ไขข้อมูลผู้ใช้
- `DELETE /api/v1/users/:id` - ลบผู้ใช้

## Categories
### Public Routes
- `GET /api/v1/categories` - ดูรายการหมวดหมู่
- `GET /api/v1/categories/:id` - ดูข้อมูลหมวดหมู่

### Admin Only
- `POST /api/v1/categories` - สร้างหมวดหมู่
- `PUT /api/v1/categories/:id` - แก้ไขหมวดหมู่
- `DELETE /api/v1/categories/:id` - ลบหมวดหมู่

## Products
### Public Routes
- `GET /api/v1/products` - ดูรายการสินค้า
- `GET /api/v1/products/:id` - ดูข้อมูลสินค้า
- `GET /api/v1/products/category/:categoryId` - ดูสินค้าตามหมวดหมู่
- `GET /api/v1/products/search` - ค้นหาสินค้า

### Admin Only
- `POST /api/v1/products` - สร้างสินค้า
- `PUT /api/v1/products/:id` - แก้ไขสินค้า
- `DELETE /api/v1/products/:id` - ลบสินค้า

## Cart (Authenticated Users)
- `GET /api/v1/cart` - ดูตะกร้าสินค้า
- `POST /api/v1/cart` - เพิ่มสินค้าลงตะกร้า
- `PUT /api/v1/cart/:itemId` - แก้ไขจำนวนสินค้าในตะกร้า
- `DELETE /api/v1/cart/:itemId` - ลบสินค้าจากตะกร้า
- `DELETE /api/v1/cart` - ล้างตะกร้าสินค้า

## Orders
### User Routes
- `POST /api/v1/orders` - สร้างคำสั่งซื้อ
- `GET /api/v1/orders` - ดูคำสั่งซื้อของตนเอง
- `GET /api/v1/orders/:id` - ดูรายละเอียดคำสั่งซื้อ
- `PUT /api/v1/orders/:id/cancel` - ยกเลิกคำสั่งซื้อ

### Admin Routes
- `GET /api/v1/orders/admin` - ดูคำสั่งซื้อทั้งหมด
- `PUT /api/v1/orders/admin/:id/status` - อัพเดทสถานะคำสั่งซื้อ

## Payments (Authenticated Users)
- `POST /api/v1/payments` - สร้างการชำระเงิน
- `POST /api/v1/payments/:id/verify` - ยืนยันการชำระเงิน
- `PUT /api/v1/payments/:id/cancel` - ยกเลิกการชำระเงิน

## Statistics (Admin Only)
- `GET /api/v1/stats/sales` - สถิติการขาย
- `GET /api/v1/stats/products` - สถิติสินค้า
- `GET /api/v1/stats/users` - สถิติผู้ใช้

## Other Routes
- `GET /health` - ตรวจสอบสถานะ API
- `GET /swagger/*` - API Documentation

## Query Parameters

### Pagination
- `page` - หน้าที่ต้องการ (default: 1)
- `limit` - จำนวนรายการต่อหน้า (default: 10, max: 100)

### Product Search
- `search` - คำค้นหา
- `category_id` - ID หมวดหมู่
- `min_price` - ราคาต่ำสุด
- `max_price` - ราคาสูงสุด

### Statistics
- `period` - ช่วงเวลา (daily, weekly, monthly, yearly)

## Response Format

### Success Response
```json
{
  "success": true,
  "message": "ข้อความสำเร็จ",
  "data": {},
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "pages": 10
  }
}
```

### Error Response
```json
{
  "success": false,
  "message": "ข้อความแสดงข้อผิดพลาด"
}
```

## Authentication Header
สำหรับ endpoints ที่ต้องการการยืนยันตัวตน ให้ส่ง header:
```
Authorization: Bearer <JWT_TOKEN>
``` 