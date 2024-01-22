package request

type MatchRequest struct {
	IsLike   bool
	IsMatch  bool
	UserID   string
	TargetID string
	IsToday  bool
}
