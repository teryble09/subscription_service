package dto

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/model"
)

type CreateSubscriptionDTO struct {
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     sql.NullTime
}

func NewCreateSubscriptionDTO(s model.CreateSubscription) CreateSubscriptionDTO {
	return CreateSubscriptionDTO{
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
	}
}
