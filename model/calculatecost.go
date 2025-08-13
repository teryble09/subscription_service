package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

type CalculateCostRequest struct {
	UserID      uuid.UUID
	StartDate   sql.NullTime
	EndDate     sql.NullTime
	ServiceName sql.NullString
}

func NewCalculateCostRequestFromApiReq(req *api.CalculateTotalCostReq) (CalculateCostRequest, error) {
	c := CalculateCostRequest{}
	if req.UserID.IsSet() {
		c.UserID = req.UserID.Value
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
			return CalculateCostRequest{}, fmt.Errorf("parse start period: %w", err)
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
			return CalculateCostRequest{}, fmt.Errorf("parse end period: %w", err)
		}
		c.EndDate = sql.NullTime{
			Time:  date,
			Valid: true,
		}
	}
	return c, nil
}
