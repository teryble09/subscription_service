package model

import (
	"fmt"

	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

// for squirrel (sql builder), set only given columns
type SubscriptionUpdate map[string]any

func NewSubscriptionUpdateFromReq(req *api.UpdateSubscriptionReq) (SubscriptionUpdate, error) {
	su := make(map[string]any, 5)

	if req.ServiceName.IsSet() {
		su["service_name"] = req.ServiceName.Value
	}

	if req.Price.IsSet() {
		su["price"] = req.Price.Value
	}

	if req.UserID.IsSet() {
		su["user_id"] = req.UserID.Value
	}

	if req.StartDate.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.StartDate.Value)
		// ogen should handle validation, error is internal
		if err != nil {
			return nil, fmt.Errorf("parse start date: %w", err)
		}
		su["start_date"] = date
	}

	if req.EndDate.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.EndDate.Value)
		// ogen should handle validation, error is internal
		if err != nil {
			return nil, fmt.Errorf("parse end date: %w", err)
		}
		su["end_date"] = date
	}

	return su, nil
}
