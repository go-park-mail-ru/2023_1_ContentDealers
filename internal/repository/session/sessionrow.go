package session

type sessionRow struct {
	UserID          uint64 `json:"user_id"`
	ExpiresAtString string `json:"expires_at"`
}
