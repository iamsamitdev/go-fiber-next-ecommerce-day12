package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repositories.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entities.Cart, error) {
	var cart models.Cart

	// หาตะกร้าของผู้ใช้ หากไม่มีให้สร้างใหม่
	if err := r.db.WithContext(ctx).Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// สร้างตะกร้าใหม่
			newCart := &models.Cart{
				UserID:     userID,
				TotalPrice: 0,
			}
			if err := r.db.WithContext(ctx).Create(newCart).Error; err != nil {
				return nil, err
			}
			return r.modelToEntity(newCart), nil
		}
		return nil, err
	}

	return r.modelToEntity(&cart), nil
}

func (r *cartRepository) AddItem(ctx context.Context, userID uuid.UUID, item *entities.AddToCartRequest) error {
	// หาตะกร้าของผู้ใช้
	cart, err := r.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// ตรวจสอบว่าสินค้ามีอยู่หรือไม่
	var product models.Product
	if err := r.db.WithContext(ctx).First(&product, "id = ?", item.ProductID).Error; err != nil {
		return err
	}

	// ตรวจสอบสต็อก
	if product.Stock < item.Quantity {
		return gorm.ErrInvalidData
	}

	// ตรวจสอบว่าสินค้านี้มีในตะกร้าแล้วหรือไม่
	var existingItem models.CartItem
	if err := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cart.ID, item.ProductID).First(&existingItem).Error; err == nil {
		// อัพเดทจำนวน
		newQuantity := existingItem.Quantity + item.Quantity
		if product.Stock < newQuantity {
			return gorm.ErrInvalidData
		}
		return r.db.WithContext(ctx).Model(&existingItem).Updates(map[string]interface{}{
			"quantity": newQuantity,
			"price":    product.Price,
		}).Error
	}

	// เพิ่มสินค้าใหม่ลงตะกร้า
	cartItem := &models.CartItem{
		CartID:    cart.ID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     product.Price,
	}

	return r.db.WithContext(ctx).Create(cartItem).Error
}

func (r *cartRepository) UpdateItem(ctx context.Context, cartItemID uuid.UUID, quantity int) error {
	// หา cart item
	var cartItem models.CartItem
	if err := r.db.WithContext(ctx).Preload("Product").First(&cartItem, "id = ?", cartItemID).Error; err != nil {
		return err
	}

	// ตรวจสอบสต็อก
	if cartItem.Product.Stock < quantity {
		return gorm.ErrInvalidData
	}

	return r.db.WithContext(ctx).Model(&cartItem).Update("quantity", quantity).Error
}

func (r *cartRepository) RemoveItem(ctx context.Context, cartItemID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.CartItem{}, "id = ?", cartItemID).Error
}

func (r *cartRepository) ClearCart(ctx context.Context, userID uuid.UUID) error {
	// หาตะกร้าของผู้ใช้
	var cart models.Cart
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}

	// ลบรายการทั้งหมดในตะกร้า
	return r.db.WithContext(ctx).Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
}

func (r *cartRepository) GetCartItem(ctx context.Context, cartItemID uuid.UUID) (*entities.CartItem, error) {
	var cartItem models.CartItem
	if err := r.db.WithContext(ctx).Preload("Product").First(&cartItem, "id = ?", cartItemID).Error; err != nil {
		return nil, err
	}

	return r.cartItemModelToEntity(&cartItem), nil
}

func (r *cartRepository) modelToEntity(cart *models.Cart) *entities.Cart {
	cartEntity := &entities.Cart{
		ID:         cart.ID,
		UserID:     cart.UserID,
		TotalPrice: cart.TotalPrice,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}

	// คำนวณราคารวม
	var totalPrice float64
	for _, item := range cart.CartItems {
		cartEntity.CartItems = append(cartEntity.CartItems, *r.cartItemModelToEntity(&item))
		totalPrice += item.Price * float64(item.Quantity)
	}
	cartEntity.TotalPrice = totalPrice

	return cartEntity
}

func (r *cartRepository) cartItemModelToEntity(cartItem *models.CartItem) *entities.CartItem {
	item := &entities.CartItem{
		ID:        cartItem.ID,
		CartID:    cartItem.CartID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
		Price:     cartItem.Price,
		CreatedAt: cartItem.CreatedAt,
		UpdatedAt: cartItem.UpdatedAt,
	}

	if cartItem.Product.ID != uuid.Nil {
		item.Product = &entities.Product{
			ID:          cartItem.Product.ID,
			Name:        cartItem.Product.Name,
			Description: cartItem.Product.Description,
			Price:       cartItem.Product.Price,
			Stock:       cartItem.Product.Stock,
			Image:       cartItem.Product.Image,
			CategoryID:  cartItem.Product.CategoryID,
			CreatedAt:   cartItem.Product.CreatedAt,
			UpdatedAt:   cartItem.Product.UpdatedAt,
		}
	}

	return item
}
