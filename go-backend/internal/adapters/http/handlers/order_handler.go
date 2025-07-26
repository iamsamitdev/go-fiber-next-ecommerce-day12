package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder สร้างคำสั่งซื้อ
// @Summary สร้างคำสั่งซื้อ
// @Description สร้างคำสั่งซื้อใหม่จากตะกร้าสินค้า
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body entities.CreateOrderRequest true "ข้อมูลการสร้างคำสั่งซื้อ"
// @Success 201 {object} entities.ApiResponse{data=entities.Order}
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req entities.CreateOrderRequest
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

	order, err := h.orderService.CreateOrder(c.Context(), userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถสร้างคำสั่งซื้อได้",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "สร้างคำสั่งซื้อสำเร็จ",
		Data:    order,
	})
}

// GetOrders ดูคำสั่งซื้อของผู้ใช้
// @Summary ดูคำสั่งซื้อของผู้ใช้
// @Description ดูคำสั่งซื้อของผู้ใช้ปัจจุบัน
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Order,pagination=entities.PaginationResponse}
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders [get]
func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	orders, pagination, err := h.orderService.GetOrders(c.Context(), userID, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลคำสั่งซื้อได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลคำสั่งซื้อสำเร็จ",
		Data:       orders,
		Pagination: pagination,
	})
}

// GetOrderByID ดูคำสั่งซื้อตาม ID
// @Summary ดูคำสั่งซื้อตาม ID
// @Description ดูรายละเอียดคำสั่งซื้อตาม ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} entities.ApiResponse{data=entities.Order}
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	order, err := h.orderService.GetOrderByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่พบคำสั่งซื้อ",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลคำสั่งซื้อสำเร็จ",
		Data:    order,
	})
}

// CancelOrder ยกเลิกคำสั่งซื้อ
// @Summary ยกเลิกคำสั่งซื้อ
// @Description ยกเลิกคำสั่งซื้อ
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders/{id}/cancel [put]
func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	if err := h.orderService.CancelOrder(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถยกเลิกคำสั่งซื้อได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ยกเลิกคำสั่งซื้อสำเร็จ",
	})
}

// GetAllOrders ดูคำสั่งซื้อทั้งหมด (Admin)
// @Summary ดูคำสั่งซื้อทั้งหมด (Admin)
// @Description ดูคำสั่งซื้อทั้งหมดพร้อม pagination (เฉพาะ Admin)
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Order,pagination=entities.PaginationResponse}
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders/admin [get]
func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	orders, pagination, err := h.orderService.GetAllOrders(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลคำสั่งซื้อได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลคำสั่งซื้อทั้งหมดสำเร็จ",
		Data:       orders,
		Pagination: pagination,
	})
}

// UpdateOrderStatus อัพเดทสถานะคำสั่งซื้อ (Admin)
// @Summary อัพเดทสถานะคำสั่งซื้อ (Admin)
// @Description อัพเดทสถานะคำสั่งซื้อ (เฉพาะ Admin)
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param request body entities.UpdateOrderStatusRequest true "ข้อมูลการอัพเดทสถานะคำสั่งซื้อ"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /orders/admin/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	var req entities.UpdateOrderStatusRequest
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

	if err := h.orderService.UpdateOrderStatus(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดทสถานะคำสั่งซื้อได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดทสถานะคำสั่งซื้อสำเร็จ",
	})
}
