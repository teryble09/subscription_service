package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

type CalculateCost struct {
	UserID      uuid.NullUUID
	StartDate   sql.NullTime
	EndDate     sql.NullTime
	ServiceName sql.NullString
}

func NewCalculateCostFromReq(req *api.CalculateTotalCostReq) (CalculateCost, error) {
	var c CalculateCost
	if req.UserID.IsSet() {
		c.UserID = uuid.NullUUID{
			UUID:  req.UserID.Value,
			Valid: true,
		}
	}
	if req.ServiceName.IsSet() {
		c.ServiceName = sql.NullString{
			String: req.ServiceName.Value,
			Valid:  true,
		}
	}
	if req.StartPeriod.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.StartPeriod.Value)
		// ogen should handle validation
		if err != nil {
			return CalculateCost{}, fmt.Errorf("parse start period: %w", err)
		}
		c.StartDate = sql.NullTime{
			Time:  date,
			Valid: true,
		}
	}
	if req.EndPeriod.IsSet() {
		date, err := dateparse.ParseMMYYYY(req.EndPeriod.Value)
		// ogen should handle validation
		if err != nil {
			return CalculateCost{}, fmt.Errorf("parse end period: %w", err)
		}
		c.EndDate = sql.NullTime{
			Time:  date,
			Valid: true,
		}
	}
	return c, nil
}
