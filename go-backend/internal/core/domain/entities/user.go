package entities

import (
	"time"

	"github.com/google/uuid"
)

// AuthRequest และ Response structs
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password_complex"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type AdminRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password_complex"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	RoleID    string `json:"role_id" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,password_complex"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,password_complex"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// User Entity
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Active    bool      `json:"active"`
	RoleID    uuid.UUID `json:"role_id"`
	Role      *Role     `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

// Role Entity
type Role struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Permission Entity
type Permission struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Category Entity
type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// Product Entity
type Product struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	Image       string         `json:"image"`
	Images      []ProductImage `json:"images,omitempty"`
	CategoryID  uuid.UUID      `json:"category_id"`
	Category    *Category      `json:"category,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ProductImage struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" validate:"required,min=0"`
	Stock       int       `json:"stock" validate:"min=0"`
	Image       string    `json:"image"`
	CategoryID  uuid.UUID `json:"category_id" validate:"required"`
	Images      []string  `json:"images"`
}

type UpdateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price" validate:"min=0"`
	Stock       int       `json:"stock" validate:"min=0"`
	Image       string    `json:"image"`
	CategoryID  uuid.UUID `json:"category_id"`
	Images      []string  `json:"images"`
}

type ProductSearchRequest struct {
	Query      string    `json:"query"`
	CategoryID uuid.UUID `json:"category_id"`
	MinPrice   float64   `json:"min_price"`
	MaxPrice   float64   `json:"max_price"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

// Cart Entity
type Cart struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	CartItems  []CartItem `json:"cart_items"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	Product   *Product  `json:"product,omitempty"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddToCartRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

// Order Entity
type Order struct {
	ID              uuid.UUID     `json:"id"`
	UserID          uuid.UUID     `json:"user_id"`
	User            *User         `json:"user,omitempty"`
	OrderItems      []OrderItem   `json:"order_items"`
	TotalPrice      float64       `json:"total_price"`
	Status          string        `json:"status"`
	PaymentMethod   string        `json:"payment_method"`
	PaymentStatus   string        `json:"payment_status"`
	ShippingMethod  string        `json:"shipping_method"`
	ShippingStatus  string        `json:"shipping_status"`
	ShippingAddress string        `json:"shipping_address"`
	TrackingNumber  string        `json:"tracking_number"`
	Notes           string        `json:"notes"`
	Transactions    []Transaction `json:"transactions,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Product   *Product  `json:"product,omitempty"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	PaymentMethod   string `json:"payment_method" validate:"required"`
	ShippingMethod  string `json:"shipping_method" validate:"required"`
	ShippingAddress string `json:"shipping_address" validate:"required"`
	Notes           string `json:"notes"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type UpdatePaymentStatusRequest struct {
	PaymentStatus string `json:"payment_status" validate:"required"`
}

type UpdateShippingStatusRequest struct {
	ShippingStatus string `json:"shipping_status" validate:"required"`
	TrackingNumber string `json:"tracking_number"`
}

// Transaction Entity
type Transaction struct {
	ID            uuid.UUID `json:"id"`
	OrderID       uuid.UUID `json:"order_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"`
	PaymentData   string    `json:"payment_data"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreatePaymentRequest struct {
	OrderID       uuid.UUID `json:"order_id" validate:"required"`
	PaymentMethod string    `json:"payment_method" validate:"required"`
	PaymentData   string    `json:"payment_data"`
}

type VerifyPaymentRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	PaymentData   string `json:"payment_data"`
}

// Stats Entity
type SalesStats struct {
	TotalSales    float64 `json:"total_sales"`
	TotalOrders   int     `json:"total_orders"`
	TodaySales    float64 `json:"today_sales"`
	TodayOrders   int     `json:"today_orders"`
	MonthlySales  float64 `json:"monthly_sales"`
	MonthlyOrders int     `json:"monthly_orders"`
	YearlySales   float64 `json:"yearly_sales"`
	YearlyOrders  int     `json:"yearly_orders"`
}

type ProductStats struct {
	TotalProducts      int `json:"total_products"`
	LowStockProducts   int `json:"low_stock_products"`
	OutOfStockProducts int `json:"out_of_stock_products"`
	TotalCategories    int `json:"total_categories"`
}

type UserStats struct {
	TotalUsers    int `json:"total_users"`
	ActiveUsers   int `json:"active_users"`
	InactiveUsers int `json:"inactive_users"`
	NewUsers      int `json:"new_users"`
}

// Common Response Types
type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ApiResponse struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       interface{}         `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
