package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, req *entities.CreateCategoryRequest) (*entities.Category, error) {
	categoryModel := &models.Category{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}

	if err := r.db.WithContext(ctx).Create(categoryModel).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(categoryModel), nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Category, error) {
	var categoryModel models.Category
	if err := r.db.WithContext(ctx).First(&categoryModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&categoryModel), nil
}

func (r *categoryRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Category, int, error) {
	var categories []models.Category
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Category
	for _, category := range categories {
		result = append(result, r.modelToEntity(&category))
	}

	return result, int(total), nil
}

func (r *categoryRepository) Update(ctx context.Context, id uuid.UUID, req *entities.UpdateCategoryRequest) error {
	updates := map[string]interface{}{}

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Image != "" {
		updates["image"] = req.Image
	}

	return r.db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", id).Updates(updates).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, "id = ?", id).Error
}

func (r *categoryRepository) modelToEntity(categoryModel *models.Category) *entities.Category {
	return &entities.Category{
		ID:          categoryModel.ID,
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
		Image:       categoryModel.Image,
		CreatedAt:   categoryModel.CreatedAt,
		UpdatedAt:   categoryModel.UpdatedAt,
	}
}
