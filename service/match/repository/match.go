package repository

import (
	"context"
	"time"

	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
	"gorm.io/gorm"
)

type matchRepository struct {
	Conn *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *matchRepository {
	return &matchRepository{Conn: db}
}

func (r matchRepository) FindAll(ctx context.Context, filter request.MatchRequest) (matches []domain.Match, err error) {
	db := r.Conn.WithContext(ctx)
	if filter.IsToday {
		todayStart := time.Now().Truncate(24 * time.Hour)
		todayEnd := todayStart.Add(24 * time.Hour)
		db = db.Where("created_at BETWEEN ? AND ?", todayStart, todayEnd)
	}
	if filter.UserID != "" {
		db = db.Where("user_id = ?", filter.UserID)
	}
	if filter.IsLike || filter.IsMatch {
		db = db.Where("is_like = ? or is_match = ?", filter.IsLike, filter.IsMatch)
	}

	err = db.Preload("User").Preload("TargetUser").Find(&matches).Error
	return matches, err
}

func (r *matchRepository) CheckForMatch(ctx context.Context, userID, matchedID string) (match domain.Match, err error) {
	err = r.Conn.
		Where("(user_id = ? AND target_id = ? AND liked = ?) OR (user_id = ? AND target_id = ? AND liked = ?)", userID, matchedID, true, matchedID, userID, true).
		Find(&match).Error
	if err != nil {
		return match, err
	}
	return match, nil
}

func (r *matchRepository) CheckForLike(ctx context.Context, filter request.MatchRequest) (match domain.Match, err error) {
	db := r.Conn.WithContext(ctx).Where("user_id = ? AND target_user_id = ?", filter.UserID, filter.TargetID)
	if filter.IsLike {
		db = db.Where("is_like = ?", filter.IsLike)
	}
	err = db.Find(&match).Error
	return match, err
}

func (r *matchRepository) Create(ctx context.Context, match *domain.Match) error {
	return r.Conn.WithContext(ctx).Create(match).Error
}

func (r *matchRepository) Update(ctx context.Context, match *domain.Match) error {
	return r.Conn.WithContext(ctx).Updates(match).Error
}
