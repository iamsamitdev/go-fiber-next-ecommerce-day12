package config

import (
	"log"

	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
	"gorm.io/gorm"
)

// SeedDatabase สำหรับ seed ข้อมูลเริ่มต้น
func SeedDatabase(db *gorm.DB, config *Config) error {
	log.Println("🌱 Starting database seeding...")

	// Seed roles ก่อน
	if err := seedRoles(db); err != nil {
		return err
	}

	// Seed admin user
	if err := seedAdminUser(db, config); err != nil {
		return err
	}

	// Seed categories
	if err := seedCategories(db); err != nil {
		return err
	}

	// Seed products
	if err := seedProducts(db); err != nil {
		return err
	}

	log.Println("✅ Database seeding completed successfully!")
	return nil
}

// seedAdminUser สร้าง admin user เริ่มต้น
func seedAdminUser(db *gorm.DB, config *Config) error {
	// ตรวจสอบว่ามี admin credentials หรือไม่
	if config.AdminEmail == "" {
		log.Println("⚠️  ADMIN_EMAIL not set, skipping admin user seeding")
		log.Println("💡 To create admin user, set ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_FIRST_NAME, ADMIN_LAST_NAME in .env")
		return nil
	}

	// ตรวจสอบว่ามี admin user อยู่แล้วหรือไม่
	var existingUser models.User
	if err := db.Where("email = ?", config.AdminEmail).First(&existingUser).Error; err == nil {
		log.Printf("ℹ️  Admin user already exists: %s", config.AdminEmail)
		return nil
	}

	// หา admin role
	var adminRole models.Role
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		log.Printf("❌ Admin role not found: %v", err)
		return err
	}

	// ตรวจสอบรหัสผ่าน
	if config.AdminPassword == "" {
		log.Println("⚠️  ADMIN_PASSWORD not set, skipping admin user creation")
		return nil
	}

	if err := utils.ValidatePassword(config.AdminPassword); err != nil {
		log.Printf("⚠️  ADMIN_PASSWORD validation failed: %v", err)
		log.Println("💡 Admin password must contain at least 8 characters with uppercase, lowercase, number, and special character")
		return nil
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(config.AdminPassword)
	if err != nil {
		log.Printf("❌ Error hashing admin password: %v", err)
		return err
	}

	// สร้าง admin user
	adminUser := &models.User{
		Email:     config.AdminEmail,
		Password:  hashedPassword,
		FirstName: config.AdminFirstName,
		LastName:  config.AdminLastName,
		RoleID:    adminRole.ID,
		Active:    true,
	}

	if err := db.Create(adminUser).Error; err != nil {
		log.Printf("❌ Error creating admin user: %v", err)
		return err
	}

	log.Printf("✅ Admin user created successfully: %s", config.AdminEmail)
	log.Printf("👤 Name: %s %s", config.AdminFirstName, config.AdminLastName)
	log.Println("🔐 Password meets security requirements (uppercase, lowercase, number, special character)")

	return nil
}

// seedRoles สร้าง roles เริ่มต้น
func seedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{
			Name:        "admin",
			Description: "Administrator with full access",
		},
		{
			Name:        "user",
			Description: "Regular user with limited access",
		},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// สร้าง role ใหม่
				if err := db.Create(&role).Error; err != nil {
					log.Printf("❌ Error creating role %s: %v", role.Name, err)
					return err
				}
				log.Printf("✅ Role created: %s", role.Name)
			} else {
				log.Printf("❌ Error checking role %s: %v", role.Name, err)
				return err
			}
		}
	}

	return nil
}

