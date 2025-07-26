package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// AuthService interface สำหรับการจัดการการยืนยันตัวตน
type AuthService interface {
	Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error)
	AdminRegister(ctx context.Context, req *entities.AdminRegisterRequest) (*entities.User, error)
	Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
	RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error)
	Logout(ctx context.Context, userID uuid.UUID) error
	ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePasswordRequest) error
	ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error
	ValidateToken(ctx context.Context, token string) (*entities.User, error)
}
