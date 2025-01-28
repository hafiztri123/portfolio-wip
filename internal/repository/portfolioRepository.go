package repository

import (
	"github.com/hafiztri123/internal/models"
	"gorm.io/gorm"
)

type PortfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) *PortfolioRepository {
	return &PortfolioRepository{
		db: db,
	}
}

func (r *PortfolioRepository) Create(portfolio *models.Portfolio) error {
	return r.db.Create(portfolio).Error
}

func (r *PortfolioRepository) GetAll() ([]models.Portfolio, error) {
	var portfolios []models.Portfolio
	err := r.db.Find(&portfolios).Error
	return portfolios, err
}

