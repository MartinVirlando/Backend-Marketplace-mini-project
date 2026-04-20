package controllers

import (
	"net/http"
	"strconv"

	"backend/services"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`
}

type AuthController struct {
	service services.AuthServiceInterface
}

func NewAuthController(service services.AuthServiceInterface) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	user, err := controller.service.Register(req.Name, req.Email, req.Password, req.Phone, req.Role)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//Service Register
	token, user, err := controller.service.Login(req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Register Success", map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

// Login handler
func (controller *AuthController) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	//Service Login
	token, user, err := controller.service.Login(req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login Successful", map[string]interface{}{
		"token": token,
		"user":  user,
	})

}

// GetMe Handler
func (controller *AuthController) GetMe(c echo.Context) error {
	claims := c.Get("user").(*utils.JwtCustomClaims)

	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID format")
	}
	userID := uint(userIDUint64)

	user, err := controller.service.GetMe(userID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "User Not Found")
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success fetch profile", user)

}

func (controller *AuthController) UpdateProfile(c echo.Context) error {
	var req UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Request")
	}

	claims := c.Get("user").(*utils.JwtCustomClaims)

	userIDUint64, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid User ID Format")
	}
	userID := uint(userIDUint64)

	//Service Update Profile
	updatedUser, err := controller.service.UpdateProfile(userID, req.Name, req.Phone, req.Avatar)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Profile Updated", updatedUser)
}
