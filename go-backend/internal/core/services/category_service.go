package services

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) services.CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req *entities.CreateCategoryRequest) (*entities.Category, error) {
	return s.categoryRepo.Create(ctx, req)
}

func (s *categoryService) GetCategories(ctx context.Context, page, limit int) ([]*entities.Category, *entities.PaginationResponse, error) {
	categories, total, err := s.categoryRepo.GetAll(ctx, page, limit)
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

	return categories, pagination, nil
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*entities.Category, error) {
	return s.categoryRepo.GetByID(ctx, id)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id uuid.UUID, req *entities.UpdateCategoryRequest) error {
	return s.categoryRepo.Update(ctx, id, req)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.categoryRepo.Delete(ctx, id)
}
