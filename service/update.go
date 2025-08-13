package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionUpdater interface {
	UpdateSubscription(su model.SubscriptionUpdate, id int64) (model.Subscription, error)
}

func (srv *SubscriptionService) SubscriptionIDPatch(
	ctx context.Context, req *api.UpdateSubscriptionReq, params api.SubscriptionIDPatchParams,
) (api.SubscriptionIDPatchRes, error) {
	reqID := ctx.Value("req_id").(string)
	logger := srv.Logger.With("req_id", reqID)

	su, err := model.NewSubscriptionUpdateFromReq(req)
	if err != nil {
		logger.Error("Failed to parse request into update",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDPatchInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	sub, err := srv.Storage.UpdateSubscription(su, int64(params.ID))
	if err != nil {
		logger.Error("Internal storage error in update req",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDPatchInternalServerError{
			Error: "Internal server error",
		}, nil
	}

	res := sub.IntoApiSub()

	return &res, nil
}