// seedCategories สร้างหมวดหมู่สินค้าเริ่มต้น
func seedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{
			Name:        "Electronics",
			Description: "อุปกรณ์อิเล็กทรอนิกส์และเทคโนโลยี",
			Image:       "https://images.unsplash.com/photo-1498049794561-7780e7231661?w=400",
		},
		{
			Name:        "Fashion",
			Description: "เสื้อผ้าและแฟชั่น",
			Image:       "https://images.unsplash.com/photo-1445205170230-053b83016050?w=400",
		},
		{
			Name:        "Home & Garden",
			Description: "ของใช้ในบ้านและสวน",
			Image:       "https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=400",
		},
		{
			Name:        "Sports & Outdoors",
			Description: "อุปกรณ์กีฬาและกิจกรรมกลางแจ้ง",
			Image:       "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400",
		},
		{
			Name:        "Books & Media",
			Description: "หนังสือและสื่อต่างๆ",
			Image:       "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400",
		},
		{
			Name:        "Health & Beauty",
			Description: "ผลิตภัณฑ์สุขภาพและความงาม",
			Image:       "https://images.unsplash.com/photo-1556228720-195a672e8a03?w=400",
		},
		{
			Name:        "Toys & Games",
			Description: "ของเล่นและเกมส์",
			Image:       "https://images.unsplash.com/photo-1558060370-d644479cb6f7?w=400",
		},
		{
			Name:        "Automotive",
			Description: "อุปกรณ์และอะไหล่รถยนต์",
			Image:       "https://images.unsplash.com/photo-1492144534655-ae79c964c9d7?w=400",
		},
		{
			Name:        "Food & Beverages",
			Description: "อาหารและเครื่องดื่ม",
			Image:       "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=400",
		},
		{
			Name:        "Office Supplies",
			Description: "อุปกรณ์สำนักงาน",
			Image:       "https://images.unsplash.com/photo-1497032628192-86f99bcd76bc?w=400",
		},
	}

	for _, category := range categories {
		var existingCategory models.Category
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&category).Error; err != nil {
					log.Printf("❌ Error creating category %s: %v", category.Name, err)
					return err
				}
				log.Printf("✅ Category created: %s", category.Name)
			} else {
				log.Printf("❌ Error checking category %s: %v", category.Name, err)
				return err
			}
		}
	}

	return nil
}

