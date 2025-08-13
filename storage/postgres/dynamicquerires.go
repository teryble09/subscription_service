package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/teryble09/subscription_service/model"
	"github.com/teryble09/subscription_service/storage"
)

func (s *Storage) UpdateSubscription(su model.SubscriptionUpdate, id int64) (model.Subscription, error) {
	query, args, err := squirrel.Update("subscriptions").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(su).
		Where("id = ?", id).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return model.Subscription{}, fmt.Errorf("build update query: %w", err)
	}

	var sub model.Subscription
	err = s.db.Get(&sub, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Subscription{}, storage.ErrSubNotFound
	}

	if err != nil {
		return model.Subscription{}, fmt.Errorf("execute query: %w", err)
	}

	return sub, nil
}

func (s *Storage) CalculateCost(req model.CalculateCostRequest) (sum int, count int, err error) {
	builder := squirrel.Select("COALESCE(SUM(price), 0), COUNT(*)").From("subscriptions").
		PlaceholderFormat(squirrel.Dollar)
	if (req.UserID != uuid.UUID{}) {
		builder = builder.Where("user_id = ?", req.UserID)
	}
	if req.ServiceName.Valid {
		builder = builder.Where("service_name =  ?", req.ServiceName.String)
	}
	if req.StartDate.Valid {
		builder = builder.Where("start_date >= ?", req.StartDate.Time)
	}
	if req.EndDate.Valid {
		builder = builder.Where("end_date <= ?", req.EndDate.Time)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, 0, fmt.Errorf("build query: %w", err)
	}

	row := s.db.QueryRow(query, args...)
	err = row.Scan(&sum, &count)
	if err != nil {
		err = fmt.Errorf("calculate query: %w", err)
	}

	return
}
