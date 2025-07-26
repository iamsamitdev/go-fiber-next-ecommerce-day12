package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

// CartService interface สำหรับการจัดการตะกร้าสินค้า
type CartService interface {
	GetCart(ctx context.Context, userID uuid.UUID) (*entities.Cart, error)
	AddToCart(ctx context.Context, userID uuid.UUID, req *entities.AddToCartRequest) error
	UpdateCartItem(ctx context.Context, cartItemID uuid.UUID, req *entities.UpdateCartItemRequest) error
	RemoveFromCart(ctx context.Context, cartItemID uuid.UUID) error
	ClearCart(ctx context.Context, userID uuid.UUID) error
}
