package domain

import (
	"time"
)

// Interest represents a user's interest.
type Interest struct {
	Model
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
