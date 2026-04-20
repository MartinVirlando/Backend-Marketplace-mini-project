package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo,
	authCtrl *controllers.AuthController,
	productCtrl *controllers.ProductController,
	categoryCtrl *controllers.CategoryController,
	cartCtrl *controllers.CartController,
	orderCtrl *controllers.OrderController,
	reviewCtrl *controllers.ReviewController,
	messageCtrl *controllers.MessageController,
	notificationCtrl *controllers.NotificationController,
	adminCtrl *controllers.AdminController,
	uploadCtrl *controllers.UploadController,
) {
	e.Static("/uploads", "uploads")
	api := e.Group("/api")
	auth := api.Group("/auth")

	//Public
	auth.POST("/register", authCtrl.Register)
	auth.POST("/login", authCtrl.Login)

	api.GET("/products", productCtrl.GetAll)
	api.GET("/products/:id", productCtrl.GetByID)
	api.GET("/categories", categoryCtrl.GetAll)
	api.GET("/categories/:id", categoryCtrl.GetByID)
	api.GET("/products/:productId/reviews", reviewCtrl.GetReviews)

	// Protected
	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware())

	protected.GET("/auth/me", authCtrl.GetMe)
	protected.PUT("/auth/me", authCtrl.UpdateProfile)

	protected.POST("/cart", cartCtrl.AddItem)
	protected.GET("/cart", cartCtrl.GetCart)
	protected.PUT("/cart/:id", cartCtrl.UpdateCart)
	protected.DELETE("/cart/:id", cartCtrl.DeleteCart)
	protected.DELETE("/cart", cartCtrl.ClearCart)

	protected.POST("/orders", orderCtrl.CreateOrder)
	protected.GET("/orders", orderCtrl.GetOrders)
	protected.GET("/orders/:id", orderCtrl.GetOrderByID)
	protected.DELETE("/orders/:id", orderCtrl.CancelOrder)

	protected.POST("/products/:productId/reviews", reviewCtrl.CreateReview)
	protected.DELETE("/reviews/:id", reviewCtrl.DeleteReview)

	protected.GET("/messages", messageCtrl.GetConversations)
	protected.GET("/messages/:userId", messageCtrl.GetMessages)
	protected.POST("/messages", messageCtrl.SendMessage)
	protected.PUT("/messages/:userId/read", messageCtrl.MarkAsRead)

	protected.GET("/notifications", notificationCtrl.GetNotifications)
	protected.PUT("/notifications/:id/read", notificationCtrl.MarkAsRead)
	protected.PUT("/notifications/read-all", notificationCtrl.MarkAllAsRead)
	protected.DELETE("/notifications/:id", notificationCtrl.Delete)

	protected.POST("/upload", uploadCtrl.UploadImage)


	//Seller Only
	seller := api.Group("")
	seller.Use(middleware.JWTMiddleware(), middleware.SellerOnly())
	seller.GET("/seller/products", productCtrl.GetBySeller)

	seller.POST("/products", productCtrl.Create)
	seller.PUT("/products/:id", productCtrl.Update)
	seller.PUT("/products/:id/status", productCtrl.UpdateStatus)
	seller.DELETE("/products/:id", productCtrl.Delete)

	//Admin Only
	admin := api.Group("/admin")
	admin.Use(middleware.JWTMiddleware(), middleware.AdminOnly())

	admin.GET("/dashboard", adminCtrl.GetDashboardStats)
	admin.GET("/products/pending", adminCtrl.GetPendingProducts)
	admin.PUT("/products/:id/approve", adminCtrl.ApproveProduct)
	admin.PUT("/products/:id/reject", adminCtrl.RejectProduct)
	admin.PUT("/products/approve-all", adminCtrl.ApproveAllProducts)
	admin.GET("/users", adminCtrl.GetUsers)
	admin.DELETE("/users/:id", adminCtrl.DeleteUser)
	admin.POST("/categories", categoryCtrl.Create)
	admin.PUT("/categories/:id", categoryCtrl.Update)
	admin.DELETE("/categories/:id", categoryCtrl.Delete)

}
