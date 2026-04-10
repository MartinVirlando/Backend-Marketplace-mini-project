package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type MessageRequest struct {
	ReceiverID uint   `json:"receiver_id"`
	Message    string `json:"message"`
	ProductID  *uint  `json:"product_id"`
}

type MessageController struct {
	service services.MessageServiceInterface
}

func NewMessageController(service services.MessageServiceInterface) *MessageController {
	return &MessageController{service: service}
}

func (ctrl *MessageController) GetConversations(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	conversations, err := ctrl.service.GetConversations(uint(userIDUint64))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", conversations)
}

func (ctrl *MessageController) GetMessages(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	otherUserID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid User ID")
	}

	messages, err := ctrl.service.GetMessages(uint(userIDUint64), uint(otherUserID))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success", messages)
}

func (ctrl *MessageController) SendMessage(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	var req MessageRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	message, err := ctrl.service.SendMessage(uint(userIDUint64), req.ReceiverID, req.ProductID, req.Message)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Message Sent", message)
}

func (ctrl *MessageController) MarkAsRead(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)
	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID")
	}

	senderID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid User ID")
	}

	err = ctrl.service.MarkAsRead(uint(userIDUint64), uint(senderID))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Messages marked as read", nil)
}