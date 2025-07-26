package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthHandler(authService services.AuthService, userService services.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

// Register ลงทะเบียนผู้ใช้ใหม่
// @Summary ลงทะเบียนผู้ใช้ใหม่
// @Description ลงทะเบียนผู้ใช้ใหม่ในระบบ
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.RegisterRequest true "ข้อมูลการลงทะเบียน"
// @Success 201 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req entities.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	user, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถลงทะเบียนได้",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "ลงทะเบียนสำเร็จ",
		Data:    user,
	})
}

// AdminRegister ลงทะเบียนผู้ดูแลระบบ
// @Summary ลงทะเบียนผู้ดูแลระบบ
// @Description ลงทะเบียนผู้ดูแลระบบใหม่ (เฉพาะ Admin เท่านั้น)
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entities.AdminRegisterRequest true "ข้อมูลการลงทะเบียนผู้ดูแลระบบ"
// @Success 201 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 403 {object} entities.ErrorResponse
// @Router /auth/admin/register [post]
func (h *AuthHandler) AdminRegister(c *fiber.Ctx) error {
	var req entities.AdminRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	user, err := h.authService.AdminRegister(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถลงทะเบียนได้",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "ลงทะเบียนผู้ดูแลระบบสำเร็จ",
		Data:    user,
	})
}

// Login เข้าสู่ระบบ
// @Summary เข้าสู่ระบบ
// @Description เข้าสู่ระบบด้วยอีเมลและรหัสผ่าน
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.LoginRequest true "ข้อมูลการเข้าสู่ระบบ"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req entities.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	response, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถเข้าสู่ระบบได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "เข้าสู่ระบบสำเร็จ",
		Data:    response,
	})
}

// RefreshToken รีเฟรช token
// @Summary รีเฟรช token
// @Description รีเฟรช JWT token ด้วย refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req entities.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	response, err := h.authService.RefreshToken(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถรีเฟรช token ได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "รีเฟรช token สำเร็จ",
		Data:    response,
	})
}

// Logout ออกจากระบบ
// @Summary ออกจากระบบ
// @Description ออกจากระบบและลบ refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} entities.ApiResponse
// @Failure 401 {object} entities.ErrorResponse
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	id, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบ user ID ไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := h.authService.Logout(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถออกจากระบบได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ออกจากระบบสำเร็จ",
	})
}

// ChangePassword เปลี่ยนรหัสผ่าน
// @Summary เปลี่ยนรหัสผ่าน
// @Description เปลี่ยนรหัสผ่านของผู้ใช้
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entities.ChangePasswordRequest true "ข้อมูลการเปลี่ยนรหัสผ่าน"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/change-password [post]
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	var req entities.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	userID := c.Locals("userID").(string)
	id, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบ user ID ไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := h.authService.ChangePassword(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถเปลี่ยนรหัสผ่านได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "เปลี่ยนรหัสผ่านสำเร็จ",
	})
}

// ForgotPassword ลืมรหัสผ่าน
// @Summary ลืมรหัสผ่าน
// @Description ส่งลิงก์รีเซ็ตรหัสผ่านไปยังอีเมล
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.ForgotPasswordRequest true "อีเมลสำหรับรีเซ็ตรหัสผ่าน"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	var req entities.ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	if err := h.authService.ForgotPassword(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถส่งลิงก์รีเซ็ตรหัสผ่านได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ส่งลิงก์รีเซ็ตรหัสผ่านไปยังอีเมลของคุณแล้ว",
	})
}

// ResetPassword รีเซ็ตรหัสผ่าน
// @Summary รีเซ็ตรหัสผ่าน
// @Description รีเซ็ตรหัสผ่านด้วย token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entities.ResetPasswordRequest true "ข้อมูลการรีเซ็ตรหัสผ่าน"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req entities.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ครบถ้วน",
			Error:   err.Error(),
		})
	}

	if err := h.authService.ResetPassword(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถรีเซ็ตรหัสผ่านได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "รีเซ็ตรหัสผ่านสำเร็จ",
	})
}

// GetUsers ดูรายการผู้ใช้ทั้งหมด
// @Summary ดูรายการผู้ใช้ทั้งหมด
// @Description ดูรายการผู้ใช้ทั้งหมดในระบบ (เฉพาะ Admin)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /users [get]
func (h *AuthHandler) GetUsers(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลผู้ใช้ได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลผู้ใช้สำเร็จ",
		Data:       users,
		Pagination: pagination,
	})
}

// GetUserByID ดูข้อมูลผู้ใช้รายบุคคล
// @Summary ดูข้อมูลผู้ใช้รายบุคคล
// @Description ดูข้อมูลผู้ใช้รายบุคคลตาม ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /users/{id} [get]
func (h *AuthHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่พบผู้ใช้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลผู้ใช้สำเร็จ",
		Data:    user,
	})
}

// UpdateUser แก้ไขข้อมูลผู้ใช้
// @Summary แก้ไขข้อมูลผู้ใช้
// @Description แก้ไขข้อมูลผู้ใช้
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body entities.UpdateUserRequest true "ข้อมูลที่ต้องการแก้ไข"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /users/{id} [put]
func (h *AuthHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	var req entities.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ข้อมูลไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := h.userService.UpdateUser(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถแก้ไขข้อมูลผู้ใช้ได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "แก้ไขข้อมูลผู้ใช้สำเร็จ",
	})
}

// DeleteUser ลบผู้ใช้
// @Summary ลบผู้ใช้
// @Description ลบผู้ใช้ออกจากระบบ (เฉพาะ Admin)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Router /users/{id} [delete]
func (h *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
			Error:   err.Error(),
		})
	}

	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ErrorResponse{
			Success: false,
			Message: "ไม่สามารถลบผู้ใช้ได้",
			Error:   err.Error(),
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบผู้ใช้สำเร็จ",
	})
}
