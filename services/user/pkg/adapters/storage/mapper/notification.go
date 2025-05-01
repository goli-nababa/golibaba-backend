package mapper

import (
	"user_service/internal/domain"
	storageTypes "user_service/pkg/adapters/storage/types"
)

func NotificationToStorage(notification *domain.Notification) *storageTypes.Notification {
	if notification == nil {
		return nil
	}

	return &storageTypes.Notification{
		ID:        uint(notification.ID),
		UserID:    notification.UserID,
		Message:   notification.Message,
		Seen:      notification.Seen,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}

func NotificationFromStorage(notification *storageTypes.Notification) *domain.Notification {
	if notification == nil {
		return nil
	}

	return &domain.Notification{
		ID:        domain.NotificationID(notification.ID),
		UserID:    notification.UserID,
		Message:   notification.Message,
		Seen:      notification.Seen,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}

func NotificationsFromStorage(notifications []storageTypes.Notification) []domain.Notification {
	result := make([]domain.Notification, len(notifications))
	for i, notif := range notifications {
		result[i] = *NotificationFromStorage(&notif)
	}
	return result
}
