package service

import (
	"context"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionLister interface {
	ListSubscriptions() ([]model.Subscription, error)
}

// returns all subscriptions
func (srv *SubscriptionService) SubscriptionGet(ctx context.Context) (api.SubscriptionGetRes, error) {
	subs, err := srv.Storage.ListSubscriptions()
	if err != nil {
		return &api.Error{
			Error: "Internal error",
		}, nil
	}

	res := api.ListSubscriptionsRes{}
	for _, sub := range subs {
		res = append(res, sub.IntoApiSub())
	}
	return &res, nil

}
