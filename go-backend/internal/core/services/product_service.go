package services

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) services.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *entities.CreateProductRequest) (*entities.Product, error) {
	return s.productRepo.Create(ctx, req)
}

func (s *productService) GetProducts(ctx context.Context, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error) {
	products, total, err := s.productRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	pagination := &entities.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return products, pagination, nil
}

func (s *productService) GetProductByID(ctx context.Context, id uuid.UUID) (*entities.Product, error) {
	return s.productRepo.GetByID(ctx, id)
}

func (s *productService) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID, page, limit int) ([]*entities.Product, *entities.PaginationResponse, error) {
	products, total, err := s.productRepo.GetByCategory(ctx, categoryID, page, limit)
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	pagination := &entities.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return products, pagination, nil
}

func (s *productService) SearchProducts(ctx context.Context, req *entities.ProductSearchRequest) ([]*entities.Product, *entities.PaginationResponse, error) {
	products, total, err := s.productRepo.Search(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	page := req.Page
	if page == 0 {
		page = 1
	}
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	pagination := &entities.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return products, pagination, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id uuid.UUID, req *entities.UpdateProductRequest) error {
	return s.productRepo.Update(ctx, id, req)
}

func (s *productService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.productRepo.Delete(ctx, id)
}
