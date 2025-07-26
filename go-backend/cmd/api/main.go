// @title Go E-commerce API
// @version 1.0.0
// @description API สำหรับระบบ E-commerce ที่พัฒนาด้วย Go Fiber
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4001
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT token สำหรับการยืนยันตัวตน ให้ใส่ token ในรูปแบบ: Bearer <token>
package main

import (
	"log"

	_ "github.com/iamsamitdev/fiber-ecommerce-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/http/middleware"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/http/routes"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/config"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup database connection
	db := config.SetupDatabase(cfg)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)

	categoryRepo := repositories.NewCategoryRepository(db)
	productRepo := repositories.NewProductRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	statsRepo := repositories.NewStatsRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, roleRepo)
	userService := services.NewUserService(userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo)
	paymentService := services.NewPaymentService(transactionRepo)
	statsService := services.NewStatsService(statsRepo)

	// Initialize middleware
	authMW := middleware.NewAuthMiddleware(cfg.JWTSecret)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(userService)

	// เพิ่ม handlers อื่นๆ
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)
	orderHandler := handlers.NewOrderHandler(orderService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	statsHandler := handlers.NewStatsHandler(statsService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Initialize routes
	routes := routes.NewRoutes(
		authHandler,
		userHandler,
		categoryHandler,
		productHandler,
		cartHandler,
		orderHandler,
		paymentHandler,
		statsHandler,
		authMW,
	)
	routes.SetupRoutes(app)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
