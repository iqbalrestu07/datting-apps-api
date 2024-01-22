package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/service/match/usecase"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	matchMock "github.com/iqbalrestu07/datting-apps-api/service/match/mock"
	userMock "github.com/iqbalrestu07/datting-apps-api/service/user/mock"
)

func TestMatchUsecase_Match(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockMatchRepo := matchMock.NewMockMatchRepository(ctrl)
	matchUsecase := usecase.NewMatchUsecase(mockMatchRepo, mockUserRepo)

	mockCtx := context.Background()
	mockUserID := uuid.NewV4()
	targetUserID := uuid.NewV4()
	t.Run("FindUserByIDError", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{}, errors.New("user not found"))

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "user not found", "Expected error message to contain 'user not found'")
	})

	t.Run("DailyLimitReachedForNonPremiumUser", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{IsPremium: false}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}, nil)

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "subscribe feature to get more loves", "Expected error message to contain 'subscribe feature to get more loves'")
	})

	t.Run("DailyLimitNotReachedForPremiumUser", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{IsPremium: true}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}, nil)
		mockMatchRepo.EXPECT().CheckForLike(mockCtx, gomock.Any()).Return(domain.Match{}, nil)
		mockMatchRepo.EXPECT().Create(mockCtx, gomock.Any()).Return(nil)

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.NoError(t, err, "Expected no error for premium user")
	})

	t.Run("CheckForLikeError", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, TargetUserID: targetUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{}, nil)

		// Set expectations for CheckForLike
		mockMatchRepo.EXPECT().CheckForLike(mockCtx, gomock.Any()).Return(domain.Match{}, errors.New("check for like error"))

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "check for like error", "Expected error message to contain 'check for like error'")
	})

	t.Run("UpdateMatchError", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, TargetUserID: targetUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{}, nil)

		// Set expectations for CheckForLike
		mockMatchRepo.EXPECT().CheckForLike(mockCtx, gomock.Any()).Return(domain.Match{IsLike: true}, nil)

		// Set expectations for Update
		mockMatchRepo.EXPECT().Update(mockCtx, gomock.Any()).Return(errors.New("update match error"))

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "update match error", "Expected error message to contain 'update match error'")
	})

	t.Run("CreateMatchError", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, TargetUserID: targetUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{}, nil)

		// Set expectations for CheckForLike
		mockMatchRepo.EXPECT().CheckForLike(mockCtx, gomock.Any()).Return(domain.Match{IsLike: false}, nil)

		// Set expectations for Create
		mockMatchRepo.EXPECT().Create(mockCtx, mockData).Return(errors.New("create match error"))

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "create match error", "Expected error message to contain 'create match error'")
	})

	t.Run("Success", func(t *testing.T) {
		// Prepare test data
		mockData := &domain.Match{UserID: mockUserID, TargetUserID: targetUserID, IsLike: true}

		// Set expectations for FindByID
		mockUserRepo.EXPECT().FindByID(mockCtx, mockData.UserID.String()).Return(domain.User{}, nil)

		// Set expectations for FindAll
		mockMatchRepo.EXPECT().FindAll(mockCtx, gomock.Any()).Return([]domain.Match{}, nil)

		// Set expectations for CheckForLike
		mockMatchRepo.EXPECT().CheckForLike(mockCtx, gomock.Any()).Return(domain.Match{IsLike: false}, nil)

		// Set expectations for Create
		mockMatchRepo.EXPECT().Create(mockCtx, mockData).Return(nil)

		// Execute the Match method
		err := matchUsecase.Match(mockCtx, mockData)

		// Assertions
		assert.NoError(t, err, "Expected no error")
	})

}
