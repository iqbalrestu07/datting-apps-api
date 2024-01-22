package usecase

import (
	"context"
	"errors"

	"github.com/iqbalrestu07/datting-apps-api/common"
	domain "github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
)

type matchUsecase struct {
	matchRepo domain.MatchRepository
	userRepo  domain.UserRepository
}

// NewMatchUsecase will create new an matchUsecase object representation of domain.MatchUsecase interface
func NewMatchUsecase(mr domain.MatchRepository, ur domain.UserRepository) domain.MatchUsecase {
	return &matchUsecase{
		matchRepo: mr,
		userRepo:  ur,
	}
}

func (cu *matchUsecase) FindAll(ctx context.Context, filter request.MatchRequest) (matchs []domain.Match, err error) {

	matchs, err = cu.matchRepo.FindAll(ctx, filter)
	if err != nil {
		common.LogErrorWithLine(err)
		return matchs, err
	}

	return matchs, err
}

func (u matchUsecase) Match(ctx context.Context, data *domain.Match) (err error) {
	user, err := u.userRepo.FindByID(ctx, data.UserID.String())
	if err != nil {
		return err
	}

	matches, err := u.matchRepo.FindAll(ctx, request.MatchRequest{
		UserID:  data.UserID.String(),
		IsToday: true,
	})
	if err != nil {
		return err
	}

	if len(matches) >= 10 && !user.IsPremium {
		return errors.New("subscribe feature to get more loves")
	}

	if data.IsLike {
		targetUser, err := u.matchRepo.CheckForLike(ctx, request.MatchRequest{
			UserID:   data.TargetUserID.String(),
			TargetID: data.UserID.String(),
			IsLike:   true,
		})
		if err != nil {
			return err
		}

		if targetUser.IsLike {
			targetUser.IsMatch = true
			data.IsMatch = true
			err = u.matchRepo.Update(ctx, &targetUser)
			if err != nil {
				return err
			}
		}
	}

	err = u.matchRepo.Create(ctx, data)
	if err != nil {
		return err
	}

	return err
}
