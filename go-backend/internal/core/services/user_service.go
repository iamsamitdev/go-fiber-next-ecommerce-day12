package services

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUsers(ctx context.Context, page, limit int) ([]*entities.User, *entities.PaginationResponse, error) {
	users, total, err := s.userRepo.GetAll(ctx, page, limit)
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

	return users, pagination, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, req *entities.UpdateUserRequest) error {
	return s.userRepo.Update(ctx, id, req)
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}
