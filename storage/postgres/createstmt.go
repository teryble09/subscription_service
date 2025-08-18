package postgres

import (
	_ "embed"

	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/storage/postgres/model"
)

//go:embed queries/insertsub.sql
var insertSub string

func (s *Storage) CreateSubstriction(createDTO dto.CreateSubscriptionDTO) (dto.SubscriptionDTO, error) {

	params := model.DBCreateSubParamsFromDTO(createDTO)

	var dbsub model.DBSubscription
	err := s.stmtInsertSubsription.Get(&dbsub, params)

	subDto := model.DBSubscriptionToSubscriptionDTO(dbsub)

	return subDto, err
}
