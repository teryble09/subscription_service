package dto

import "github.com/teryble09/subscription_service/model"

type GetSubscriptionDTO struct {
	ID int64
}

func NewGetSubscriptionDTO(s model.GetSubscription) GetSubscriptionDTO {
	return GetSubscriptionDTO{
		ID: s.ID,
	}
}
