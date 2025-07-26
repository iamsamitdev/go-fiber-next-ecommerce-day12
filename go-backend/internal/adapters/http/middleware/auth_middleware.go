package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: jwtSecret,
	}
}

// AuthRequired middleware ตรวจสอบ JWT token
func (m *AuthMiddleware) AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "ไม่พบ Authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "รูปแบบ token ไม่ถูกต้อง",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "Token ไม่ถูกต้องหรือหมดอายุ",
			})
		}

		claims, ok := token.Claims.(*utils.Claims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "Claims ไม่ถูกต้อง",
			})
		}

		// แปลง UserID เป็น UUID
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: "User ID ไม่ถูกต้อง",
			})
		}

		// เก็บข้อมูลผู้ใช้ใน context
		c.Locals("userID", userID)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// AdminRequired middleware ตรวจสอบว่าเป็น admin
func (m *AuthMiddleware) AdminRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(entities.ApiResponse{
				Success: false,
				Message: "ไม่มีสิทธิ์เข้าถึง",
			})
		}

		return c.Next()
	}
}

// RoleRequired middleware ตรวจสอบ role ที่กำหนด
func (m *AuthMiddleware) RoleRequired(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(entities.ApiResponse{
				Success: false,
				Message: "ไม่พบข้อมูล role",
			})
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่มีสิทธิ์เข้าถึง",
		})
	}
}
