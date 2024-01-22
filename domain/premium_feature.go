package domain

import (
	"context"
)

type PremiumFeature struct {
	UserID         string `json:"user_id"`
	Verified       bool   `json:"verified"`
	UnlimitedSwipe bool   `json:"unlimited_swipe"`
}

type PremiumFeatureUsecase interface {
	Subscribe(ctx context.Context, data PremiumFeature) error
}
