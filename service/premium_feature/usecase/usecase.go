package usecase

import (
	"context"

	"github.com/iqbalrestu07/datting-apps-api/domain"
)

type premiumFeatureUsecase struct {
	userRepo domain.UserRepository
}

func NewPremiumFeatureUsecase(ur domain.UserRepository) premiumFeatureUsecase {
	return premiumFeatureUsecase{
		userRepo: ur,
	}
}

func (u premiumFeatureUsecase) Subscribe(ctx context.Context, data domain.PremiumFeature) (err error) {
	user, err := u.userRepo.FindByID(ctx, data.UserID)
	if err != nil {
		return err
	}
	if data.UnlimitedSwipe {
		user.IsPremium = true
	}
	if data.Verified {
		user.IsVerified = true
	}

	err = u.userRepo.Update(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
