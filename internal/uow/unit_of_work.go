package uow

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/infrastructure/repositories"
)

type Uow struct {
	db *sql.DB
}

func NewUow(db *sql.DB) UnitOfWork {
	return &Uow{
		db: db,
	}
}

func (u *Uow) Do(
	ctx context.Context,
	transition bool,
	fn func(repositories repositories.AllRepositories) error,
) error {

	if transition {

		tx, err := u.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		repositories := repositories.NewAllRepositories(tx)

		if err := fn(*repositories); err != nil {
			_ = tx.Rollback()
			return err
		}

		return tx.Commit()
	} else {
		repositories := repositories.NewAllRepositories(u.db)

		if err := fn(*repositories); err != nil {
			return err
		}
		return nil
	}
}

type UnitOfWork interface {
	Do(
		ctx context.Context,
		transition bool,
		fn func(repositories repositories.AllRepositories) error,
	) error
}
