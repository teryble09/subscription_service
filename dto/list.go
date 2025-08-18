package dto

import "github.com/teryble09/subscription_service/model"

type ListSubscriptionDTO []SubscriptionDTO

func ListSubscriptionsDtoToModel(s ListSubscriptionDTO) model.ListSubscriptions {
	var list model.ListSubscriptions
	for _, subDto := range s {
		list = append(list, SubscriptionDtoToModel(subDto))
	}
	return list
}
