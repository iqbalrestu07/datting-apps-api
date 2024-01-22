package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iqbalrestu07/datting-apps-api/app/server/middleware"
	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/service/auth/usecase"
	userMock "github.com/iqbalrestu07/datting-apps-api/service/user/mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthUsecase_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	mockCtx := context.Background()

	t.Run("InvalidEmailError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{}, errors.New("some error"))

		// Execute the Login method
		token, err := authUsecase.Login(mockCtx, mockData)

		// Assertions
		assert.Empty(t, token, "Expected an empty token")
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "invalid", "Expected error message to contain 'invalid'")
	})

	t.Run("InvalidPasswordError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail, Password: "password"}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{Password: "hashed_password"}, nil)

		// Execute the Login method
		token, err := authUsecase.Login(mockCtx, mockData)

		// Assertions
		assert.Empty(t, token, "Expected an empty token")
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "invalid password", "Expected error message to contain 'invalid password'")
	})

	t.Run("Success", func(t *testing.T) {
		// Prepare test data
		password := "satuduatiga"
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail, Password: password}
		hashedPassword, _ := common.HashPassword(password)
		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{Password: hashedPassword}, nil)
		token := middleware.NewJWTAuthService().GenerateToken(domain.User{
			Email:    mockEmail,
			Password: password,
		})
		// Execute the Login method
		userToken, err := authUsecase.Login(mockCtx, mockData)
		// Assertions
		assert.Equal(t, userToken, token, "Expected generated token to match")
		assert.NoError(t, err, "Expected no error")
	})
}
func TestAuthUsecase_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userMock.NewMockUserRepository(ctrl)
	authUsecase := usecase.NewAuthUsecase(mockUserRepo)

	mockCtx := context.Background()

	t.Run("UserFoundError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail}
		mockUser := domain.User{Email: mockEmail}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(mockUser, errors.New("some error"))

		// Execute the Register method
		err := authUsecase.Register(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "some error", "Expected error message to contain 'some error'")
	})

	t.Run("UserAlreadyRegisteredError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail}
		mockUser := domain.User{Email: mockEmail}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(mockUser, nil)

		// Execute the Register method
		err := authUsecase.Register(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "email has been registered", "Expected error message to contain 'email has been registered'")
	})

	t.Run("PasswordMismatchError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail, Password: "password", VerifyPassword: "mismatch"}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{}, gorm.ErrRecordNotFound)

		// Execute the Register method
		err := authUsecase.Register(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "password doesn't match", "Expected error message to contain 'password doesn't match'")
	})

	t.Run("CreateUserError", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail, Password: "password", VerifyPassword: "password"}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{}, gorm.ErrRecordNotFound)

		// Set expectations for Create
		mockUserRepo.EXPECT().Create(mockCtx, gomock.Any()).Return(errors.New("create user error"))

		// Execute the Register method
		err := authUsecase.Register(mockCtx, mockData)

		// Assertions
		assert.Error(t, err, "Expected an error")
		assert.Contains(t, err.Error(), "create user error", "Expected error message to contain 'create user error'")
	})

	t.Run("Success", func(t *testing.T) {
		// Prepare test data
		mockEmail := "test@example.com"
		mockData := domain.Auth{Email: mockEmail, Password: "password", VerifyPassword: "password"}

		// Set expectations for FindByEmail
		mockUserRepo.EXPECT().FindByEmail(mockCtx, mockEmail).Return(domain.User{}, gorm.ErrRecordNotFound)

		// Set expectations for Create
		mockUserRepo.EXPECT().Create(mockCtx, gomock.Any()).Return(nil)

		// Execute the Register method
		err := authUsecase.Register(mockCtx, mockData)

		// Assertions
		assert.NoError(t, err, "Expected no error")
	})
}
