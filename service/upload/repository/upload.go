package repository

import (
	"context"

	"gorm.io/gorm"

	domain "github.com/iqbalrestu07/datting-apps-api/domain"
)

type uploadRepository struct {
	Conn *gorm.DB
}

// NewUploadRepository will create an object that represent the uploadoice.Repository interface
func NewUploadRepository(conn *gorm.DB) domain.UploadRepository {
	return &uploadRepository{conn}
}

func (r *uploadRepository) Create(ctx context.Context, photo *domain.Photo) (err error) {
	err = r.Conn.WithContext(ctx).Create(photo).Error
	if err != nil {
		return
	}

	return
}
