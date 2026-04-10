package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type NotificationController struct {
	service services.NotificationServiceInterface
}

func NewNotificationController(service services.NotificationServiceInterface) *NotificationController {
	return &NotificationController{service: service}
}

func (ctrl *NotificationController) GetNotifications(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	notifications, err := ctrl.service.GetNotifications(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", notifications)
}

func (ctrl *NotificationController) MarkAsRead(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = ctrl.service.MarkAsRead(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Notification marked as read", nil)
}

func (ctrl *NotificationController) MarkAllAsRead(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	err = ctrl.service.MarkAllAsRead(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "All notifications marked as read", nil)
}
