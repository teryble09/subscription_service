package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/api"
	"github.com/teryble09/subscription_service/lib/dateparse"
)

type Subscription struct {
	ID          int64        `db:"id"`
	ServiceName string       `db:"service_name"`
	Price       int          `db:"price"`
	UserID      uuid.UUID    `db:"user_id"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}

// преобразует в дто которое сгенерировал оген

func SubscriptionIntoApi(sub *Subscription) api.Subscription {
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
