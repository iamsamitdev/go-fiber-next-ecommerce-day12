package config

import (
	"fmt"
	"log"
	"os"

	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// ตรวจสอบว่าต้องการ migrate หรือไม่
	if shouldRunMigration() {

		// runMigration จะทำการ migrate ถ้าเงื่อนไขเป็นจริง
		runMigration(db)

		// Seed database หลังจาก migration เสร็จ
		if err := SeedDatabase(db, config); err != nil {
			log.Printf("Database seeding failed: %v", err)
		}

	} else {
		// แสดง message ที่ชัดเจนขึ้นตามสาเหตุ
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Printf("Skipping database migration (AUTO_MIGRATE=false)")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Printf("Skipping database migration (production environment, set AUTO_MIGRATE=true to enable)")
		} else {
			log.Printf("Skipping database migration (set AUTO_MIGRATE=true to enable)")
		}

		// ลองสร้างข้อมูลตัวอย่าง แม้ว่าจะไม่ได้ migrate (กรณีที่ตารางมีอยู่แล้ว)
		if err := SeedDatabase(db, config); err != nil {
			log.Printf("Database seeding failed: %v", err)
		}
	}

	return db

}

// สร้างฟังก์ชัน ตรวจสอบว่าควร migrate หรือไม่
func shouldRunMigration() bool {
	// ถ้ากำหนด AUTO_MIGRATE=false ให้ไม่ migrate เลย (ทุก environment)
	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}
	// ถ้ากำหนด AUTO_MIGRATE=true ให้ migrate เลย (ทุก environment)
	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}
	// ถ้าไม่ได้กำหนด AUTO_MIGRATE ให้ใช้ default ตาม environment
	// Development - migrate อัตโนมัติ
	if os.Getenv("APP_ENV") == "development" {
		return true
	}
	// Production - ไม่ migrate อัตโนมัติ
	return false
}

// ฟังก์ชันสำหรับ migrate
func runMigration(db *gorm.DB) {
	log.Println("Starting database migration...")

	// Migrate all models
	err := db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}

// ฟังก์ชันสำหรับ migrate แบบ manual (สำหรับ CLI)
func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)

	log.Println("Running manual migration...")

	// Migrate all models
	err := db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	log.Println("Manual migration completed successfully")
	return nil
}
