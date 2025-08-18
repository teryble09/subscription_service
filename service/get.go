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

type SubscriptionGetter interface {
	GetSubscription(dto.GetSubscriptionDTO) (dto.SubscriptionDTO, error)
}

func (srv *SubscriptionService) SubscriptionIDGet(
	ctx context.Context, params api.SubscriptionIDGetParams,
) (api.SubscriptionIDGetRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	getSub := model.GetSubscriptionFromGetReq(params)

	dtoGet := dto.NewGetSubscriptionDTO(getSub)

	dtoSub, err := srv.Storage.GetSubscription(dtoGet)
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

	sub := dto.SubscriptionDtoToModel(dtoSub)

	resp := model.SubscriptionIntoApi(&sub)
	return &resp, nil
}
