package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

// CreatePayment สร้างการชำระเงิน
// @Summary สร้างการชำระเงิน
// @Description สร้างการชำระเงินสำหรับคำสั่งซื้อ
// @Tags Payments
// @Accept json
// @Produce json
// @Param request body entities.CreatePaymentRequest true "ข้อมูลการสร้างการชำระเงิน"
// @Success 201 {object} entities.ApiResponse{data=entities.Transaction}
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /payments [post]
func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	var req entities.CreatePaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	transaction, err := h.paymentService.CreatePayment(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถสร้างการชำระเงินได้",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "สร้างการชำระเงินสำเร็จ",
		Data:    transaction,
	})
}

// VerifyPayment ยืนยันการชำระเงิน
// @Summary ยืนยันการชำระเงิน
// @Description ยืนยันการชำระเงิน
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param request body entities.VerifyPaymentRequest true "ข้อมูลการยืนยันการชำระเงิน"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /payments/{id}/verify [post]
func (h *PaymentHandler) VerifyPayment(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	var req entities.VerifyPaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := h.paymentService.VerifyPayment(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถยืนยันการชำระเงินได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ยืนยันการชำระเงินสำเร็จ",
	})
}

// CancelPayment ยกเลิกการชำระเงิน
// @Summary ยกเลิกการชำระเงิน
// @Description ยกเลิกการชำระเงิน
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /payments/{id}/cancel [put]
func (h *PaymentHandler) CancelPayment(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	if err := h.paymentService.CancelPayment(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถยกเลิกการชำระเงินได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ยกเลิกการชำระเงินสำเร็จ",
	})
}
