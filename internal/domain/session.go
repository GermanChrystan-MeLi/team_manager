package domain

type Session struct {
	ID        string `json:"id"`
	UserID    int    `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
}
