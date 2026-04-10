package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type CartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity"`
}

type CartController struct {
	service services.CartServiceInterface
}

func NewCartController(service services.CartServiceInterface) *CartController {
	return &CartController{service: service}
}

func (ctrl *CartController) GetCart(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	carts, err := ctrl.service.GetCart(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", carts)
}

func (ctrl *CartController) AddItem(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	var req CartRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	cart, err := ctrl.service.AddItem(uint(userIDUint64), req.ProductID, req.Quantity)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Item Added to Cart", cart)
}

func (ctrl *CartController) UpdateCart(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	var req UpdateCartRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	cart, err := ctrl.service.UpdateCart(uint(userIDUint64), uint(id), req.Quantity)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Cart Updated", cart)
}

func (ctrl *CartController) DeleteCart(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.DeleteCart(uint(userIDUint64), uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Cart Item Deleted", nil)
}

func (ctrl *CartController) ClearCart(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	err = ctrl.service.ClearCart(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Cart Cleared", nil)
}
