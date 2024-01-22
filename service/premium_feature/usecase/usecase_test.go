package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/service/premium_feature/usecase"
	userMock "github.com/iqbalrestu07/datting-apps-api/service/user/mock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestPremiumFeatureUsecase_Subscribe(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	premiumFeatureUsecase := usecase.NewPremiumFeatureUsecase(mockUserRepo)

	mockCtx := context.Background()
	mockUserID := uuid.NewV4()
	t.Run("FindUserByIDError", func(t *testing.T) {
		// Prepare test data
		mockData := domain.PremiumFeature{UserID: mockUserID.String()}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID).Return(domain.User{}, errors.New("find user by ID error"))

		// Execute the Subscribe method
		err := premiumFeatureUsecase.Subscribe(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "find user by ID error", "Expected error message to contain 'find user by ID error'")
	})

	t.Run("SubsribeError", func(t *testing.T) {
		// Prepare test data
		mockData := domain.PremiumFeature{UserID: mockUserID.String(), UnlimitedSwipe: true, Verified: true}
		mockUser := domain.User{Model: domain.Model{ID: mockUserID}}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID).Return(mockUser, nil)

		// Set expectations for Update
		mockUserRepo.EXPECT().Update(mockCtx, gomock.Any()).Return(errors.New("update user error"))

		// Execute the Subscribe method
		err := premiumFeatureUsecase.Subscribe(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "update user error", "Expected error message to contain 'update user error'")
	})

	t.Run("Success", func(t *testing.T) {
		// Prepare test data
		mockData := domain.PremiumFeature{UserID: "user_id", UnlimitedSwipe: true, Verified: true}
		mockUser := domain.User{Model: domain.Model{ID: mockUserID}}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID).Return(mockUser, nil)

		// Set expectations for Update
		mockUserRepo.EXPECT().Update(mockCtx, gomock.Any()).Return(nil)

		// Execute the Subscribe method
		err := premiumFeatureUsecase.Subscribe(mockCtx, mockData)

		// Assertions
		assert.NoError(t, err, "Expected no error")
	})
}
