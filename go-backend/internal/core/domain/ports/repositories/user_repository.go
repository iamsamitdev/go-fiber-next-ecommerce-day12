package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// UserRepository interface สำหรับการจัดการข้อมูลผู้ใช้
type UserRepository interface {
	Create(ctx context.Context, user *entities.User, password string) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.User, int, error)
	Update(ctx context.Context, id uuid.UUID, user *entities.UpdateUserRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error
	SetRefreshToken(ctx context.Context, id uuid.UUID, token string) error
	GetByRefreshToken(ctx context.Context, token string) (*entities.User, error)
	SetResetToken(ctx context.Context, email string, token string) error
	GetByResetToken(ctx context.Context, token string) (*entities.User, error)
	ClearResetToken(ctx context.Context, id uuid.UUID) error
	GetPasswordHash(ctx context.Context, id uuid.UUID) (string, error)
}

// RoleRepository interface สำหรับการจัดการบทบาท
type RoleRepository interface {
	Create(ctx context.Context, role *entities.Role) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Role, error)
	GetByName(ctx context.Context, name string) (*entities.Role, error)
	GetAll(ctx context.Context) ([]*entities.Role, error)
	Update(ctx context.Context, id uuid.UUID, role *entities.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// PermissionRepository interface สำหรับการจัดการสิทธิ์
type PermissionRepository interface {
	Create(ctx context.Context, permission *entities.Permission) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Permission, error)
	GetByName(ctx context.Context, name string) (*entities.Permission, error)
	GetAll(ctx context.Context) ([]*entities.Permission, error)
	Update(ctx context.Context, id uuid.UUID, permission *entities.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// CategoryRepository interface สำหรับการจัดการหมวดหมู่
type CategoryRepository interface {
	Create(ctx context.Context, category *entities.CreateCategoryRequest) (*entities.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Category, int, error)
	Update(ctx context.Context, id uuid.UUID, category *entities.UpdateCategoryRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// ProductRepository interface สำหรับการจัดการสินค้า
type ProductRepository interface {
	Create(ctx context.Context, product *entities.CreateProductRequest) (*entities.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Product, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Product, int, error)
	GetByCategory(ctx context.Context, categoryID uuid.UUID, page, limit int) ([]*entities.Product, int, error)
	Search(ctx context.Context, req *entities.ProductSearchRequest) ([]*entities.Product, int, error)
	Update(ctx context.Context, id uuid.UUID, product *entities.UpdateProductRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateStock(ctx context.Context, id uuid.UUID, stock int) error
	GetLowStockProducts(ctx context.Context, threshold int) ([]*entities.Product, error)
}

// CartRepository interface สำหรับการจัดการตะกร้าสินค้า
type CartRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entities.Cart, error)
	AddItem(ctx context.Context, userID uuid.UUID, item *entities.AddToCartRequest) error
	UpdateItem(ctx context.Context, cartItemID uuid.UUID, quantity int) error
	RemoveItem(ctx context.Context, cartItemID uuid.UUID) error
	ClearCart(ctx context.Context, userID uuid.UUID) error
	GetCartItem(ctx context.Context, cartItemID uuid.UUID) (*entities.CartItem, error)
}

// OrderRepository interface สำหรับการจัดการคำสั่งซื้อ
type OrderRepository interface {
	Create(ctx context.Context, userID uuid.UUID, order *entities.CreateOrderRequest) (*entities.Order, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Order, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entities.Order, int, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Order, int, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, paymentStatus string) error
	UpdateShippingStatus(ctx context.Context, id uuid.UUID, shippingStatus, trackingNumber string) error
	Cancel(ctx context.Context, id uuid.UUID) error
}

// TransactionRepository interface สำหรับการจัดการธุรกรรม
type TransactionRepository interface {
	Create(ctx context.Context, transaction *entities.CreatePaymentRequest) (*entities.Transaction, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error)
	GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]*entities.Transaction, error)
	GetByTransactionID(ctx context.Context, transactionID string) (*entities.Transaction, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	Cancel(ctx context.Context, id uuid.UUID) error
}

// StatsRepository interface สำหรับสถิติ
type StatsRepository interface {
	GetSalesStats(ctx context.Context) (*entities.SalesStats, error)
	GetProductStats(ctx context.Context) (*entities.ProductStats, error)
	GetUserStats(ctx context.Context) (*entities.UserStats, error)
}
