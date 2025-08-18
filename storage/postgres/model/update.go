package model

import (
	"github.com/teryble09/subscription_service/dto"
)

type DBUpdateSubscription struct {
	ID int64
	// for squirrel (sql builder), set only given columns
	Map map[string]any
}

func NewSubscriptionUpdateFromDto(updateDto dto.UpdateSubscriptionDTO) DBUpdateSubscription {
	var us DBUpdateSubscription

	us.ID = updateDto.ID

	us.Map = make(map[string]any, 5)

	if updateDto.ServiceName.Valid {
		us.Map["service_name"] = updateDto.ServiceName.String
	}

	if updateDto.Price.Valid {
		us.Map["price"] = updateDto.Price.Int64
	}

	if updateDto.UserID.Valid {
		us.Map["user_id"] = updateDto.UserID.UUID
	}

	if updateDto.StartDate.Valid {
		us.Map["start_date"] = updateDto.StartDate.Time
	}

	if updateDto.EndDate.Valid {
		us.Map["end_date"] = updateDto.EndDate.Time
	}

	return us
}
