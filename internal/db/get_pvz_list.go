package db

import (
	"context"
	"errors"
	"pvz-service/internal/entities"
	"pvz-service/internal/models"
)

func (s *storage) GetPvzList(ctx context.Context, filters *models.PvzFilterParams) ([]entities.Pvz, []entities.Reception, error) {
	return nil, nil, errors.New("testing plug")
}
