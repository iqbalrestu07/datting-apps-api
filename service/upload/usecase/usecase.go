package usecase

import (
	"context"

	"github.com/iqbalrestu07/datting-apps-api/domain"
)

type uploadUsecase struct {
	uploadRepo domain.UploadRepository
}

func NewUploadUsecase(ur domain.UploadRepository) uploadUsecase {
	return uploadUsecase{
		uploadRepo: ur,
	}
}

func (u uploadUsecase) Create(ctx context.Context, data *domain.Photo) (err error) {
	err = u.uploadRepo.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
