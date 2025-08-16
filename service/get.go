package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/xlogger"
	"github.com/teryble09/subscription_service/model"
	"github.com/teryble09/subscription_service/storage"
)

type SubscriptionGetter interface {
	GetSubscription(id int64) (model.Subscription, error)
}

func (srv *SubscriptionService) SubscriptionIDGet(
	ctx context.Context, params api.SubscriptionIDGetParams,
) (api.SubscriptionIDGetRes, error) {

	logger := xlogger.WithReqID(srv.Logger, ctx)

	sub, err := srv.Storage.GetSubscription(int64(params.ID))
	if errors.Is(err, storage.ErrSubNotFound) {
		logger.Info("Subscription not found",
			slog.Int("id", params.ID),
		)
		return &api.SubscriptionIDGetNotFound{
			Error: "Subscription not found",
		}, nil
	}
	if err != nil {
		logger.Error("Delete subscription internal error",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDGetInternalServerError{
			Error: "Internal error",
		}, nil
	}
	resp := model.IntoApiSub(&sub)
	return &resp, nil
}
