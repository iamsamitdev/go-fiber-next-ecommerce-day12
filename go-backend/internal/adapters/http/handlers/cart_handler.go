package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

// GetCart ดูตะกร้าสินค้า
// @Summary ดูตะกร้าสินค้า
// @Description ดูตะกร้าสินค้าของผู้ใช้ปัจจุบัน
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} entities.ApiResponse{data=entities.Cart}
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /cart [get]
func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	cart, err := h.cartService.GetCart(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลตะกร้าสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลตะกร้าสินค้าสำเร็จ",
		Data:    cart,
	})
}

// AddToCart เพิ่มสินค้าลงตะกร้า
// @Summary เพิ่มสินค้าลงตะกร้า
// @Description เพิ่มสินค้าลงในตะกร้าสินค้า
// @Tags Cart
// @Accept json
// @Produce json
// @Param request body entities.AddToCartRequest true "ข้อมูลการเพิ่มสินค้าลงตะกร้า"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /cart [post]
func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req entities.AddToCartRequest
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

	if err := h.cartService.AddToCart(c.Context(), userID, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถเพิ่มสินค้าลงตะกร้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "เพิ่มสินค้าลงตะกร้าสำเร็จ",
	})
}

// UpdateCartItem อัพเดทสินค้าในตะกร้า
// @Summary อัพเดทสินค้าในตะกร้า
// @Description อัพเดทจำนวนสินค้าในตะกร้า
// @Tags Cart
// @Accept json
// @Produce json
// @Param itemId path string true "Cart Item ID"
// @Param request body entities.UpdateCartItemRequest true "ข้อมูลการอัพเดทสินค้าในตะกร้า"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /cart/{itemId} [put]
func (h *CartHandler) UpdateCartItem(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("itemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ Item ID ไม่ถูกต้อง",
		})
	}

	var req entities.UpdateCartItemRequest
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

	if err := h.cartService.UpdateCartItem(c.Context(), itemID, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดทสินค้าในตะกร้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดทสินค้าในตะกร้าสำเร็จ",
	})
}

// RemoveFromCart ลบสินค้าจากตะกร้า
// @Summary ลบสินค้าจากตะกร้า
// @Description ลบสินค้าออกจากตะกร้าสินค้า
// @Tags Cart
// @Accept json
// @Produce json
// @Param itemId path string true "Cart Item ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /cart/{itemId} [delete]
func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("itemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ Item ID ไม่ถูกต้อง",
		})
	}

	if err := h.cartService.RemoveFromCart(c.Context(), itemID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถลบสินค้าจากตะกร้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบสินค้าจากตะกร้าสำเร็จ",
	})
}

// ClearCart ล้างตะกร้าสินค้า
// @Summary ล้างตะกร้าสินค้า
// @Description ลบสินค้าทั้งหมดออกจากตะกร้าสินค้า
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /cart [delete]
func (h *CartHandler) ClearCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	if err := h.cartService.ClearCart(c.Context(), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถล้างตะกร้าสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ล้างตะกร้าสินค้าสำเร็จ",
	})
}
