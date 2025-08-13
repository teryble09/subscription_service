package postgres

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/teryble09/subscription_service/model"
	"github.com/teryble09/subscription_service/storage"
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

	s.stmtDeleteSubscription, err = s.db.Preparex(deleteSub)
	if err != nil {
		return fmt.Errorf("delete subscriptions stmt: %w", err)
	}

	s.stmtSelectSubscription, err = s.db.Preparex(selectSub)
	if err != nil {
		return fmt.Errorf("select subscriptions stmt: %w", err)
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
	err := s.stmtListSubsriptions.Select(&result, struct{}{})
	return result, err
}

//go:embed queries/deletesub.sql
var deleteSub string

func (s *Storage) DeleteSubscription(id int64) error {
	res, err := s.stmtDeleteSubscription.Exec(id)
	if err != nil {
		return fmt.Errorf("delete sub: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected after sub delete: %w", err)
	}

	if n == 0 {
		return storage.ErrSubNotFound
	}

	return nil
}

//go:embed queries/selectsub.sql
var selectSub string

func (s *Storage) GetSubscription(id int64) (model.Subscription, error) {
	var sub model.Subscription
	err := s.stmtSelectSubscription.Get(&sub, id)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Subscription{}, storage.ErrSubNotFound
	}
	if err != nil {
		return model.Subscription{}, fmt.Errorf("get subscription: %w", err)
	}
	return sub, nil
}
