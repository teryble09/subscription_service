package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/model"
)

type CostCalculator interface {
	CalculateCost(model.CalculateCostRequest) (sum int, count int, err error)
}

func (srv *SubscriptionService) SubscriptionCalculateTotalCostPost(
	ctx context.Context, req *api.CalculateTotalCostReq,
) (api.SubscriptionCalculateTotalCostPostRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	calcReq, err := model.NewCalculateCostReq(req)
	if err != nil {
		logger.Error("Failed to parse calculate request",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionCalculateTotalCostPostInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	priceSum, count, err := srv.Storage.CalculateCost(calcReq)
	if err != nil {
		logger.Error("Failed to calculate from storage",
			slog.String("erro", err.Error()),
		)
		return &api.SubscriptionCalculateTotalCostPostInternalServerError{
			Error: "Internal",
		}, nil
	}
	return &api.TotalCostRes{
		TotalCost: api.NewOptInt(priceSum),
		Count:     api.NewOptInt(count),
	}, nil
}
