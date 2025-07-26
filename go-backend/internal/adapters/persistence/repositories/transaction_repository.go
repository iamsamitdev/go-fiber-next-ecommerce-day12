package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repositories.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, req *entities.CreatePaymentRequest) (*entities.Transaction, error) {
	// ตรวจสอบว่าคำสั่งซื้อมีอยู่จริง
	var order models.Order
	if err := r.db.WithContext(ctx).First(&order, "id = ?", req.OrderID).Error; err != nil {
		return nil, err
	}

	// สร้าง transaction ID แบบ unique
	transactionID := fmt.Sprintf("TXN_%d_%s", time.Now().Unix(), uuid.New().String()[:8])

	transaction := &models.Transaction{
		OrderID:       req.OrderID,
		Amount:        order.TotalPrice,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
		TransactionID: transactionID,
		PaymentData:   req.PaymentData,
	}

	if err := r.db.WithContext(ctx).Create(transaction).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(transaction), nil
}

func (r *transactionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.WithContext(ctx).First(&transaction, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&transaction), nil
}

func (r *transactionRepository) GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]*entities.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	var result []*entities.Transaction
	for _, transaction := range transactions {
		result = append(result, r.modelToEntity(&transaction))
	}

	return result, nil
}

func (r *transactionRepository) GetByTransactionID(ctx context.Context, transactionID string) (*entities.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&transaction).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&transaction), nil
}

func (r *transactionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	tx := r.db.WithContext(ctx).Begin()

	// อัพเดทสถานะ transaction
	if err := tx.Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	// หา transaction เพื่อดึง order ID
	var transaction models.Transaction
	if err := tx.First(&transaction, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// อัพเดทสถานะการชำระเงินของคำสั่งซื้อ
	var paymentStatus string
	switch status {
	case "completed":
		paymentStatus = "paid"
	case "failed":
		paymentStatus = "failed"
	case "cancelled":
		paymentStatus = "cancelled"
	default:
		paymentStatus = "pending"
	}

	if err := tx.Model(&models.Order{}).Where("id = ?", transaction.OrderID).Update("payment_status", paymentStatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *transactionRepository) Cancel(ctx context.Context, id uuid.UUID) error {
	return r.UpdateStatus(ctx, id, "cancelled")
}

func (r *transactionRepository) modelToEntity(transaction *models.Transaction) *entities.Transaction {
	return &entities.Transaction{
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
}
