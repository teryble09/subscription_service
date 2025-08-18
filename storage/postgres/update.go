package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/storage"
	"github.com/teryble09/subscription_service/storage/postgres/model"
)

func (s *Storage) UpdateSubscription(updateDto dto.UpdateSubscriptionDTO) (dto.SubscriptionDTO, error) {

	updateSub := model.NewSubscriptionUpdateFromDto(updateDto)

	query, args, err := squirrel.Update("subscriptions").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(updateSub.Map).
		Where("id = ?", updateSub.ID).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return dto.SubscriptionDTO{}, fmt.Errorf("build update query: %w", err)
	}

	var sub model.DBSubscription
	err = s.db.Get(&sub, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return dto.SubscriptionDTO{}, storage.ErrSubNotFound
	}

	if err != nil {
		return dto.SubscriptionDTO{}, fmt.Errorf("execute query: %w", err)
	}

	subDto := model.DBSubscriptionToSubscriptionDTO(sub)

	return subDto, nil
}
