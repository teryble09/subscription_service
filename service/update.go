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

type SubscriptionUpdater interface {
	UpdateSubscription(su model.SubscriptionUpdate, id int64) (model.Subscription, error)
}

func (srv *SubscriptionService) SubscriptionIDPatch(
	ctx context.Context, req *api.UpdateSubscriptionReq, params api.SubscriptionIDPatchParams,
) (api.SubscriptionIDPatchRes, error) {

	logger := xlogger.WithReqID(srv.Logger, ctx)

	su, err := model.NewSubscriptionUpdateFromReq(req)
	if errors.Is(err, model.ErrEmptyUpdateRequest) {
		logger.Info("Empty update request")
		return &api.SubscriptionIDPatchBadRequest{
			Error: "Empty request",
		}, nil
	}

	if err != nil {
		logger.Error("Failed to parse request into update",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDPatchInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	sub, err := srv.Storage.UpdateSubscription(su, int64(params.ID))
	if errors.Is(err, storage.ErrSubNotFound) {
		logger.Info("Subscription not found")
		return &api.SubscriptionIDPatchBadRequest{
			Error: "Subscription not found",
		}, nil
	}

	if err != nil {
		logger.Error("Internal storage error in update req",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDPatchInternalServerError{
			Error: "Internal server error",
		}, nil
	}

	res := model.IntoApiSub(&sub)

	return &res, nil
}
