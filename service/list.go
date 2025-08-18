package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionLister interface {
	ListSubscriptions() (dto.ListSubscriptionDTO, error)
}

// returns all subscriptions
func (srv *SubscriptionService) SubscriptionGet(ctx context.Context) (api.SubscriptionGetRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	listDto, err := srv.Storage.ListSubscriptions()
	if err != nil {
		logger.Error("Could not list subscriptions",
			slog.String("error", err.Error()),
		)
		return &api.Error{
			Error: "Internal error",
		}, nil
	}

	list := dto.ListSubscriptionsDtoToModel(listDto)

	res := model.ListIntoApi(list)

	return &res, nil

}
