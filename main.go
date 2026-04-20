package main

import (
	"log"
	"os"

	"backend/config"
	"backend/controllers"
	"backend/repositories"
	"backend/routes"
	"backend/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Connect DB
	config.ConnectDB()

	//Migrate DB
	config.MigrateDB()

	//Init Repositories
	userRepo := repositories.NewUserRepository(config.DB)
	productRepo := repositories.NewProductRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	cartRepo := repositories.NewCartRepository(config.DB)
	orderRepo := repositories.NewOrderRepository(config.DB)
	reviewRepo := repositories.NewReviewRepository(config.DB)
	messageRepo := repositories.NewMessageRepository(config.DB)
	notificationRepo := repositories.NewNotificationRepository(config.DB)

	//Init Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo)
	reviewService := services.NewReviewService(reviewRepo, orderRepo)
	messageService := services.NewMessageService(messageRepo, notificationRepo)
	notificationService := services.NewNotificationService(notificationRepo)
	adminService := services.NewAdminService(productRepo, userRepo, orderRepo)

	//Init Controllers
	authCtrl := controllers.NewAuthController(authService)
	productCtrl := controllers.NewProductController(productService)
	categoryCtrl := controllers.NewCategoryController(categoryService)
	cartCtrl := controllers.NewCartController(cartService)
	orderCtrl := controllers.NewOrderController(orderService, cartService)
	reviewCtrl := controllers.NewReviewController(reviewService)
	messageCtrl := controllers.NewMessageController(messageService)
	notificationCtrl := controllers.NewNotificationController(notificationService)
	adminCtrl := controllers.NewAdminController(adminService)
	uploadCtrl := controllers.NewUploadController()


	//Setup Echo
	e := echo.New()

	//Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	//Setup Routes
	routes.SetupRoute(e,
		authCtrl,
		productCtrl,
		categoryCtrl,
		cartCtrl,
		orderCtrl,
		reviewCtrl,
		messageCtrl,
		notificationCtrl,
		adminCtrl,
		uploadCtrl,
	)

	//Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
