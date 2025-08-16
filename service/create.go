package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/xlogger"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionCreator interface {
	CreateSubstriction(model.Subscription) (model.Subscription, error)
}

func (srv *SubscriptionService) SubscriptionPost(
	ctx context.Context, req *api.CreateSubscriptionReq,
) (api.SubscriptionPostRes, error) {
	logger := xlogger.WithReqID(srv.Logger, ctx)

	sub, err := model.SubscriptionFromCreateReq(req)
	if err != nil {
		logger.Error("Failed to parse create request",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionPostInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	result, err := srv.Storage.CreateSubstriction(sub)
	if err != nil {
		// no restrictions on unique so far, should be created
		logger.Error("Failed to create subscription",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionPostInternalServerError{
			Error: "Failed to create subscription",
		}, nil
	}

	resp := model.IntoApiSub(&result)
	return &resp, nil
}
