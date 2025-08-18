package dto

import "github.com/teryble09/subscription_service/model"

type DeleteSubscriptionDTO struct {
	ID int64
}

func NewDeleteSubscriptionDTO(s model.DeleteSubscription) DeleteSubscriptionDTO {
	return DeleteSubscriptionDTO{
		ID: s.ID,
	}
}
