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
	reqID := ctx.Value("req_id").(string)
	logger := srv.Logger.With("req_id", reqID)

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
		res = append(res, sub.IntoApiSub())
	}
	return &res, nil

}
