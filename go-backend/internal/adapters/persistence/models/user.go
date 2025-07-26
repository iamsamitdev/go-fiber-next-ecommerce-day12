package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel เป็นโครงสร้างพื้นฐานสำหรับทุกโมเดล
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Role สำหรับเก็บข้อมูลสิทธิ์การใช้งาน
type Role struct {
	BaseModel
	Name        string       `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string       `gorm:"type:text" json:"description"`
	Users       []User       `gorm:"foreignKey:RoleID" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// Permission สำหรับเก็บสิทธิ์การเข้าถึงระบบ
type Permission struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string `gorm:"type:text" json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

// User สำหรับเก็บข้อมูลผู้ใช้งาน
type User struct {
	BaseModel
	Email            string    `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	Password         string    `gorm:"type:varchar(100)" json:"-" validate:"required,password_complex"`
	FirstName        string    `gorm:"type:varchar(100)" json:"first_name" validate:"required"`
	LastName         string    `gorm:"type:varchar(100)" json:"last_name" validate:"required"`
	Avatar           string    `gorm:"type:varchar(255)" json:"avatar"`
	Phone            string    `gorm:"type:varchar(20)" json:"phone"`
	Address          string    `gorm:"type:text" json:"address"`
	Active           bool      `gorm:"default:true" json:"active"`
	RoleID           uuid.UUID `json:"role_id" validate:"required"`
	Role             Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Orders           []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	WishList         []Product `gorm:"many2many:user_wishlist;" json:"wishlist,omitempty"`
	RefreshToken     string    `gorm:"type:text" json:"-"`
	ResetToken       string    `gorm:"type:text" json:"-"`
	ResetTokenExpiry time.Time `json:"-"`
}

// Category สำหรับเก็บข้อมูลหมวดหมู่สินค้า
type Category struct {
	BaseModel
	Name        string    `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

// Product สำหรับเก็บข้อมูลสินค้า
type Product struct {
	BaseModel
	Name        string         `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price" validate:"required,min=0"`
	Stock       int            `gorm:"type:int" json:"stock" validate:"min=0"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	CategoryID  uuid.UUID      `json:"category_id" validate:"required"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	OrderItems  []OrderItem    `gorm:"foreignKey:ProductID" json:"order_items,omitempty"`
	CartItems   []CartItem     `gorm:"foreignKey:ProductID" json:"cart_items,omitempty"`
}

// ProductImage สำหรับเก็บรูปภาพของสินค้า
type ProductImage struct {
	BaseModel
	ProductID uuid.UUID `json:"product_id"`
	ImageURL  string    `gorm:"type:varchar(255)" json:"image_url" validate:"required"`
}

// Cart สำหรับเก็บข้อมูลตะกร้าสินค้า
type Cart struct {
	BaseModel
	UserID     uuid.UUID  `json:"user_id"`
	User       User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CartItems  []CartItem `gorm:"foreignKey:CartID" json:"cart_items,omitempty"`
	TotalPrice float64    `gorm:"type:decimal(10,2)" json:"total_price"`
}

// CartItem สำหรับเก็บรายการสินค้าในตะกร้า
type CartItem struct {
	BaseModel
	CartID    uuid.UUID `json:"cart_id"`
	Cart      Cart      `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	ProductID uuid.UUID `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int" json:"quantity" validate:"required,min=1"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
}

// Order สำหรับเก็บข้อมูลการสั่งซื้อ
type Order struct {
	BaseModel
	UserID          uuid.UUID     `json:"user_id"`
	User            User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems      []OrderItem   `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	TotalPrice      float64       `gorm:"type:decimal(10,2)" json:"total_price"`
	Status          string        `gorm:"type:varchar(50);default:'pending'" json:"status"`
	PaymentMethod   string        `gorm:"type:varchar(50)" json:"payment_method"`
	PaymentStatus   string        `gorm:"type:varchar(50);default:'pending'" json:"payment_status"`
	ShippingMethod  string        `gorm:"type:varchar(50)" json:"shipping_method"`
	ShippingStatus  string        `gorm:"type:varchar(50);default:'pending'" json:"shipping_status"`
	ShippingAddress string        `gorm:"type:text" json:"shipping_address"`
	TrackingNumber  string        `gorm:"type:varchar(100)" json:"tracking_number"`
	Notes           string        `gorm:"type:text" json:"notes"`
	Transactions    []Transaction `gorm:"foreignKey:OrderID" json:"transactions,omitempty"`
}

// OrderItem สำหรับเก็บรายการสินค้าในคำสั่งซื้อ
type OrderItem struct {
	BaseModel
	OrderID   uuid.UUID `json:"order_id"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	ProductID uuid.UUID `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int" json:"quantity" validate:"required,min=1"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
}

// Transaction สำหรับเก็บข้อมูลธุรกรรมการชำระเงิน
type Transaction struct {
	BaseModel
	OrderID       uuid.UUID `json:"order_id"`
	Order         Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Amount        float64   `gorm:"type:decimal(10,2)" json:"amount"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	Status        string    `gorm:"type:varchar(50);default:'pending'" json:"status"`
	TransactionID string    `gorm:"type:varchar(100)" json:"transaction_id"`
	PaymentData   string    `gorm:"type:text" json:"payment_data"`
}
