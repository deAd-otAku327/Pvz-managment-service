package entity

import (
	"pvz-service/internal/entities"
	"pvz-service/internal/enum"
	"pvz-service/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testInput struct {
	pvzs       []*entities.Pvz
	receptions []*entities.Reception
	products   []*entities.Product
}

func TestMapToPvzList(t *testing.T) {
	id, datetime, city := 1, time.Now(), enum.Moscow.String()
	status, pType := enum.StatusInProgress.String(), enum.TypeShoes.String()

	pvzInput := []*entities.Pvz{{
		ID:               id,
		RegistrationDate: datetime,
		City:             city,
	}}

	receptionsInput := []*entities.Reception{{
		ID:       id,
		DateTime: datetime,
		PvzID:    id,
		Status:   status,
	}}

	productsInput := []*entities.Product{{
		ID:          id,
		DateTime:    datetime,
		ReceptionID: id,
		Type:        pType,
	}}

	testCases := []struct {
		name     string
		input    testInput
		expected models.PvzList
	}{
		{name: "empty input", input: testInput{nil, nil, nil}, expected: models.PvzList{}},
		{name: "empty receptions && products",
			input: testInput{pvzInput, nil, nil},
			expected: models.PvzList{&models.PvzWithReceptions{
				Pvz:        &models.Pvz{ID: id, RegistrationDate: datetime, City: city},
				Receptions: nil,
			}},
		},
		{name: "empty products",
			input: testInput{pvzInput, receptionsInput, nil},
			expected: models.PvzList{&models.PvzWithReceptions{
				Pvz:        &models.Pvz{ID: id, RegistrationDate: datetime, City: city},
				Receptions: []*models.Reception{{ID: id, DateTime: datetime, PvzID: id, Status: status}},
			}}},
		{name: "full input",
			input: testInput{pvzInput, receptionsInput, productsInput},
			expected: models.PvzList{&models.PvzWithReceptions{
				Pvz: &models.Pvz{ID: id, RegistrationDate: datetime, City: city},
				Receptions: []*models.Reception{{
					ID: id, DateTime: datetime, PvzID: id, Status: status,
					Products: []*models.Product{{ID: id, DateTime: datetime, ReceptionID: id, Type: pType}}}},
			}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := MapToPvzList(testCase.input.pvzs, testCase.input.receptions, testCase.input.products)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
