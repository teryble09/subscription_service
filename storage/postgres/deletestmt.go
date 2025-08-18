package postgres

import (
	_ "embed"
	"fmt"

	"github.com/teryble09/subscription_service/dto"
	"github.com/teryble09/subscription_service/storage"
)

//go:embed queries/deletesub.sql
var deleteSub string

func (s *Storage) DeleteSubscription(deleteDto dto.DeleteSubscriptionDTO) error {
	// тут нужно только id поэтому не стал делать отдельную модель в storage/model
	res, err := s.stmtDeleteSubscription.Exec(deleteDto.ID)
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
