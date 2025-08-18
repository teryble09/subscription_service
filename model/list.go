package model

import "github.com/teryble09/subscription_service/api"

type ListSubscriptions []Subscription

func ListIntoApi(list ListSubscriptions) api.ListSubscriptionsRes {
	var res api.ListSubscriptionsRes
	for _, s := range list {
		res = append(res, SubscriptionIntoApi(&s))
	}
	return res
}
