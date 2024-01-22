package repository

import (
	"context"

	"gorm.io/gorm"

	domain "github.com/iqbalrestu07/datting-apps-api/domain"
)

type userInterestRepository struct {
	Conn *gorm.DB
}

// NewUserInterestRepository will create an object that represent the invoice.Repository interface
func NewUserInterestRepository(conn *gorm.DB) domain.UserInterestRepository {
	return &userInterestRepository{conn}
}

func (r *userInterestRepository) FindAll(ctx context.Context, filter domain.UserInterest) (userInterests []domain.UserInterest, err error) {
	// Add filters to the query
	db := r.Conn.WithContext(ctx)

	if filter.UserID != "" {
		db = db.Where("user_id = ?", filter.UserID)
	}

	err = db.Find(&userInterests).Error
	return userInterests, err
}

func (r *userInterestRepository) Create(ctx context.Context, userInteres *domain.UserInterest) (err error) {
	err = r.Conn.WithContext(ctx).Create(userInteres).Error
	if err != nil {
		return
	}

	return
}
