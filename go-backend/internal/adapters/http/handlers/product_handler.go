package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct สร้างสินค้า
// @Summary สร้างสินค้า
// @Description สร้างสินค้าใหม่ (เฉพาะ Admin)
// @Tags Products
// @Accept json
// @Produce json
// @Param request body entities.CreateProductRequest true "ข้อมูลการสร้างสินค้า"
// @Success 201 {object} entities.ApiResponse{data=entities.Product}
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req entities.CreateProductRequest
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

	product, err := h.productService.CreateProduct(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถสร้างสินค้าได้",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.ApiResponse{
		Success: true,
		Message: "สร้างสินค้าสำเร็จ",
		Data:    product,
	})
}

// GetProducts ดูสินค้าทั้งหมด
// @Summary ดูสินค้าทั้งหมด
// @Description ดูสินค้าทั้งหมดพร้อม pagination
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Product,pagination=entities.PaginationResponse}
// @Failure 500 {object} entities.ApiResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, pagination, err := h.productService.GetProducts(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลสินค้าสำเร็จ",
		Data:       products,
		Pagination: pagination,
	})
}

// GetProductByID ดูสินค้าตาม ID
// @Summary ดูสินค้าตาม ID
// @Description ดูรายละเอียดสินค้าตาม ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entities.ApiResponse{data=entities.Product}
// @Failure 400 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	product, err := h.productService.GetProductByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่พบสินค้า",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ดึงข้อมูลสินค้าสำเร็จ",
		Data:    product,
	})
}

// GetProductsByCategory ดูสินค้าตามหมวดหมู่
// @Summary ดูสินค้าตามหมวดหมู่
// @Description ดูสินค้าที่กรองตามหมวดหมู่
// @Tags Products
// @Accept json
// @Produce json
// @Param categoryId path string true "Category ID"
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Product,pagination=entities.PaginationResponse}
// @Failure 400 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Router /products/category/{categoryId} [get]
func (h *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	categoryID, err := uuid.Parse(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ Category ID ไม่ถูกต้อง",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, pagination, err := h.productService.GetProductsByCategory(c.Context(), categoryID, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถดึงข้อมูลสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ดึงข้อมูลสินค้าตามหมวดหมู่สำเร็จ",
		Data:       products,
		Pagination: pagination,
	})
}

// SearchProducts ค้นหาสินค้า
// @Summary ค้นหาสินค้า
// @Description ค้นหาสินค้าตามชื่อ คำอธิบาย หรือเงื่อนไขอื่นๆ
// @Tags Products
// @Accept json
// @Produce json
// @Param search query string false "คำค้นหา"
// @Param category_id query string false "Category ID"
// @Param min_price query number false "ราคาต่ำสุด"
// @Param max_price query number false "ราคาสูงสุด"
// @Param page query int false "หน้าที่ต้องการ" default(1)
// @Param limit query int false "จำนวนรายการต่อหน้า" default(10)
// @Success 200 {object} entities.ApiResponse{data=[]entities.Product,pagination=entities.PaginationResponse}
// @Failure 500 {object} entities.ApiResponse
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	req := &entities.ProductSearchRequest{
		Query:    c.Query("search"),
		Page:     1,
		Limit:    10,
		MinPrice: 0,
		MaxPrice: 0,
	}

	if page, err := strconv.Atoi(c.Query("page", "1")); err == nil && page > 0 {
		req.Page = page
	}

	if limit, err := strconv.Atoi(c.Query("limit", "10")); err == nil && limit > 0 && limit <= 100 {
		req.Limit = limit
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := uuid.Parse(categoryID); err == nil {
			req.CategoryID = id
		}
	}

	if minPrice, err := strconv.ParseFloat(c.Query("min_price"), 64); err == nil && minPrice >= 0 {
		req.MinPrice = minPrice
	}

	if maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64); err == nil && maxPrice >= 0 {
		req.MaxPrice = maxPrice
	}

	products, pagination, err := h.productService.SearchProducts(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถค้นหาสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success:    true,
		Message:    "ค้นหาสินค้าสำเร็จ",
		Data:       products,
		Pagination: pagination,
	})
}

// UpdateProduct แก้ไขสินค้า
// @Summary แก้ไขสินค้า
// @Description แก้ไขข้อมูลสินค้า (เฉพาะ Admin)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body entities.UpdateProductRequest true "ข้อมูลการแก้ไขสินค้า"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	var req entities.UpdateProductRequest
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

	if err := h.productService.UpdateProduct(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถอัพเดทสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "อัพเดทสินค้าสำเร็จ",
	})
}

// DeleteProduct ลบสินค้า
// @Summary ลบสินค้า
// @Description ลบสินค้าตาม ID (เฉพาะ Admin)
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entities.ApiResponse
// @Failure 400 {object} entities.ApiResponse
// @Failure 401 {object} entities.ApiResponse
// @Failure 403 {object} entities.ApiResponse
// @Failure 404 {object} entities.ApiResponse
// @Failure 500 {object} entities.ApiResponse
// @Security BearerAuth
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: "รูปแบบ ID ไม่ถูกต้อง",
		})
	}

	if err := h.productService.DeleteProduct(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: "ไม่สามารถลบสินค้าได้",
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "ลบสินค้าสำเร็จ",
	})
}
