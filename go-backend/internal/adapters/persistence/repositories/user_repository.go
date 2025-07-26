package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepositoryImpl เป็นการใช้งาน GORM สำหรับจัดการผู้ใช้
func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entities.User, password string) error {
	userModel := &models.User{
		Email:     user.Email,
		Password:  password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Phone:     user.Phone,
		Address:   user.Address,
		Active:    true,
		RoleID:    user.RoleID,
	}

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}

	user.ID = userModel.ID
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).Preload("Role").First(&userModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).Preload("Role").First(&userModel, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

func (r *userRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.User, int, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("Role").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.User
	for _, user := range users {
		result = append(result, r.modelToEntity(&user))
	}

	return result, int(total), nil
}

func (r *userRepository) Update(ctx context.Context, id uuid.UUID, req *entities.UpdateUserRequest) error {
	updates := map[string]interface{}{}

	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}

	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepository) UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

func (r *userRepository) SetRefreshToken(ctx context.Context, id uuid.UUID, token string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("refresh_token", token).Error
}

func (r *userRepository) GetByRefreshToken(ctx context.Context, token string) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).Preload("Role").First(&userModel, "refresh_token = ?", token).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

func (r *userRepository) SetResetToken(ctx context.Context, email string, token string) error {
	expiry := time.Now().Add(24 * time.Hour) // Token หมดอายุใน 24 ชั่วโมง
	return r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"reset_token":        token,
		"reset_token_expiry": expiry,
	}).Error
}

func (r *userRepository) GetByResetToken(ctx context.Context, token string) (*entities.User, error) {
	var userModel models.User
	if err := r.db.WithContext(ctx).Preload("Role").Where("reset_token = ? AND reset_token_expiry > ?", token, time.Now()).First(&userModel).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&userModel), nil
}

func (r *userRepository) ClearResetToken(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"reset_token":        "",
		"reset_token_expiry": nil,
	}).Error
}

func (r *userRepository) GetPasswordHash(ctx context.Context, id uuid.UUID) (string, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Select("password").First(&user, "id = ?", id).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}

func (r *userRepository) modelToEntity(userModel *models.User) *entities.User {
	user := &entities.User{
		ID:        userModel.ID,
		Email:     userModel.Email,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Avatar:    userModel.Avatar,
		Phone:     userModel.Phone,
		Address:   userModel.Address,
		Active:    userModel.Active,
		RoleID:    userModel.RoleID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	if userModel.Role.ID != uuid.Nil {
		user.Role = &entities.Role{
			ID:          userModel.Role.ID,
			Name:        userModel.Role.Name,
			Description: userModel.Role.Description,
			CreatedAt:   userModel.Role.CreatedAt,
			UpdatedAt:   userModel.Role.UpdatedAt,
		}
	}

	return user
}
