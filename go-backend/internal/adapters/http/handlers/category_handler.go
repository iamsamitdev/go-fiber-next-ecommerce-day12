package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory สร้างหมวดหมู่
// @Summary สร้างหมวดหมู่
// @Description สร้างหมวดหมู่ใหม่ (เฉพาะ Admin)
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body entities.CreateCategoryRequest true "ข้อมูลการสร้างหมวดหมู่"
// @Success 201 {object} entities.ApiResponse{data=entities.Category}
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req entities.CreateCategoryRequest
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

	category, err := h.categoryService.CreateCategory(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถสร้างหมวดหมู่ได้",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "สร้างหมวดหมู่สำเร็จ",
		Data:    category,
	})
}

// GetCategories ดูหมวดหมู่ทั้งหมด
// @Summary ดูหมวดหมู่ทั้งหมด
// @Description ดูหมวดหมู่ทั้งหมดพร้อม pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Category,pagination=entities.PaginationResponse}
// @Failure 500 {object} entities.ApiResponse
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	categories, pagination, err := h.categoryService.GetCategories(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลหมวดหมู่ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลหมวดหมู่สำเร็จ",
		Data:       categories,
		Pagination: pagination,
	})
}

// GetCategoryByID ดูหมวดหมู่ตาม ID
// @Summary ดูหมวดหมู่ตาม ID
// @Description ดูรายละเอียดหมวดหมู่ตาม ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entities.ApiResponse{data=entities.Category}
// @Failure 400 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	category, err := h.categoryService.GetCategoryByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่พบหมวดหมู่",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลหมวดหมู่สำเร็จ",
		Data:    category,
	})
}

// UpdateCategory แก้ไขหมวดหมู่
// @Summary แก้ไขหมวดหมู่
// @Description แก้ไขข้อมูลหมวดหมู่ (เฉพาะ Admin)
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body entities.UpdateCategoryRequest true "ข้อมูลการแก้ไขหมวดหมู่"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	var req entities.UpdateCategoryRequest
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

	if err := h.categoryService.UpdateCategory(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดทหมวดหมู่ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดทหมวดหมู่สำเร็จ",
	})
}

// DeleteCategory ลบหมวดหมู่
// @Summary ลบหมวดหมู่
// @Description ลบหมวดหมู่ตาม ID (เฉพาะ Admin)
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	if err := h.categoryService.DeleteCategory(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถลบหมวดหมู่ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบหมวดหมู่สำเร็จ",
	})
}
