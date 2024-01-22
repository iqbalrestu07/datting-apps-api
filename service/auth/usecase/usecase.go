package usecase

import (
	"context"
	"errors"

	"github.com/iqbalrestu07/datting-apps-api/app/server/middleware"
	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"gorm.io/gorm"
)

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(ur domain.UserRepository) authUsecase {
	return authUsecase{
		userRepo: ur,
	}
}

func (u authUsecase) Login(ctx context.Context, data domain.Auth) (token string, err error) {

	user, err := u.userRepo.FindByEmail(ctx, data.Email)
	if err != nil || user.ID.String() == "" {
		return "", errors.New("invalid email")
	}

	err = common.ComparePassword(user.Password, data.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	token = middleware.NewJWTAuthService().GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u authUsecase) Register(ctx context.Context, data domain.Auth) (err error) {
	user, err := u.userRepo.FindByEmail(ctx, data.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if user.Email != "" {
		return errors.New("email has been registered")
	}

	if data.Password != data.VerifyPassword {
		return errors.New("password doesn't match")
	}

	passwordHash, err := common.HashPassword(data.Password)
	if err != nil {
		return err
	}
	user = domain.User{
		Email:    data.Email,
		Password: passwordHash,
		Gender:   data.Gender,
	}

	err = u.userRepo.Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
