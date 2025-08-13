package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

// model is identical to database table
type Subscription struct {
	ID          int64        `db:"id"`
	ServiceName string       `db:"service_name"`
	Price       int          `db:"price"`
	UserID      uuid.UUID    `db:"user_id"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}

func (sub *Subscription) IntoApiSub() api.Subscription {
	n := api.Subscription{
		ID:          int(sub.ID),
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   dateparse.ParseIntoMMYYYY(sub.StartDate),
	}
	if sub.EndDate.Valid {
		n.EndDate = api.NewOptString(dateparse.ParseIntoMMYYYY(sub.EndDate.Time))
	}
	return n
}

func SubscriptionFromCreateReq(req *api.CreateSubscriptionReq) (Subscription, error) {
	dateStart, err := dateparse.ParseMMYYYY(req.GetStartDate())

	if err != nil {
		// ogen should handle validation with regexp, if date reached here it is internal error
		return Subscription{}, fmt.Errorf("create subscription from req: %w", err)
	}
	sub := Subscription{
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
			return Subscription{}, fmt.Errorf("create subscription from req: %w", err)
		}
		sub.EndDate = sql.NullTime{Time: dateEnd, Valid: true}
	}

	return sub, nil
}
