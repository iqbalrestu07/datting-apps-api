package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	domain "github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository will create an object that represent the useroice.Repository interface
func NewUserRepository(conn *gorm.DB) domain.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) FindAll(ctx context.Context, filter request.UserRequest) (users []domain.User, err error) {
	// Add filters to the query
	db := r.Conn.WithContext(ctx)

	if filter.Gender != "" {
		db = db.Where("gender = ?", filter.Gender)
	}
	fmt.Println("excluded id", filter.ExcludedID)
	if len(filter.ExcludedID) > 0 {
		db = db.Where("id not in (?) ", filter.ExcludedID)
	}
	err = db.Preload("UserInterests.Interest").Preload("Photos").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(ctx context.Context, id string) (res domain.User, err error) {
	err = r.Conn.WithContext(ctx).Preload("UserInterests.Interest").Preload("Photos").Where("id = ?", id).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	err = r.Conn.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *userRepository) FindByMultipleID(ctx context.Context, ids []string) (users []domain.User, err error) {
	err = r.Conn.WithContext(ctx).Where("id in (?)", ids).Find(&users).Error
	return users, err
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (err error) {
	err = r.Conn.WithContext(ctx).Create(user).Error
	if err != nil {
		return
	}

	return
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (err error) {
	err = r.Conn.WithContext(ctx).Updates(user).Error
	return
}
