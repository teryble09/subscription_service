package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/dto"
)

type DBCreateSubParams struct {
	ServiceName string       `db:"service_name"`
	Price       int          `db:"price"`
	UserID      uuid.UUID    `db:"user_id"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}

func DBCreateSubParamsFromDTO(s dto.CreateSubscriptionDTO) DBCreateSubParams {
	return DBCreateSubParams{
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
	}
}
