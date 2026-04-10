package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	service services.AdminServiceInterface
}

func NewAdminController(service services.AdminServiceInterface) *AdminController {
	return &AdminController{service: service}
}

func (ctrl *AdminController) GetDashboardStats(c echo.Context) error {
	stats, err := ctrl.service.GetDashboardStats()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", stats)
}

func (ctrl *AdminController) GetPendingProducts(c echo.Context) error {
	products, err := ctrl.service.GetPendingProducts()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", products)
}

func (ctrl *AdminController) ApproveProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.ApproveProduct(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Product Approved", nil)
}

func (ctrl *AdminController) RejectProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.RejectProduct(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Product Rejected", nil)
}

func (ctrl *AdminController) ApproveAllProducts(c echo.Context) error {
	err := ctrl.service.ApproveAllProducts()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "All Products Approved", nil)
}

func (ctrl *AdminController) GetUsers(c echo.Context) error {
	users, err := ctrl.service.GetUsers()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", users)
}

func (ctrl *AdminController) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.DeleteUser(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "User Deleted", nil)
}
