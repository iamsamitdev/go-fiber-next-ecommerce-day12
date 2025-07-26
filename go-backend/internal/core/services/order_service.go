package services

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) services.OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, userID uuid.UUID, req *entities.CreateOrderRequest) (*entities.Order, error) {
	return s.orderRepo.Create(ctx, userID, req)
}

func (s *orderService) GetOrders(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error) {
	orders, total, err := s.orderRepo.GetByUserID(ctx, userID, page, limit)
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

	return orders, pagination, nil
}

func (s *orderService) GetOrderByID(ctx context.Context, id uuid.UUID) (*entities.Order, error) {
	return s.orderRepo.GetByID(ctx, id)
}

func (s *orderService) CancelOrder(ctx context.Context, id uuid.UUID) error {
	return s.orderRepo.Cancel(ctx, id)
}

func (s *orderService) GetAllOrders(ctx context.Context, page, limit int) ([]*entities.Order, *entities.PaginationResponse, error) {
	orders, total, err := s.orderRepo.GetAll(ctx, page, limit)
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

	return orders, pagination, nil
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateOrderStatusRequest) error {
	return s.orderRepo.UpdateStatus(ctx, id, req.Status)
}

func (s *orderService) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, req *entities.UpdatePaymentStatusRequest) error {
	return s.orderRepo.UpdatePaymentStatus(ctx, id, req.PaymentStatus)
}

func (s *orderService) UpdateShippingStatus(ctx context.Context, id uuid.UUID, req *entities.UpdateShippingStatusRequest) error {
	return s.orderRepo.UpdateShippingStatus(ctx, id, req.ShippingStatus, req.TrackingNumber)
}
