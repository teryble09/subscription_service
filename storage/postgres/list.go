package postgres

import (
	_ "embed"

	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/storage/postgres/model"
)

//go:embed queries/listsubs.sql
var listSubs string

func (s *Storage) ListSubscriptions() (dto.ListSubscriptionDTO, error) {

	var list []model.DBSubscription
	err := s.stmtListSubsriptions.Select(&list, struct{}{})

	dtoList := model.ListDBSubscriptionToDto(list)
	return dtoList, err
}
