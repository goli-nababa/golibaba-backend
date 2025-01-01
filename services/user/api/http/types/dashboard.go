package types


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
