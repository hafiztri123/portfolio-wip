package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafiztri123/internal/models"
	"github.com/hafiztri123/internal/repository"
)




type PortfolioHandler struct {
	repository *repository.PortfolioRepository
}

func NewPortfolioHandler(repository *repository.PortfolioRepository) *PortfolioHandler {
	return &PortfolioHandler{
		repository: repository,
	}
}

func (r *PortfolioHandler) CreatePortfolio(c *gin.Context) {
	userID := c.GetUint("user_id")
	var portfolio models.Portfolio
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio.UserID = uint(userID)
	if err := r.repository.Create(&portfolio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Portfolio created successfully"})
}

func (r *PortfolioHandler) GetAllPortfolios(c *gin.Context) {
	portfolios, err := r.repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, portfolios)
}

