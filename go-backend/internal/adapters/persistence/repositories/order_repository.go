package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(ctx context.Context, userID uuid.UUID, req *entities.CreateOrderRequest) (*entities.Order, error) {
	tx := r.db.WithContext(ctx).Begin()

	// หาตะกร้าของผู้ใช้
	var cart models.Cart
	if err := tx.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(cart.CartItems) == 0 {
		tx.Rollback()
		return nil, errors.New("ตะกร้าสินค้าว่าง")
	}

	// คำนวณราคารวม
	var totalPrice float64
	for _, item := range cart.CartItems {
		totalPrice += item.Price * float64(item.Quantity)
	}

	// สร้างคำสั่งซื้อ
	order := &models.Order{
		UserID:          userID,
		TotalPrice:      totalPrice,
		Status:          "pending",
		PaymentMethod:   req.PaymentMethod,
		PaymentStatus:   "pending",
		ShippingMethod:  req.ShippingMethod,
		ShippingStatus:  "pending",
		ShippingAddress: req.ShippingAddress,
		Notes:           req.Notes,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// สร้างรายการสินค้าในคำสั่งซื้อ
	for _, cartItem := range cart.CartItems {
		orderItem := &models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}

		if err := tx.Create(orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// อัพเดทสต็อกสินค้า
		if err := tx.Model(&models.Product{}).Where("id = ?", cartItem.ProductID).Update("stock", gorm.Expr("stock - ?", cartItem.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// ล้างตะกร้าสินค้า
	if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return r.GetByID(ctx, order.ID)
}

func (r *orderRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).Preload("User").Preload("OrderItems.Product").Preload("Transactions").First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&order), nil
}

func (r *orderRepository) GetByUserID(ctx context.Context, userID uuid.UUID, page, limit int) ([]*entities.Order, int, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Order{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("OrderItems.Product").Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Order
	for _, order := range orders {
		result = append(result, r.modelToEntity(&order))
	}

	return result, int(total), nil
}

func (r *orderRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Order, int, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("User").Preload("OrderItems.Product").Order("created_at DESC").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Order
	for _, order := range orders {
		result = append(result, r.modelToEntity(&order))
	}

	return result, int(total), nil
}

func (r *orderRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepository) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, paymentStatus string) error {
	return r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Update("payment_status", paymentStatus).Error
}

func (r *orderRepository) UpdateShippingStatus(ctx context.Context, id uuid.UUID, shippingStatus, trackingNumber string) error {
	updates := map[string]interface{}{
		"shipping_status": shippingStatus,
	}
	if trackingNumber != "" {
		updates["tracking_number"] = trackingNumber
	}

	return r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Updates(updates).Error
}

func (r *orderRepository) Cancel(ctx context.Context, id uuid.UUID) error {
	tx := r.db.WithContext(ctx).Begin()

	// ดึงข้อมูลคำสั่งซื้อ
	var order models.Order
	if err := tx.Preload("OrderItems").First(&order, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// ตรวจสอบสถานะ
	if order.Status != "pending" {
		tx.Rollback()
		return errors.New("ไม่สามารถยกเลิกคำสั่งซื้อนี้ได้")
	}

	// คืนสต็อกสินค้า
	for _, item := range order.OrderItems {
		if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// อัพเดทสถานะ
	if err := tx.Model(&order).Update("status", "cancelled").Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *orderRepository) modelToEntity(order *models.Order) *entities.Order {
	orderEntity := &entities.Order{
		ID:              order.ID,
		UserID:          order.UserID,
		TotalPrice:      order.TotalPrice,
		Status:          order.Status,
		PaymentMethod:   order.PaymentMethod,
		PaymentStatus:   order.PaymentStatus,
		ShippingMethod:  order.ShippingMethod,
		ShippingStatus:  order.ShippingStatus,
		ShippingAddress: order.ShippingAddress,
		TrackingNumber:  order.TrackingNumber,
		Notes:           order.Notes,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	if order.User.ID != uuid.Nil {
		orderEntity.User = &entities.User{
			ID:        order.User.ID,
			Email:     order.User.Email,
			FirstName: order.User.FirstName,
			LastName:  order.User.LastName,
			Phone:     order.User.Phone,
			Address:   order.User.Address,
			Active:    order.User.Active,
			RoleID:    order.User.RoleID,
			CreatedAt: order.User.CreatedAt,
			UpdatedAt: order.User.UpdatedAt,
		}
	}

	for _, item := range order.OrderItems {
		orderItem := entities.OrderItem{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		if item.Product.ID != uuid.Nil {
			orderItem.Product = &entities.Product{
				ID:          item.Product.ID,
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				Image:       item.Product.Image,
				CategoryID:  item.Product.CategoryID,
				CreatedAt:   item.Product.CreatedAt,
				UpdatedAt:   item.Product.UpdatedAt,
			}
		}

		orderEntity.OrderItems = append(orderEntity.OrderItems, orderItem)
	}

	for _, transaction := range order.Transactions {
		transactionEntity := entities.Transaction{
			ID:            transaction.ID,
			OrderID:       transaction.OrderID,
			Amount:        transaction.Amount,
			PaymentMethod: transaction.PaymentMethod,
			Status:        transaction.Status,
			TransactionID: transaction.TransactionID,
			PaymentData:   transaction.PaymentData,
			CreatedAt:     transaction.CreatedAt,
			UpdatedAt:     transaction.UpdatedAt,
		}
		orderEntity.Transactions = append(orderEntity.Transactions, transactionEntity)
	}

	return orderEntity
}
