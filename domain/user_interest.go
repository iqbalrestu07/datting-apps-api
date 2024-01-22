package domain

import (
	"context"
	"time"
)

// UserInterest represents a user's interest.
type UserInterest struct {
	Model
	UserID     string    `gorm:"user_id" json:"user_id"`
	InterestID string    `gorm:"interest_id" json:"interest_id"`
	Interest   Interest  `gorm:"foreignKey:InterestID" json:"interest"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

// UserInterestRepository represent the UserInterest's repository contract
type UserInterestRepository interface {
	FindAll(ctx context.Context, req UserInterest) (res []UserInterest, err error)
	Create(ctx context.Context, ui *UserInterest) error
}
