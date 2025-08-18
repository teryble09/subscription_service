package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DBSubscription struct {
	ID          int64        `db:"id"`
	ServiceName string       `db:"service_name"`
	Price       int          `db:"price"`
	UserID      uuid.UUID    `db:"user_id"`
	StartDate   time.Time    `db:"start_date"`
	EndDate     sql.NullTime `db:"end_date"`
}
