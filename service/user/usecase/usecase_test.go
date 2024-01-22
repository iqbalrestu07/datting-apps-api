package usecase_test

import (
	"context"
	"sort"
	"testing"

	domain "github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
	matchMock "github.com/iqbalrestu07/datting-apps-api/service/match/mock"
	userMock "github.com/iqbalrestu07/datting-apps-api/service/user/mock"
	"github.com/iqbalrestu07/datting-apps-api/service/user/usecase"
	userInterestMock "github.com/iqbalrestu07/datting-apps-api/service/user_interest/mock"
	uuid "github.com/satori/go.uuid"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_FindAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockUserInterestRepo := userInterestMock.NewMockUserInterestRepository(ctrl)
	mockMatchRepo := matchMock.NewMockMatchRepository(ctrl)
	customerUsecase := usecase.NewUserUsecase(mockUserRepo, mockUserInterestRepo, mockMatchRepo)

	// Prepare mock data
	mockFilter := request.UserRequest{}

	// Mock the repository calls
	mockUsers := []domain.User{
		{Model: domain.Model{
			ID: uuid.FromStringOrNil("b130b469-eee8-4733-8928-9c4405ffbac9"),
		},
			Name:   "iqbal",
			Email:  "iqbal.restu07@gmail.com",
			Gender: "male",
			Bio:    "enter for winner",
		},

		{Model: domain.Model{
			ID: uuid.FromStringOrNil("82f725e4-f74a-4d77-8b2e-c4da8fdc5238"),
		}, Name: "restu",
			Email:  "iqbalrestumaulana@gmail.com",
			Gender: "female",
			Bio:    "enter for winner",
		},
	}

	// Set expectations for FindAll
	mockUserRepo.EXPECT().FindAll(gomock.Any(), mockFilter).
		Return(mockUsers, nil)

	// Execute the FindAll method
	result, err := customerUsecase.FindAll(context.Background(), mockFilter)

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Len(t, result, len(mockUsers), "Expected customers length to match")
}

func TestUserUsecase_FindByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockUserInterestRepo := userInterestMock.NewMockUserInterestRepository(ctrl)
	mockUserMatchRepo := matchMock.NewMockMatchRepository(ctrl)

	userUsecase := usecase.NewUserUsecase(mockUserRepo, mockUserInterestRepo, mockUserMatchRepo)

	// Prepare mock data
	mockUserID := uuid.NewV4()

	// Mock the repository calls
	expectedUser := domain.User{Model: domain.Model{ID: mockUserID}, Name: "Alice", Gender: "female"}

	// Set expectations for FindByID
	mockUserRepo.EXPECT().FindByID(gomock.Any(), mockUserID.String()).Return(expectedUser, nil).AnyTimes()

	// Execute the FindByID method
	result, err := userUsecase.FindByID(context.Background(), mockUserID.String())

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedUser, result, "Expected user to match")
}

func TestUserUsecase_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockUserInterestRepo := userInterestMock.NewMockUserInterestRepository(ctrl)
	mockUserMatchRepo := matchMock.NewMockMatchRepository(ctrl)

	userUsecase := usecase.NewUserUsecase(mockUserRepo, mockUserInterestRepo, mockUserMatchRepo)

	mockUserID := uuid.NewV4()

	// Prepare mock data
	mockUserToUpdate := &domain.User{Model: domain.Model{ID: mockUserID}, Name: "Alice", Gender: "female"}

	// Set expectations for Update
	mockUserRepo.EXPECT().Update(gomock.Any(), mockUserToUpdate).Return(nil)

	// Execute the Update method
	err := userUsecase.Update(context.Background(), mockUserToUpdate)

	// Assertions
	assert.NoError(t, err, "Expected no error")
}

func TestUserUsecase_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockUserInterestRepo := userInterestMock.NewMockUserInterestRepository(ctrl)
	mockUserMatchRepo := matchMock.NewMockMatchRepository(ctrl)

	userUsecase := usecase.NewUserUsecase(mockUserRepo, mockUserInterestRepo, mockUserMatchRepo)

	// Prepare mock data
	mockNewUser := &domain.User{Name: "Alice", Gender: "female"}

	// Set expectations for Create
	mockUserRepo.EXPECT().Create(gomock.Any(), mockNewUser).Return(nil)

	// Execute the Create method
	err := userUsecase.Create(context.Background(), mockNewUser)

	// Assertions
	assert.NoError(t, err, "Expected no error")
}
func TestUserUsecase_GetUserListSortedByInterest(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	mockUserInterestRepo := userInterestMock.NewMockUserInterestRepository(ctrl)
	mockUserMatchRepo := matchMock.NewMockMatchRepository(ctrl)

	uc := usecase.NewUserUsecase(mockUserRepo, mockUserInterestRepo, mockUserMatchRepo)

	// Prepare test data
	mockUserID := uuid.NewV4()
	mockUserIDMatch1 := uuid.NewV4()
	mockUserIDMatch2 := uuid.NewV4()
	mockTargetUserID1 := uuid.NewV4()
	mockTargetUserID2 := uuid.NewV4()

	mockUserMatches := []domain.Match{
		{UserID: mockUserID, TargetUserID: mockUserIDMatch1, IsLike: true},
		{UserID: mockUserID, TargetUserID: mockUserIDMatch2, IsLike: true},
	}

	mockUserInterests := []domain.UserInterest{
		{Model: domain.Model{
			ID: mockUserID,
		}, UserID: mockUserID.String(),
		},
		{Model: domain.Model{
			ID: mockUserID,
		}, UserID: mockUserID.String(),
		},
	}

	mockUsers := []domain.User{
		{Model: domain.Model{ID: mockTargetUserID1}, Name: "Bob", Gender: "male"},
		{Model: domain.Model{ID: mockTargetUserID2}, Name: "Charlie", Gender: "male"},
	}

	// Set expectations for repository calls
	mockUserMatchRepo.EXPECT().FindAll(gomock.Any(), request.MatchRequest{UserID: mockUserID.String(), IsLike: true, IsMatch: true}).Return(mockUserMatches, nil)
	mockUserInterestRepo.EXPECT().FindAll(gomock.Any(), domain.UserInterest{UserID: mockUserID.String()}).Return(mockUserInterests, nil)
	mockUserRepo.EXPECT().FindAll(gomock.Any(), request.UserRequest{ExcludedID: []string{mockUserID.String(), mockUserIDMatch1.String(), mockUserIDMatch2.String()}}).Return(mockUsers, nil)

	// Execute the use case method
	result, err := uc.GetUserListSortedByInterest(context.Background(), mockUserID.String())

	// Assertions
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected a result")

	// Check if the result is sorted correctly
	expectedOrder := []string{mockTargetUserID1.String(), mockTargetUserID2.String()}
	actualOrder := make([]string, len(result))
	for i, user := range result {
		actualOrder[i] = user.ID.String()
	}

	sort.Strings(expectedOrder)
	sort.Strings(actualOrder)

	assert.Equal(t, expectedOrder, actualOrder, "Expected users to be sorted by shared interests")
}
