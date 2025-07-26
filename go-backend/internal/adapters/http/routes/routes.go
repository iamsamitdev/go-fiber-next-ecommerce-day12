package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/http/middleware"
)

type Routes struct {
	authHandler     *handlers.AuthHandler
	userHandler     *handlers.UserHandler
	categoryHandler *handlers.CategoryHandler
	productHandler  *handlers.ProductHandler
	cartHandler     *handlers.CartHandler
	orderHandler    *handlers.OrderHandler
	paymentHandler  *handlers.PaymentHandler
	statsHandler    *handlers.StatsHandler
	authMW          *middleware.AuthMiddleware
}

func NewRoutes(
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler,
	statsHandler *handlers.StatsHandler,
	authMW *middleware.AuthMiddleware,
) *Routes {
	return &Routes{
		authHandler:     authHandler,
		userHandler:     userHandler,
		categoryHandler: categoryHandler,
		productHandler:  productHandler,
		cartHandler:     cartHandler,
		orderHandler:    orderHandler,
		paymentHandler:  paymentHandler,
		statsHandler:    statsHandler,
		authMW:          authMW,
	}
}

func (r *Routes) SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Fiber Ecommerce API is running",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", r.authHandler.Register)
	auth.Post("/login", r.authHandler.Login)
	auth.Post("/refresh", r.authHandler.RefreshToken)
	auth.Post("/forgot-password", r.authHandler.ForgotPassword)
	auth.Post("/reset-password", r.authHandler.ResetPassword)

	// Protected auth routes
	authProtected := auth.Group("", r.authMW.AuthRequired())
	authProtected.Post("/logout", r.authHandler.Logout)
	authProtected.Post("/change-password", r.authHandler.ChangePassword)

	// Admin only auth routes
	authAdmin := auth.Group("", r.authMW.AuthRequired(), r.authMW.AdminRequired())
	authAdmin.Post("/admin/register", r.authHandler.AdminRegister)

	// User routes (admin only)
	users := api.Group("/users", r.authMW.AuthRequired(), r.authMW.AdminRequired())
	users.Get("/", r.userHandler.GetUsers)
	users.Get("/:id", r.userHandler.GetUserByID)
	users.Put("/:id", r.userHandler.UpdateUser)
	users.Delete("/:id", r.userHandler.DeleteUser)

	// Categories (admin only for CUD, public for read)
	categories := api.Group("/categories")
	categories.Get("/", r.categoryHandler.GetCategories)
	categories.Get("/:id", r.categoryHandler.GetCategoryByID)
	categoriesAdmin := categories.Group("", r.authMW.AuthRequired(), r.authMW.AdminRequired())
	categoriesAdmin.Post("/", r.categoryHandler.CreateCategory)
	categoriesAdmin.Put("/:id", r.categoryHandler.UpdateCategory)
	categoriesAdmin.Delete("/:id", r.categoryHandler.DeleteCategory)

	// Products (admin only for CUD, public for read)
	products := api.Group("/products")
	products.Get("/", r.productHandler.GetProducts)
	products.Get("/:id", r.productHandler.GetProductByID)
	products.Get("/category/:categoryId", r.productHandler.GetProductsByCategory)
	products.Get("/search", r.productHandler.SearchProducts)
	productsAdmin := products.Group("", r.authMW.AuthRequired(), r.authMW.AdminRequired())
	productsAdmin.Post("/", r.productHandler.CreateProduct)
	productsAdmin.Put("/:id", r.productHandler.UpdateProduct)
	productsAdmin.Delete("/:id", r.productHandler.DeleteProduct)

	// Cart (user only)
	cart := api.Group("/cart", r.authMW.AuthRequired())
	cart.Get("/", r.cartHandler.GetCart)
	cart.Post("/", r.cartHandler.AddToCart)
	cart.Put("/:itemId", r.cartHandler.UpdateCartItem)
	cart.Delete("/:itemId", r.cartHandler.RemoveFromCart)
	cart.Delete("/", r.cartHandler.ClearCart)

	// Orders (user for own orders, admin for all)
	orders := api.Group("/orders", r.authMW.AuthRequired())
	orders.Post("/", r.orderHandler.CreateOrder)
	orders.Get("/", r.orderHandler.GetOrders)
	orders.Get("/:id", r.orderHandler.GetOrderByID)
	orders.Put("/:id/cancel", r.orderHandler.CancelOrder)
	ordersAdmin := orders.Group("/admin", r.authMW.AdminRequired())
	ordersAdmin.Get("/", r.orderHandler.GetAllOrders)
	ordersAdmin.Put("/:id/status", r.orderHandler.UpdateOrderStatus)

	// Payments (user only)
	payments := api.Group("/payments", r.authMW.AuthRequired())
	payments.Post("/", r.paymentHandler.CreatePayment)
	payments.Post("/:id/verify", r.paymentHandler.VerifyPayment)
	payments.Put("/:id/cancel", r.paymentHandler.CancelPayment)

	// Stats (admin only)
	stats := api.Group("/stats", r.authMW.AuthRequired(), r.authMW.AdminRequired())
	stats.Get("/sales", r.statsHandler.GetSalesStats)
	stats.Get("/products", r.statsHandler.GetProductStats)
	stats.Get("/users", r.statsHandler.GetUserStats)
}
