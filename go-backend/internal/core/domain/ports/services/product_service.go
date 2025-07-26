package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// ProductService interface สำหรับการจัดการสินค้า
type ProductService interface {
	CreateProduct(ctx context.Context, req *entities.CreateProductRequest) (*entities.Product, error)
	GetProducts(ctx context.Context, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID uuid.UUID, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error)
	SearchProducts(ctx context.Context, req *entities.ProductSearchRequest) ([]*entities.Product, *entities.PaginationResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *entities.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
