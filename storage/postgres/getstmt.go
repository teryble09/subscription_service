package postgres

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/storage"
	"github.com/teryble09/subscription_service/storage/postgres/model"
)

//go:embed queries/selectsub.sql
var selectSub string

func (s *Storage) GetSubscription(getDto dto.GetSubscriptionDTO) (dto.SubscriptionDTO, error) {
	var sub model.DBSubscription
	err := s.stmtSelectSubscription.Get(&sub, getDto.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return dto.SubscriptionDTO{}, storage.ErrSubNotFound
	}
	if err != nil {
		return dto.SubscriptionDTO{}, fmt.Errorf("get subscription: %w", err)
	}

	dtoSub := model.DBSubscriptionToSubscriptionDTO(sub)

	return dtoSub, nil
}
