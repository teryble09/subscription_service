package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/teryble09/subscription_service/dto"
)

func (s *Storage) CalculateCost(calcDto dto.CalculateCostDTO) (dto.CalculateCostResponseDTO, error) {

	builder := squirrel.Select("COALESCE(SUM(price), 0), COUNT(*)").
		From("subscriptions").
		PlaceholderFormat(squirrel.Dollar)

	if calcDto.UserID.Valid {
		builder = builder.Where("user_id = ?", calcDto.UserID.UUID)
	}

	if calcDto.ServiceName.Valid {
		builder = builder.Where("service_name =  ?", calcDto.ServiceName.String)
	}

	if calcDto.StartDate.Valid {
		builder = builder.Where("start_date >= ?", calcDto.StartDate.Time)
	}

	if calcDto.EndDate.Valid {
		builder = builder.Where("end_date <= ?", calcDto.EndDate.Time)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return dto.CalculateCostResponseDTO{}, fmt.Errorf("build query: %w", err)
	}

	var resDto dto.CalculateCostResponseDTO

	row := s.db.QueryRow(query, args...)
	err = row.Scan(&resDto.Cost, &resDto.Count)
	if err != nil {
		err = fmt.Errorf("calculate query: %w", err)
	}

	return resDto, nil
}
