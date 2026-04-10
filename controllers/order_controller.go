package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	Items           []OrderItemRequest `json:"items"`
	ShippingAddress string             `json:"shipping_address"`
	City            string             `json:"city"`
	Province        string             `json:"province"`
	PostalCode      string             `json:"postal_code"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type OrderController struct {
	service     services.OrderServiceInterface
	cartService services.CartServiceInterface
}

func NewOrderController(service services.OrderServiceInterface, cartService services.CartServiceInterface) *OrderController {
	return &OrderController{
		service:     service,
		cartService: cartService,
	}
}

func (ctrl *OrderController) GetOrders(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	orders, err := ctrl.service.GetOrders(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", orders)
}

func (ctrl *OrderController) GetOrderByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	order, err := ctrl.service.GetOrderByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Order not found")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", order)
}

func (ctrl *OrderController) CreateOrder(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}
	userID := uint(userIDUint64)

	var req OrderRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	// Ambil cart items user
	cartItems, err := ctrl.cartService.GetCart(userID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if len(cartItems) == 0 {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Cart is empty")
	}

	// Buat order dari cart items
	order, err := ctrl.service.CreateOrder(
		userID,
		cartItems,
		req.ShippingAddress,
		req.City,
		req.Province,
		req.PostalCode,
	)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Clear cart setelah order berhasil
	ctrl.cartService.ClearCart(userID)

	return utils.SuccessResponse(c, http.StatusCreated, "Order Created", order)
}

func (ctrl *OrderController) CancelOrder(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.CancelOrder(uint(id), uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Order Cancelled", nil)
}
