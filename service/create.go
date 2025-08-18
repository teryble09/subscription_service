package service

import (
	"context"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionCreator interface {
	CreateSubstriction(dto.CreateSubscriptionDTO) (dto.SubscriptionDTO, error)
}

func (srv *SubscriptionService) SubscriptionPost(
	ctx context.Context, req *api.CreateSubscriptionReq,
) (api.SubscriptionPostRes, error) {

	logger := ctx.Value("logger").(*slog.Logger)

	//модель бизнес логики из дто ogen
	createSub, err := model.CreateSubscriptionFromCreateReq(req)
	if err != nil {
		logger.Error("Failed to parse create request",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionPostInternalServerError{
			Error: "Failed to parse request",
		}, nil
	}

	// дто для бд
	createDTO := dto.NewCreateSubscriptionDTO(createSub)

	// получили обратно дто
	subDto, err := srv.Storage.CreateSubstriction(createDTO)
	if err != nil {
		// no restrictions on unique so far, should be created
		logger.Error("Failed to create subscription",
			slog.String("error", err.Error()),
		)
		return &api.SubscriptionPostInternalServerError{
			Error: "Failed to create subscription",
		}, nil
	}
	// модель бизнес логики
	subModel := dto.SubscriptionDtoToModel(subDto)
	// превращается опять в дто (весьма "специфичное" для ogen)
	resp := model.SubscriptionIntoApi(&subModel)
	return &resp, nil
}
