package dto

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/model"
)

// result of get request from database, contains full information on subscription
type SubscriptionDTO struct {
	ID          int64
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     sql.NullTime
}

func SubscriptionDtoToModel(s SubscriptionDTO) model.Subscription {
	return model.Subscription{
		ID:          s.ID,
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
	}
}
