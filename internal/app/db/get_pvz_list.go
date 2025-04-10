package db

import (
	"context"
	"errors"
	"pvz-service/internal/models"
)

func (s *storage) GetPvzList(ctx context.Context, filters *models.FilterParams) (*models.SummaryInfo, error) {
	return nil, errors.New("testing plug")
}
