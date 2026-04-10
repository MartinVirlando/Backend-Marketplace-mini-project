package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type ReviewRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type ReviewController struct {
	service services.ReviewServiceInterface
}

func NewReviewController(service services.ReviewServiceInterface) *ReviewController {
	return &ReviewController{service: service}
}

func (ctrl *ReviewController) GetReviews(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Product ID")
	}

	reviews, err := ctrl.service.GetReviews(uint(productID))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", reviews)
}

func (ctrl *ReviewController) CreateReview(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Product ID")
	}

	var req ReviewRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	review, err := ctrl.service.CreateReview(uint(userIDUint64), uint(productID), req.Rating, req.Comment)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Review Created", review)
}

func (ctrl *ReviewController) DeleteReview(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.DeleteReview(uint(id), uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Review Deleted", nil)
}
