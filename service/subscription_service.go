package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/storage/postgres"
)

type SubscriptionService struct {
	Logger  *slog.Logger
	Storage *postgres.Storage

	api.UnimplementedHandler
}

func (srv *SubscriptionService) SubscriptionPost(
	ctx context.Context, req *api.CreateSubscriptionReq,
) (
	api.SubscriptionPostRes, error,
) {

}
