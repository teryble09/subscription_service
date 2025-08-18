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

type SubscriptionDeleter interface {
	DeleteSubscription(dto.DeleteSubscriptionDTO) error
}

func (srv *SubscriptionService) SubscriptionIDDelete(
	ctx context.Context, params api.SubscriptionIDDeleteParams,
) (api.SubscriptionIDDeleteRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	deleteSub := model.DeleteSubscriptionFromDeleteReq(params)

	deleteDto := dto.NewDeleteSubscriptionDTO(deleteSub)

	err := srv.Storage.DeleteSubscription(deleteDto)

	if errors.Is(err, storage.ErrSubNotFound) {
		logger.Info("Subscription not found",
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
