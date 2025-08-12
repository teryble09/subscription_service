package postgres

import (
	_ "embed"
	"fmt"

	"github.com/teryble09/subscription_service/model"
)

func (s *Storage) prepareStatements() error {
	var err error
	s.stmtInsertSubsription, err = s.db.PrepareNamed(insertSub)
	if err != nil {
		return fmt.Errorf("insert subscription stmt: %w", err)
	}

	s.stmtListSubsriptions, err = s.db.PrepareNamed(listSubs)
	if err != nil {
		return fmt.Errorf("list subscriptions stmt: %w", err)
	}

	return err
}

//go:embed queries/insertsub.sql
var insertSub string

func (s *Storage) CreateSubstriction(newSub model.Subscription) (model.Subscription, error) {
	var result model.Subscription
	err := s.stmtInsertSubsription.Get(&result, newSub)
	return result, err
}

//go:embed queries/listsubs.sql
var listSubs string

func (s *Storage) ListSubscriptions() ([]model.Subscription, error) {
	var result []model.Subscription
	err := s.stmtListSubsriptions.Select(result, struct{}{})
	return result, err
}
