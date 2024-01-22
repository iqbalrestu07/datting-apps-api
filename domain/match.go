package domain

import (
	"context"

	"github.com/iqbalrestu07/datting-apps-api/request"
	uuid "github.com/satori/go.uuid"
)

type Match struct {
	Model
	UserID       uuid.UUID `json:"user_id"`
	TargetUserID uuid.UUID `json:"target_user_id"`
	IsLike       bool      `json:"is_like"`
	IsMatch      bool      `json:"is_match"`
	User         *User     `gorm:"foreignKey:user_id"`
	TargetUser   *User     `gorm:"foreignKey:target_user_id"`
}

type MatchUsecase interface {
	FindAll(ctx context.Context, filter request.MatchRequest) (matchs []Match, err error)
	Match(ctx context.Context, data *Match) error
}

type MatchRepository interface {
	FindAll(ctx context.Context, filter request.MatchRequest) (matches []Match, err error)
	CheckForLike(ctx context.Context, filter request.MatchRequest) (Match, error)
	Create(ctx context.Context, data *Match) error
	Update(ctx context.Context, data *Match) error
}
