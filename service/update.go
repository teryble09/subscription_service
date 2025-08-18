package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/model"
	"github.com/teryble09/subscription_service/storage"
)

type SubscriptionUpdater interface {
	UpdateSubscription(dto.UpdateSubscriptionDTO) (dto.SubscriptionDTO, error)
}

func (srv *SubscriptionService) SubscriptionIDPatch(
	ctx context.Context, req *api.UpdateSubscriptionReq, params api.SubscriptionIDPatchParams,
) (api.SubscriptionIDPatchRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	updateSub, err := model.UpdateSubscriptionFromReq(req, params)
	if err != nil {
		logger.Error("Failed to parse request into update",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionIDPatchInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	updateDto := dto.NewUpdateSubscriptionDTO(updateSub)

	subDto, err := srv.Storage.UpdateSubscription(updateDto)
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

	sub := dto.SubscriptionDtoToModel(subDto)

	res := model.SubscriptionIntoApi(&sub)

	return &res, nil
}
