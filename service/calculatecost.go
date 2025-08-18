package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/model"
)

type CostCalculator interface {
	CalculateCost(dto.CalculateCostDTO) (dto.CalculateCostResponseDTO, error)
}

func (srv *SubscriptionService) SubscriptionCalculateTotalCostPost(
	ctx context.Context, req *api.CalculateTotalCostReq,
) (api.SubscriptionCalculateTotalCostPostRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	calculateCost, err := model.NewCalculateCostFromReq(req)
	if err != nil {
		logger.Error("Failed to parse calculate request",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionCalculateTotalCostPostInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	calculateDto := dto.NewCalculateCostDTO(calculateCost)

	calcAnswerDto, err := srv.Storage.CalculateCost(calculateDto)
	if err != nil {
		logger.Error("Failed to calculate from storage",
			slog.String("erro", err.Error()),
		)
		return &api.SubscriptionCalculateTotalCostPostInternalServerError{
			Error: "Internal",
		}, nil
	}

	return &api.TotalCostRes{
		TotalCost: api.NewOptInt(calcAnswerDto.Cost),
		Count:     api.NewOptInt(calcAnswerDto.Count),
	}, nil
}
