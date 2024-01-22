package request

type UserRequest struct {
	ExcludedID []string `json:"excluded_id" query:"excluded_id"`
	UserID     string   `json:"user_id" query:"user_id"`
	Email      string   `json:"email" query:"email"`
	Username   string   `json:"username" query:"username"`
	Gender     string   `json:"gender" query:"gender"`
}
