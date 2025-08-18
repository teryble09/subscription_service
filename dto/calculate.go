package dto

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/model"
)

type CalculateCostDTO struct {
	UserID      uuid.NullUUID
	StartDate   sql.NullTime
	EndDate     sql.NullTime
	ServiceName sql.NullString
}

func NewCalculateCostDTO(c model.CalculateCost) CalculateCostDTO {
	return CalculateCostDTO{
		UserID:      c.UserID,
		StartDate:   c.StartDate,
		EndDate:     c.EndDate,
		ServiceName: c.ServiceName,
	}
}
