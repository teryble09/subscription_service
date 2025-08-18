package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

type UpdateSubscription struct {
	ID          int64
	ServiceName sql.NullString
	Price       sql.NullInt64
	UserID      uuid.NullUUID
	StartDate   sql.NullTime
	EndDate     sql.NullTime
}

func UpdateSubscriptionFromReq(
	req *api.UpdateSubscriptionReq, params api.SubscriptionIDPatchParams,
) (UpdateSubscription, error) {

	var us UpdateSubscription
	us.ID = int64(params.ID)

	if req.ServiceName.IsSet() {
		us.ServiceName = sql.NullString{
			String: req.ServiceName.Value,
			Valid:  true,
		}
	}

	if req.Price.IsSet() {
		us.Price = sql.NullInt64{
			Int64: int64(req.Price.Value),
		}
	}

	if req.UserID.IsSet() {
		us.UserID = uuid.NullUUID{
			UUID:  req.UserID.Value,
			Valid: true,
		}
	}

	if req.StartDate.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.StartDate.Value)
		// ogen should handle validation, error is internal
		if err != nil {
			return UpdateSubscription{}, fmt.Errorf("parse start date: %w", err)
		}

		us.StartDate = sql.NullTime{
			Time:  date,
			Valid: true,
		}
	}

	if req.EndDate.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.EndDate.Value)
		// ogen should handle validation, error is internal
		if err != nil {
			return UpdateSubscription{}, fmt.Errorf("parse end date: %w", err)
		}

		us.EndDate = sql.NullTime{
			Time:  date,
			Valid: true,
		}
	}

	return us, nil
}
