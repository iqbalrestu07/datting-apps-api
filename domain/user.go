package domain

import (
	"context"
	"time"

	"github.com/iqbalrestu07/datting-apps-api/request"
)

// User is representing the User data struct
type User struct {
	Model
	Name                string         `json:"name" gorm:"type:varchar(255);not null"`
	Email               string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password            string         `json:"-" gorm:"type:varchar(255);not null"`
	DateOfBirth         time.Time      `json:"dateOfBirth"`
	Bio                 string         `json:"bio" gorm:"type:text"`
	Gender              string         `json:"gender" gorm:"type:varchar(255)"`
	Photos              []Photo        `gorm:"foreignKey:UserID" json:"photos"`
	IsPremium           bool           `gorm:"is_premium" json:"is_premium"`
	IsVerified          bool           `gorm:"is_verified" json:"is_verified"`
	UserInterests       []UserInterest `gorm:"foreignKey:UserID" json:"user_interests"`
	SharedInterestCount int            `gorm:"-" json:"-"`
}

// UserUsecase represent the User's usecases
type UserUsecase interface {
	FindAll(ctx context.Context, filter request.UserRequest) ([]User, error)
	GetUserListSortedByInterest(ctx context.Context, userID string) ([]User, error)
	FindByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, m *User) (err error)
	Update(ctx context.Context, cust *User) error
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	FindAll(ctx context.Context, filter request.UserRequest) (users []User, err error)
	FindByID(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByMultipleID(ctx context.Context, ids []string) (users []User, err error)
	Create(ctx context.Context, inv *User) error
	Update(ctx context.Context, inv *User) error
}
