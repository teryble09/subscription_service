package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

type CreateSubscription struct {
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     sql.NullTime
}

func CreateSubscriptionFromCreateReq(req *api.CreateSubscriptionReq) (CreateSubscription, error) {
	dateStart, err := dateparse.ParseMMYYYY(req.GetStartDate())

	if err != nil {
		// ogen should handle validation with regexp, if date reached here it is internal error
		return CreateSubscription{}, fmt.Errorf("create subscription from req: %w", err)
	}
	sub := CreateSubscription{
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
			return CreateSubscription{}, fmt.Errorf("create subscription from req: %w", err)
		}
		sub.EndDate = sql.NullTime{Time: dateEnd, Valid: true}
	}

	return sub, nil
}
