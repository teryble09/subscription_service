package dto

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/model"
)

type UpdateSubscriptionDTO struct {
	ID          int64
	ServiceName sql.NullString
	Price       sql.NullInt64
	UserID      uuid.NullUUID
	StartDate   sql.NullTime
	EndDate     sql.NullTime
}

func NewUpdateSubscriptionDTO(m model.UpdateSubscription) UpdateSubscriptionDTO {
	return UpdateSubscriptionDTO{
		ID:          m.ID,
		ServiceName: m.ServiceName,
		Price:       m.Price,
		UserID:      m.UserID,
		StartDate:   m.StartDate,
		EndDate:     m.EndDate,
	}
}
