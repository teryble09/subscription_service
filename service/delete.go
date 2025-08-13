package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/storage"
)

type SubscriptionDeleter interface {
	DeleteSubscription(id int64) error
}

func (srv *SubscriptionService) SubscriptionIDDelete(
	ctx context.Context, params api.SubscriptionIDDeleteParams,
) (api.SubscriptionIDDeleteRes, error) {
	reqID := ctx.Value("req_id").(string)
	logger := srv.Logger.With("req_id", reqID)

	err := srv.Storage.DeleteSubscription(int64(params.ID))

	if errors.Is(err, storage.ErrSubNotFound) {
		srv.Logger.Info("Subscription not found",
			slog.Int("id", params.ID),
		)
		return &api.SubscriptionIDDeleteNotFound{
			Error: "Subscription not found",
		}, nil
	}

	if err != nil {
		logger.Error("Delete subscription internal error",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDDeleteInternalServerError{
			Error: "Internal error",
		}, nil
	}

	return &api.SubscriptionIDDeleteNoContent{}, nil
}
