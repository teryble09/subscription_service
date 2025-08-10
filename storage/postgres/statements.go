package postgres

import (
	_ "embed"
	"fmt"

	"github.com/teryble09/subscription_service/model"
)

//go:embed queries/insertsub.sql
var insertSub string

func (s *Storage) prepareStatements() error {
	var err error
	s.stmtInsertSubsription, err = s.db.PrepareNamed(insertSub)
	if err != nil {
		return fmt.Errorf("insert subscription stmt: %w", err)
	}

	return err
}

func (s *Storage) CreateSubstriction(newSub model.Subscription) (model.Subscription, error) {
	var result model.Subscription
	err := s.stmtInsertSubsription.Get(&result, newSub)
	return result, err
}
