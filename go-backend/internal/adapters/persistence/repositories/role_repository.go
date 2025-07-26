package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repositories.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(ctx context.Context, role *entities.Role) error {
	roleModel := &models.Role{
		Name:        role.Name,
		Description: role.Description,
	}

	return r.db.WithContext(ctx).Create(roleModel).Error
}

func (r *roleRepository) GetByName(ctx context.Context, name string) (*entities.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&role), nil
}

func (r *roleRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&role), nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]*entities.Role, error) {
	var roles []models.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}

	var result []*entities.Role
	for _, role := range roles {
		result = append(result, r.modelToEntity(&role))
	}

	return result, nil
}

func (r *roleRepository) Update(ctx context.Context, id uuid.UUID, role *entities.Role) error {
	updates := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
	}
	return r.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Updates(updates).Error
}

func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Role{}, "id = ?", id).Error
}

func (r *roleRepository) modelToEntity(role *models.Role) *entities.Role {
	return &entities.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}
