package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/teryble09/subscription_service/model"
)

func (s *Storage) UpdateSubscription(su model.SubscriptionUpdate, id int64) (model.Subscription, error) {
	query, args, err := squirrel.Update("subscriptions").
		Where("id = $1", id).
		SetMap(su).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return model.Subscription{}, fmt.Errorf("build update query: %w", err)
	}

	var sub model.Subscription
	err = s.db.Get(&sub, query, args...)
	if err != nil {
		return model.Subscription{}, fmt.Errorf("execute query: %w", err)
	}

	return sub, nil
}
