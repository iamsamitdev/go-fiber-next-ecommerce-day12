package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, req *entities.CreateProductRequest) (*entities.Product, error) {
	productModel := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Image:       req.Image,
		CategoryID:  req.CategoryID,
	}

	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Create(productModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// เพิ่มรูปภาพเพิ่มเติม
	for _, imageURL := range req.Images {
		productImage := &models.ProductImage{
			ProductID: productModel.ID,
			ImageURL:  imageURL,
		}
		if err := tx.Create(productImage).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return r.GetByID(ctx, productModel.ID)
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Product, error) {
	var productModel models.Product
	if err := r.db.WithContext(ctx).Preload("Category").Preload("Images").First(&productModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return r.modelToEntity(&productModel), nil
}

func (r *productRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Product, int, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Product
	for _, product := range products {
		result = append(result, r.modelToEntity(&product))
	}

	return result, int(total), nil
}

func (r *productRepository) GetByCategory(ctx context.Context, categoryID uuid.UUID, page, limit int) ([]*entities.Product, int, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("category_id = ?", categoryID).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Product
	for _, product := range products {
		result = append(result, r.modelToEntity(&product))
	}

	return result, int(total), nil
}

func (r *productRepository) Search(ctx context.Context, req *entities.ProductSearchRequest) ([]*entities.Product, int, error) {
	var products []models.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Product{})

	// ค้นหาตามชื่อสินค้า
	if req.Query != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", fmt.Sprintf("%%%s%%", req.Query), fmt.Sprintf("%%%s%%", req.Query))
	}

	// กรองตามหมวดหมู่
	if req.CategoryID != uuid.Nil {
		query = query.Where("category_id = ?", req.CategoryID)
	}

	// กรองตามราคา
	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}
	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := req.Page
	if page == 0 {
		page = 1
	}
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := query.Preload("Category").Preload("Images").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	var result []*entities.Product
	for _, product := range products {
		result = append(result, r.modelToEntity(&product))
	}

	return result, int(total), nil
}

func (r *productRepository) Update(ctx context.Context, id uuid.UUID, req *entities.UpdateProductRequest) error {
	updates := map[string]interface{}{}

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Price > 0 {
		updates["price"] = req.Price
	}
	if req.Stock >= 0 {
		updates["stock"] = req.Stock
	}
	if req.Image != "" {
		updates["image"] = req.Image
	}
	if req.CategoryID != uuid.Nil {
		updates["category_id"] = req.CategoryID
	}

	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		tx.Rollback()
		return err
	}

	// อัพเดทรูปภาพเพิ่มเติม (ลบรูปเก่าและเพิ่มรูปใหม่)
	if len(req.Images) > 0 {
		// ลบรูปเก่า
		if err := tx.Where("product_id = ?", id).Delete(&models.ProductImage{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// เพิ่มรูปใหม่
		for _, imageURL := range req.Images {
			productImage := &models.ProductImage{
				ProductID: id,
				ImageURL:  imageURL,
			}
			if err := tx.Create(productImage).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, "id = ?", id).Error
}

func (r *productRepository) UpdateStock(ctx context.Context, id uuid.UUID, stock int) error {
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", id).Update("stock", stock).Error
}

func (r *productRepository) GetLowStockProducts(ctx context.Context, threshold int) ([]*entities.Product, error) {
	var products []models.Product

	if err := r.db.WithContext(ctx).Preload("Category").Where("stock <= ?", threshold).Find(&products).Error; err != nil {
		return nil, err
	}

	var result []*entities.Product
	for _, product := range products {
		result = append(result, r.modelToEntity(&product))
	}

	return result, nil
}

func (r *productRepository) modelToEntity(productModel *models.Product) *entities.Product {
	product := &entities.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
		Stock:       productModel.Stock,
		Image:       productModel.Image,
		CategoryID:  productModel.CategoryID,
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt,
	}

	if productModel.Category.ID != uuid.Nil {
		product.Category = &entities.Category{
			ID:          productModel.Category.ID,
			Name:        productModel.Category.Name,
			Description: productModel.Category.Description,
			Image:       productModel.Category.Image,
			CreatedAt:   productModel.Category.CreatedAt,
			UpdatedAt:   productModel.Category.UpdatedAt,
		}
	}

	for _, img := range productModel.Images {
		product.Images = append(product.Images, entities.ProductImage{
			ID:        img.ID,
			ProductID: img.ProductID,
			ImageURL:  img.ImageURL,
			CreatedAt: img.CreatedAt,
			UpdatedAt: img.UpdatedAt,
		})
	}

	return product
}
