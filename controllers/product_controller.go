package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Stock       int                `json:"stock"`
	CategoryID  uint               `json:"category_id"`
	Images      []string           `json:"images"`
}

type ProductController struct {
	service services.ProductServiceInterface
}

func NewProductController(service services.ProductServiceInterface) *ProductController {
	return &ProductController{service: service}
}

func (ctrl *ProductController) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	categoryID, _ := strconv.ParseUint(c.QueryParam("category_id"), 10, 64)

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 12
	}

	products, err := ctrl.service.GetAll(search, uint(categoryID), page, limit)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", products)
}

func (ctrl *ProductController) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	product, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Product not found")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", product)
}

func (ctrl *ProductController) Create(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	sellerIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	product, err := ctrl.service.Create(uint(sellerIDUint64), services.ProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Images:      req.Images,
	})
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Product Created", product)
}

func (ctrl *ProductController) Update(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	sellerIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	product, err := ctrl.service.Update(uint(id), uint(sellerIDUint64), services.ProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		Images:      req.Images,
	})
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Product Updated", product)
}

func (ctrl *ProductController) Delete(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	sellerIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.Delete(uint(id), uint(sellerIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Product Deleted", nil)
}

func (ctrl *ProductController) GetBySeller(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	sellerIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	products, err := ctrl.service.GetBySeller(uint(sellerIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", products)
}

func (ctrl *ProductController) UpdateStatus(c echo.Context) error {
    claims := c.Get("user").(*utils.JwtCustomClaims)
    sellerIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
    if err != nil {
        return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
    }

    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
    }

    var req struct {
        Status string `json:"status"`
    }
    if err := c.Bind(&req); err != nil {
        return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
    }

    // Validasi status yang boleh diset seller
    validStatus := map[string]bool{"hide": true, "ready": true, "sold": true}
    if !validStatus[req.Status] {
        return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid status")
    }

    product, err := ctrl.service.UpdateStatus(uint(id), uint(sellerIDUint64), req.Status)
    if err != nil {
        return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
    }

    return utils.SuccessResponse(c, http.StatusOK, "Product Status Updated", product)
}