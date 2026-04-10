package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type CategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CategoryController struct {
	service services.CategoryServiceInterface
}

func NewCategoryController(service services.CategoryServiceInterface) *CategoryController {
	return &CategoryController{service: service}
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	categories, err := ctrl.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success", categories)
}

func (ctrl *CategoryController) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	category, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Category not found")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", category)
}

func (ctrl *CategoryController) Create(c echo.Context) error {
	var req CategoryRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	category, err := ctrl.service.Create(req.Name, req.Icon)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Category Created", category)
}

func (ctrl *CategoryController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	var req CategoryRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	category, err := ctrl.service.Update(uint(id), req.Name, req.Icon)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Category Updated", category)
}

func (ctrl *CategoryController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.Delete(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Category Deleted", nil)
}
