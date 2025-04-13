package helper

import (
	"context"
	"database/sql"
	"errors"
	"pvz-service/internal/enum"
	"pvz-service/internal/storage/db/dberrors"
	"pvz-service/internal/storage/db/shared/consts"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type DBHepler interface {
	CheckOpenedReceptions(ctx context.Context, tx *sql.Tx, pvzID int) (bool, int, error)
	CatchPQErrors(err error) error
}

type dbHelper struct {
	db *sql.DB
}

func New(db *sql.DB) DBHepler {
	return &dbHelper{
		db: db,
	}
}

func (h *dbHelper) CheckOpenedReceptions(ctx context.Context, tx *sql.Tx, pvzID int) (bool, int, error) {
	query, args, err := sq.Select(consts.ID).
		From(consts.ReceptionsTable).
		Where(sq.Eq{
			consts.ReceptionPvzID:  pvzID,
			consts.ReceptionStatus: enum.StatusInProgress.String(),
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return false, 0, err
	}

	var receptionID int

	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(&receptionID)
	if err == nil {
		return true, receptionID, nil
	}

	if err == sql.ErrNoRows {
		return false, 0, nil
	}

	return false, 0, err
}

func (h *dbHelper) CatchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code.Name() == consts.PQInvalidTextRepresentation {
			return errors.Join(dberrors.ErrEnumTypeViolation, err) // For mode specifics in logs.
		}
		if pqErr.Code.Name() == consts.PQForeignKeyViolation {
			return dberrors.ErrForeignKeyViolation
		}
	}

	return err
}