// seedProducts สร้างสินค้าเริ่มต้น
func seedProducts(db *gorm.DB) error {
	// ดึงหมวดหมู่ทั้งหมด
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Printf("❌ Error fetching categories: %v", err)
		return err
	}

	if len(categories) == 0 {
		log.Println("⚠️  No categories found, skipping product seeding")
		return nil
	}

	products := []models.Product{
		// Electronics
		{
			Name:        "iPhone 15 Pro",
			Description: "สมาร์ทโฟนรุ่นล่าสุดจาก Apple พร้อมชิป A17 Pro",
			Price:       39900,
			Stock:       50,
			CategoryID:  categories[0].ID, // Electronics
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1592750475338-74b7b21085ab?w=400"},
			},
		},
		{
			Name:        "MacBook Air M2",
			Description: "แล็ปท็อปที่บางและเบาพร้อมชิป M2",
			Price:       42900,
			Stock:       30,
			CategoryID:  categories[0].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1541807084-5c52b6b3adef?w=400"},
			},
		},
		// Fashion
		{
			Name:        "เสื้อเชิ้ตผ้าคอตตอน",
			Description: "เสื้อเชิ้ตผ้าคอตตอน 100% สีขาว คลาสสิค",
			Price:       1290,
			Stock:       100,
			CategoryID:  categories[1].ID, // Fashion
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1602810318383-e386cc2a3ccf?w=400"},
			},
		},
		{
			Name:        "กางเกงยีนส์ Slim Fit",
			Description: "กางเกงยีนส์ทรง Slim Fit สีน้ำเงินเข้ม",
			Price:       1890,
			Stock:       75,
			CategoryID:  categories[1].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1542272604-787c3835535d?w=400"},
			},
		},
		// Home & Garden
		{
			Name:        "โซฟาผ้า 3 ที่นั่ง",
			Description: "โซฟาผ้าสีเทา 3 ที่นั่ง สไตล์โมเดิร์น",
			Price:       15900,
			Stock:       20,
			CategoryID:  categories[2].ID, // Home & Garden
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=400"},
			},
		},
		{
			Name:        "ชุดเครื่องนอน Cotton",
			Description: "ชุดเครื่องนอนผ้าคอตตอน 100% ขนาด 6 ฟุต",
			Price:       2490,
			Stock:       60,
			CategoryID:  categories[2].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=400"},
			},
		},
		// Sports & Outdoors
		{
			Name:        "รองเท้าวิ่ง Nike",
			Description: "รองเท้าวิ่งเพื่อสุขภาพ เบาสบาย",
			Price:       3290,
			Stock:       80,
			CategoryID:  categories[3].ID, // Sports & Outdoors
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=400"},
			},
		},
		{
			Name:        "ดัมเบลปรับน้ำหนักได้",
			Description: "ดัมเบลปรับน้ำหนักได้ 5-25 กก. คู่ละ",
			Price:       4590,
			Stock:       25,
			CategoryID:  categories[3].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400"},
			},
		},
		// Books & Media
		{
			Name:        "หนังสือ Clean Code",
			Description: "หนังสือสอนการเขียนโค้ดที่สะอาด",
			Price:       890,
			Stock:       40,
			CategoryID:  categories[4].ID, // Books & Media
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400"},
			},
		},
		{
			Name:        "หนังสือ Design Patterns",
			Description: "หนังสือเรียนรู้ Design Patterns",
			Price:       1290,
			Stock:       35,
			CategoryID:  categories[4].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=400"},
			},
		},
		// Health & Beauty
		{
			Name:        "ครีมบำรุงหน้า Vitamin C",
			Description: "ครีมบำรุงหน้าสูตร Vitamin C ลดจุดด่างดำ",
			Price:       1590,
			Stock:       90,
			CategoryID:  categories[5].ID, // Health & Beauty
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1556228720-195a672e8a03?w=400"},
			},
		},
		{
			Name:        "แชมพูสมุนไพร",
			Description: "แชมพูสมุนไพรสำหรับผมแห้งเสีย",
			Price:       390,
			Stock:       120,
			CategoryID:  categories[5].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1571781926291-c477ebfd024b?w=400"},
			},
		},
		// Toys & Games
		{
			Name:        "ตัวต่อเลโก้ Creator",
			Description: "ชุดตัวต่อเลโก้ Creator Expert 2000 ชิ้น",
			Price:       3990,
			Stock:       30,
			CategoryID:  categories[6].ID, // Toys & Games
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1558060370-d644479cb6f7?w=400"},
			},
		},
		{
			Name:        "บอร์ดเกม Monopoly",
			Description: "เกมโมโนโพลี่ฉบับภาษาไทย",
			Price:       1290,
			Stock:       45,
			CategoryID:  categories[6].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1606092195730-5d7b9af1efc5?w=400"},
			},
		},
		// Automotive
		{
			Name:        "น้ำมันเครื่อง Mobil 1",
			Description: "น้ำมันเครื่องสังเคราะห์ 100% 5W-30",
			Price:       890,
			Stock:       100,
			CategoryID:  categories[7].ID, // Automotive
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1492144534655-ae79c964c9d7?w=400"},
			},
		},
		{
			Name:        "ยางรถยนต์ Michelin",
			Description: "ยางรถยนต์ Michelin ขนาด 195/65R15",
			Price:       3290,
			Stock:       40,
			CategoryID:  categories[7].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400"},
			},
		},
		// Food & Beverages
		{
			Name:        "กาแฟอราบิก้า 100%",
			Description: "เมล็ดกาแฟอราบิก้า 100% คั่วกลาง 250g",
			Price:       590,
			Stock:       80,
			CategoryID:  categories[8].ID, // Food & Beverages
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=400"},
			},
		},
		{
			Name:        "ชาเขียวญี่ปุ่น",
			Description: "ชาเขียวญี่ปุ่นแท้ 100g",
			Price:       890,
			Stock:       60,
			CategoryID:  categories[8].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1594631661960-e0b9d10d6d81?w=400"},
			},
		},
		// Office Supplies
		{
			Name:        "เก้าอี้สำนักงานเออร์โกโนมิก",
			Description: "เก้าอี้สำนักงานปรับระดับได้ รองรับหลัง",
			Price:       7990,
			Stock:       15,
			CategoryID:  categories[9].ID, // Office Supplies
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1497032628192-86f99bcd76bc?w=400"},
			},
		},
		{
			Name:        "โต๊ะทำงานไม้โอ๊ค",
			Description: "โต๊ะทำงานไม้โอ๊คขนาด 120x60 ซม.",
			Price:       5490,
			Stock:       20,
			CategoryID:  categories[9].ID,
			Images: []models.ProductImage{
				{ImageURL: "https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=400"},
			},
		},
	}

	for _, product := range products {
		var existingProduct models.Product
		if err := db.Where("name = ?", product.Name).First(&existingProduct).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// สร้าง product พร้อม images
				if err := db.Create(&product).Error; err != nil {
					log.Printf("❌ Error creating product %s: %v", product.Name, err)
					return err
				}
				log.Printf("✅ Product created: %s (฿%.2f)", product.Name, product.Price)
			} else {
				log.Printf("❌ Error checking product %s: %v", product.Name, err)
				return err
			}
		}
	}

	return nil
}
