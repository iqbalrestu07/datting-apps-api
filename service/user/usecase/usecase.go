package usecase

import (
	"context"
	"sort"
	"time"

	"github.com/iqbalrestu07/datting-apps-api/common"
	domain "github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
)

type userUsecase struct {
	userRepo         domain.UserRepository
	userInterestRepo domain.UserInterestRepository
	userMatchRepo    domain.MatchRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(cr domain.UserRepository, ir domain.UserInterestRepository, mr domain.MatchRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo:         cr,
		userInterestRepo: ir,
		userMatchRepo:    mr,
	}
}

func (uc *userUsecase) FindAll(ctx context.Context, filter request.UserRequest) (users []domain.User, err error) {

	users, err = uc.userRepo.FindAll(ctx, filter)
	if err != nil {
		common.LogErrorWithLine(err)
		return users, err
	}

	return users, err
}

func (uc *userUsecase) FindByID(ctx context.Context, id string) (res domain.User, err error) {

	res, err = uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return
	}

	resUser, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return resUser, err
}

func (uc *userUsecase) Update(ctx context.Context, ar *domain.User) (err error) {

	ar.UpdatedAt = time.Now()
	return uc.userRepo.Update(ctx, ar)
}

func (uc *userUsecase) Create(ctx context.Context, m *domain.User) (err error) {
	return uc.userRepo.Create(ctx, m)
}

// GetUserListSortedByInterest returns a list of users sorted by the number of shared interests with the given user.
func (uc *userUsecase) GetUserListSortedByInterest(ctx context.Context, userID string) ([]domain.User, error) {
	userMatch, err := uc.userMatchRepo.FindAll(ctx, request.MatchRequest{
		UserID:  userID,
		IsLike:  true,
		IsMatch: true,
	})
	if err != nil {
		return nil, err
	}

	excludedUserID := []string{userID}
	for _, um := range userMatch {
		excludedUserID = append(excludedUserID, um.TargetUserID.String())
	}

	// Get the user's interests
	userInterests, err := uc.userInterestRepo.FindAll(ctx, domain.UserInterest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	// Get all users excluding the given user
	users, err := uc.userRepo.FindAll(ctx, request.UserRequest{
		ExcludedID: excludedUserID,
	})
	if err != nil {
		return nil, err
	}

	// Calculate the number of shared interests for each user
	userWithSharedInterests := make([]domain.User, 0, len(users))
	for _, user := range users {
		sharedInterests := countSharedInterests(userInterests, user.UserInterests)
		user.SharedInterestCount = sharedInterests
		userWithSharedInterests = append(userWithSharedInterests, user)
	}

	// Sort the users by the number of shared interests in descending order
	sort.Slice(userWithSharedInterests, func(i, j int) bool {
		return userWithSharedInterests[i].SharedInterestCount > userWithSharedInterests[j].SharedInterestCount
	})

	return userWithSharedInterests, nil
}

// countSharedInterests calculates the number of shared interests between two sets of interests.
func countSharedInterests(interests1, interests2 []domain.UserInterest) int {
	shared := 0
	interestMap := make(map[string]bool)

	// Create a map of interests for faster lookup
	for _, interest := range interests1 {
		interestMap[interest.ID.String()] = true
	}

	// Count the number of shared interests
	for _, interest := range interests2 {
		if interestMap[interest.ID.String()] {
			shared++
		}
	}

	return shared
}
