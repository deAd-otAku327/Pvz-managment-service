package model

import (
	"pvz-service/internal/dto"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapToGetPvzListResponse(t *testing.T) {
	id, datetime, city := 1, time.Now(), enum.Moscow.String()
	status, pType := enum.StatusInProgress.String(), enum.TypeShoes.String()

	inputWithoutReceptions := &models.PvzWithReceptions{
		Pvz: &models.Pvz{
			ID:               id,
			RegistrationDate: datetime,
			City:             city,
		},
		Receptions: nil,
	}

	inputWithoutProducts := &models.PvzWithReceptions{
		Pvz: &models.Pvz{
			ID:               id,
			RegistrationDate: datetime,
			City:             city,
		},
		Receptions: []*models.Reception{{
			ID:       id,
			DateTime: datetime,
			PvzID:    id,
			Products: nil,
			Status:   status,
		}},
	}

	inputFull := &models.PvzWithReceptions{
		Pvz: &models.Pvz{
			ID:               id,
			RegistrationDate: datetime,
			City:             city,
		},
		Receptions: []*models.Reception{{
			ID:       id,
			DateTime: datetime,
			PvzID:    id,
			Products: []*models.Product{{
				ID:          id,
				DateTime:    datetime,
				ReceptionID: id,
				Type:        pType,
			}},
			Status: status,
		}},
	}

	testCases := []struct {
		name     string
		input    models.PvzList
		expected *dto.GetPvzListResponseDTO
	}{
		{name: "empty input", input: nil, expected: &dto.GetPvzListResponseDTO{}},
		{name: "empty receptions && products",
			input: models.PvzList{inputWithoutReceptions},
			expected: &dto.GetPvzListResponseDTO{&dto.PvzWithReceptionsDTO{
				Pvz:        &dto.PvzResponseDTO{ID: id, RegistrationDate: datetime.Format(DateTimeFormat), City: city},
				Receptions: nil,
			}},
		},
		{name: "empty products",
			input: models.PvzList{inputWithoutProducts},
			expected: &dto.GetPvzListResponseDTO{&dto.PvzWithReceptionsDTO{
				Pvz: &dto.PvzResponseDTO{ID: id, RegistrationDate: datetime.Format(DateTimeFormat), City: city},
				Receptions: []*dto.ReceptionWithProductsDTO{{
					Reception: &dto.ReceptionResponseDTO{ID: id, DateTime: datetime.Format(DateTimeFormat), PvzID: id, Status: status},
					Products:  nil,
				}},
			}},
		},
		{name: "full input",
			input: models.PvzList{inputFull},
			expected: &dto.GetPvzListResponseDTO{&dto.PvzWithReceptionsDTO{
				Pvz: &dto.PvzResponseDTO{ID: id, RegistrationDate: datetime.Format(DateTimeFormat), City: city},
				Receptions: []*dto.ReceptionWithProductsDTO{{
					Reception: &dto.ReceptionResponseDTO{ID: id, DateTime: datetime.Format(DateTimeFormat), PvzID: id, Status: status},
					Products: []*dto.ProductResponseDTO{
						{ID: id, DateTime: datetime.Format(DateTimeFormat), Type: pType, ReceptionID: id},
					},
				}},
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := MapToGetPvzListResponse(testCase.input)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
