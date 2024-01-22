package domain

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// Photo represents a user's photo.
type Photo struct {
	Model
	UserID uuid.UUID `gorm:"user_id" json:"user_id"`
	URL    string    `gorm:"url" json:"url"`
}

// UploadUsecase represent the Upload's usecases
type UploadUsecase interface {
	Create(ctx context.Context, m *Photo) (err error)
}

// UploadRepository represent the Upload's repository contract
type UploadRepository interface {
	Create(ctx context.Context, inv *Photo) error
}
