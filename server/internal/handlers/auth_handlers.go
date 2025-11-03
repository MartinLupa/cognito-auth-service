package handlers

import (
	"net/http"

	"github.com/MartinLupa/go-cognito-auth/internal/models"
	"github.com/MartinLupa/go-cognito-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandlers(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Signup(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "signup successful", "user": user.Email})
}

func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	var params struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	c.ShouldBind(&params)
	err := h.authService.ConfirmEmail(params.Email, params.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email confirmed successfully"})
}

func (h *AuthHandler) ResendConfirmationCode(c *gin.Context) {
	var params struct {
		Email string `json:"email" binding:"required,email"`
	}

	c.ShouldBind(&params)
	err := h.authService.ResendConfirmationCode(params.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code resent successfully"})
}

func (h *AuthHandler) Signin(c *gin.Context) {
	var params struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	c.ShouldBind(&params)

	token, err := h.authService.Signin(params.Email, params.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signin successful", "token": token})
}

func (h *AuthHandler) Signout(c *gin.Context) {
	var params struct {
		AccessToken string `json:"access_token" binding:"required"`
	}

	c.ShouldBind(&params)

	err := h.authService.Signout(params.AccessToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signout successful"})
}

func (h *AuthHandler) ListUsers(c *gin.Context) {
	users, err := h.authService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Users listed successfully", "users": users})
}
