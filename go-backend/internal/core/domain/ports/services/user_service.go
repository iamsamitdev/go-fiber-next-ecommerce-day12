package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// UserService interface สำหรับการจัดการผู้ใช้
type UserService interface {
	GetUsers(ctx context.Context, page, limit int) ([]*entities.User, *entities.PaginationResponse, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req *entities.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
