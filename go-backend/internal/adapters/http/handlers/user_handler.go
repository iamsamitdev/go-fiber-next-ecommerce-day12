package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers ดูผู้ใช้ทั้งหมด
// @Summary ดูผู้ใช้ทั้งหมด
// @Description ดูผู้ใช้ทั้งหมดพร้อม pagination
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.User,pagination=entities.PaginationResponse}
// @Failure 400 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, pagination, err := h.userService.GetUsers(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลผู้ใช้ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลผู้ใช้สำเร็จ",
		Data:       users,
		Pagination: pagination,
	})
}

// GetUserByID ดูผู้ใช้ตาม ID
// @Summary ดูผู้ใช้ตาม ID
// @Description ดูรายละเอียดผู้ใช้ตาม ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entities.ApiResponse{data=entities.User}
// @Failure 400 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่พบผู้ใช้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลผู้ใช้สำเร็จ",
		Data:    user,
	})
}

// UpdateUser แก้ไขผู้ใช้
// @Summary แก้ไขผู้ใช้
// @Description แก้ไขข้อมูลผู้ใช้
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body entities.UpdateUserRequest true "ข้อมูลการแก้ไขผู้ใช้"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	var req entities.UpdateUserRequest
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

	if err := h.userService.UpdateUser(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดทข้อมูลผู้ใช้ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดทข้อมูลผู้ใช้สำเร็จ",
	})
}

// DeleteUser ลบผู้ใช้
// @Summary ลบผู้ใช้
// @Description ลบผู้ใช้ตาม ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถลบผู้ใช้ได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบผู้ใช้สำเร็จ",
	})
}
