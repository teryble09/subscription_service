package model

import "github.com/teryble09/subscription_service/api"

type DeleteSubscription struct {
	ID int64
}

func DeleteSubscriptionFromDeleteReq(params api.SubscriptionIDDeleteParams) DeleteSubscription {
	return DeleteSubscription{
		ID: int64(params.ID),
	}
}
