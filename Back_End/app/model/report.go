package model

type Report struct {
	ID      int    `json:"report_id"`
	UserID  int    `json:"user_id"`
	Username string `json:"username"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
}
