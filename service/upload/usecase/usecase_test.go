package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	uploadMock "github.com/iqbalrestu07/datting-apps-api/service/upload/mock"
	"github.com/iqbalrestu07/datting-apps-api/service/upload/usecase"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUploadUsecase_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUploadRepo := uploadMock.NewMockUploadRepository(ctrl)
	uploadUsecase := usecase.NewUploadUsecase(mockUploadRepo)

	mockCtx := context.Background()
	mockUserID := uuid.NewV4()

	t.Run("CreateError", func(t *testing.T) {
		mockData := &domain.Photo{UserID: mockUserID}

		mockUploadRepo.EXPECT().Create(mockCtx, mockData).Return(errors.New("create error"))

		err := uploadUsecase.Create(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "create error", "Expected error message to contain 'create error'")
	})

	t.Run("Success", func(t *testing.T) {
		mockData := &domain.Photo{UserID: mockUserID}

		mockUploadRepo.EXPECT().Create(mockCtx, mockData).Return(nil)

		err := uploadUsecase.Create(mockCtx, mockData)

		// Assertions
		assert.NoError(t, err, "Expected no error")
	})
}
