package repositories

import (
	"context"
	"time"

	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"gorm.io/gorm"
)

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) repositories.StatsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) GetSalesStats(ctx context.Context) (*entities.SalesStats, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	thisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	thisYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())

	stats := &entities.SalesStats{}

	// Total sales และ orders (ทั้งหมด)
	var totalSales float64
	var totalOrders int64
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ?", "cancelled").
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&totalSales).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ?", "cancelled").
		Count(&totalOrders).Error; err != nil {
		return nil, err
	}
	stats.TotalSales = totalSales
	stats.TotalOrders = int(totalOrders)

	// Today's sales และ orders
	var todaySales float64
	var todayOrders int64
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", today).
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&todaySales).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", today).
		Count(&todayOrders).Error; err != nil {
		return nil, err
	}
	stats.TodaySales = todaySales
	stats.TodayOrders = int(todayOrders)

	// Monthly sales และ orders
	var monthlySales float64
	var monthlyOrders int64
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", thisMonth).
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&monthlySales).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", thisMonth).
		Count(&monthlyOrders).Error; err != nil {
		return nil, err
	}
	stats.MonthlySales = monthlySales
	stats.MonthlyOrders = int(monthlyOrders)

	// Yearly sales และ orders
	var yearlySales float64
	var yearlyOrders int64
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", thisYear).
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&yearlySales).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status != ? AND created_at >= ?", "cancelled", thisYear).
		Count(&yearlyOrders).Error; err != nil {
		return nil, err
	}
	stats.YearlySales = yearlySales
	stats.YearlyOrders = int(yearlyOrders)

	return stats, nil
}

func (r *statsRepository) GetProductStats(ctx context.Context) (*entities.ProductStats, error) {
	stats := &entities.ProductStats{}

	// Total products
	var totalProducts int64
	if err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&totalProducts).Error; err != nil {
		return nil, err
	}
	stats.TotalProducts = int(totalProducts)

	// Low stock products (สต็อกต่ำกว่า 10)
	var lowStockProducts int64
	if err := r.db.WithContext(ctx).Model(&models.Product{}).
		Where("stock > 0 AND stock <= 10").
		Count(&lowStockProducts).Error; err != nil {
		return nil, err
	}
	stats.LowStockProducts = int(lowStockProducts)

	// Out of stock products
	var outOfStockProducts int64
	if err := r.db.WithContext(ctx).Model(&models.Product{}).
		Where("stock = 0").
		Count(&outOfStockProducts).Error; err != nil {
		return nil, err
	}
	stats.OutOfStockProducts = int(outOfStockProducts)

	// Total categories
	var totalCategories int64
	if err := r.db.WithContext(ctx).Model(&models.Category{}).Count(&totalCategories).Error; err != nil {
		return nil, err
	}
	stats.TotalCategories = int(totalCategories)

	return stats, nil
}

func (r *statsRepository) GetUserStats(ctx context.Context) (*entities.UserStats, error) {
	stats := &entities.UserStats{}

	// Total users
	var totalUsers int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, err
	}
	stats.TotalUsers = int(totalUsers)

	// Active users
	var activeUsers int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).
		Where("active = ?", true).
		Count(&activeUsers).Error; err != nil {
		return nil, err
	}
	stats.ActiveUsers = int(activeUsers)

	// Inactive users
	var inactiveUsers int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).
		Where("active = ?", false).
		Count(&inactiveUsers).Error; err != nil {
		return nil, err
	}
	stats.InactiveUsers = int(inactiveUsers)

	// New users (ใน 30 วันที่ผ่านมา)
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	var newUsers int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).
		Where("created_at >= ?", thirtyDaysAgo).
		Count(&newUsers).Error; err != nil {
		return nil, err
	}
	stats.NewUsers = int(newUsers)

	return stats, nil
}
