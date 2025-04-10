package service

import (
	"context"
	"net/http"
	"pvz-service/pkg/werrors"
	"strconv"
)

func (s *pvzService) DeleteProduct(ctx context.Context, id string) werrors.Werror {
	pvzId, err := strconv.Atoi(id)
	if err != nil {
		return werrors.New(errInvalidPvzID, http.StatusBadRequest)
	}

	err = s.storage.DeleteProduct(ctx, pvzId)
	if err != nil {
		s.logger.Error("delete product: " + err.Error())
		return werrors.New(errSmthWentWrong, http.StatusInternalServerError)
	}

	return nil
}
