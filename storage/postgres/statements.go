package postgres

import (
	_ "embed"
	"fmt"
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
