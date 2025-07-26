package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
)

type StatsHandler struct {
	statsService services.StatsService
}

func NewStatsHandler(statsService services.StatsService) *StatsHandler {
	return &StatsHandler{
		statsService: statsService,
	}
}

// GetSalesStats ดูสถิติการขาย
// @Summary ดูสถิติการขาย
// @Description ดูสถิติการขาย (เฉพาะ Admin)
// @Tags Statistics
// @Accept json
// @Produce json
// @Param period query string false "ช่วงเวลา: daily, weekly, monthly, yearly" default(monthly)
// @Success 200 {object} entities.ApiResponse{data=entities.SalesStats}
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /stats/sales [get]
func (h *StatsHandler) GetSalesStats(c *fiber.Ctx) error {
	stats, err := h.statsService.GetSalesStats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงสถิติการขายได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงสถิติการขายสำเร็จ",
		Data:    stats,
	})
}

// GetProductStats ดูสถิติสินค้า
// @Summary ดูสถิติสินค้า
// @Description ดูสถิติสินค้า (เฉพาะ Admin)
// @Tags Statistics
// @Accept json
// @Produce json
// @Success 200 {object} entities.ApiResponse{data=entities.ProductStats}
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /stats/products [get]
func (h *StatsHandler) GetProductStats(c *fiber.Ctx) error {
	stats, err := h.statsService.GetProductStats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงสถิติสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงสถิติสินค้าสำเร็จ",
		Data:    stats,
	})
}

// GetUserStats ดูสถิติผู้ใช้
// @Summary ดูสถิติผู้ใช้
// @Description ดูสถิติผู้ใช้ (เฉพาะ Admin)
// @Tags Statistics
// @Accept json
// @Produce json
// @Success 200 {object} entities.ApiResponse{data=entities.UserStats}
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /stats/users [get]
func (h *StatsHandler) GetUserStats(c *fiber.Ctx) error {
	stats, err := h.statsService.GetUserStats(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงสถิติผู้ใช้ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงสถิติผู้ใช้สำเร็จ",
		Data:    stats,
	})
}
