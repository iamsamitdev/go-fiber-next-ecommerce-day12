package services

import (
	"context"

	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// StatsService interface สำหรับสถิติ
type StatsService interface {
	GetSalesStats(ctx context.Context) (*entities.SalesStats, error)
	GetProductStats(ctx context.Context) (*entities.ProductStats, error)
	GetUserStats(ctx context.Context) (*entities.UserStats, error)
}
