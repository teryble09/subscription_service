package model

import "github.com/teryble09/subscription_service/api"

type GetSubscription struct {
	ID int64
}

func GetSubscriptionFromGetReq(params api.SubscriptionIDGetParams) GetSubscription {
	return GetSubscription{
		ID: int64(params.ID),
	}
}
