package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionLister interface {
	ListSubscriptions() ([]model.Subscription, error)
}

// returns all subscriptions
func (srv *SubscriptionService) SubscriptionGet(ctx context.Context) (api.SubscriptionGetRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	subs, err := srv.Storage.ListSubscriptions()
	if err != nil {
		logger.Error("Could not list subscriptions",
			slog.String("error", err.Error()),
		)
		return &api.Error{
			Error: "Internal error",
		}, nil
	}

	res := api.ListSubscriptionsRes{}
	for _, sub := range subs {
		res = append(res, model.IntoApiSub(&sub))
	}
	return &res, nil

}
