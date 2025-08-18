package model

import "github.com/teryble09/subscription_service/dto"

func DBSubscriptionToSubscriptionDTO(s DBSubscription) dto.SubscriptionDTO {
	return dto.SubscriptionDTO{
		ID:          s.ID,
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
	}
}

func ListDBSubscriptionToDto(list []DBSubscription) dto.ListSubscriptionDTO {
	var dtoList dto.ListSubscriptionDTO
	for _, s := range list {
		dtoList = append(dtoList, DBSubscriptionToSubscriptionDTO(s))
	}
	return dtoList
}
