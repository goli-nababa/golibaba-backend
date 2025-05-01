package types

type GetUserHistoryRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
type UserHistoryResponse struct {
	Action    string `json:"action"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
}

type NotificationResponse struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Seen      bool   `json:"seen"`
	CreatedAt string `json:"created_at"`
}

type CreateNotificationRequest struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}
