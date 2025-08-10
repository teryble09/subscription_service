package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
	"github.com/teryble09/subscription_service/model"
)

type SubscriptionCreator interface {
	CreateSubstriction(model.Subscription) (model.Subscription, error)
}

func (srv *SubscriptionService) SubscriptionPost(
	ctx context.Context, req *api.CreateSubscriptionReq,
) (api.SubscriptionPostRes, error) {
	reqID := ctx.Value("req_id").(string)
	logger := srv.Logger.With("req_id", reqID)

	dateStart, err := dateparse.ParseMMYYYY(req.GetStartDate())

	if err != nil {
		// ogen should handle validation with regexp, if date reached here it is internal error
		logger.Error("Failed to parse start date",
			slog.String("error", err.Error()),
			slog.String("date", req.GetStartDate()),
		)
		return &api.SubscriptionPostInternalServerError{
			Error: "Failed to parse start date",
		}, nil
	}

	sub := model.Subscription{
		ServiceName: req.GetServiceName(),
		Price:       req.GetPrice(),
		UserID:      req.GetUserID(),
		StartDate:   dateStart,
	}

	// do not forget to parse end date
	if req.GetEndDate().IsSet() {
		dateEnd, err := dateparse.ParseMMYYYY(req.GetEndDate().Value)
		if err != nil {
			// ogen should handle validation with regexp, if date reached here it is internal error
			logger.Error("Failed to parse end date",
				slog.String("error", err.Error()),
				slog.String("date", req.GetEndDate().Value),
			)
			return &api.SubscriptionPostInternalServerError{
				Error: "Failed to parse end date",
			}, nil
		}
		sub.EndDate = sql.NullTime{Time: dateEnd, Valid: true}
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

	resp := &api.Subscription{
		ID:          int(result.ID),
		ServiceName: req.ServiceName,
		Price:       result.Price,
		UserID:      result.UserID,
		StartDate:   dateparse.ParseIntoMMYYYY(result.StartDate),
	}

	// do not forget end time
	if result.EndDate.Valid {
		resp.EndDate = api.NewOptString(dateparse.ParseIntoMMYYYY(result.EndDate.Time))
	}

	return resp, nil
}
