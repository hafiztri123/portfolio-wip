package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafiztri123/internal/models"
	"github.com/hafiztri123/internal/repository"
	"github.com/hafiztri123/internal/utils"
	"golang.org/x/crypto/bcrypt"
)


type AuthHandler struct {
	repository *repository.AuthRepository
}

func NewAuthHandler(repository *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{
		repository: repository,
	}
}

func(h *AuthHandler)  HandleRegister(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	//TODO: Implement register logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Email: req.Email,
		Password: string(hashedPassword),
	}

	if err := h.repository.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func (h *AuthHandler) HandleLogin(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	var user *models.User
	user, err := h.repository.GetUserByEmail(req.Email); 
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	

}