package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type cartService struct {
	cartRepo repositories.CartRepository
}

func NewCartService(cartRepo repositories.CartRepository) services.CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

func (s *cartService) GetCart(ctx context.Context, userID uuid.UUID) (*entities.Cart, error) {
	return s.cartRepo.GetByUserID(ctx, userID)
}

func (s *cartService) AddToCart(ctx context.Context, userID uuid.UUID, req *entities.AddToCartRequest) error {
	return s.cartRepo.AddItem(ctx, userID, req)
}

func (s *cartService) UpdateCartItem(ctx context.Context, cartItemID uuid.UUID, req *entities.UpdateCartItemRequest) error {
	return s.cartRepo.UpdateItem(ctx, cartItemID, req.Quantity)
}

func (s *cartService) RemoveFromCart(ctx context.Context, cartItemID uuid.UUID) error {
	return s.cartRepo.RemoveItem(ctx, cartItemID)
}

func (s *cartService) ClearCart(ctx context.Context, userID uuid.UUID) error {
	return s.cartRepo.ClearCart(ctx, userID)
}
