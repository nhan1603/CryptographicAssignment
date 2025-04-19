package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/menu"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/order"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/user"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Registry interface {
	User() user.Repository
	Menu() menu.Repository
	Order() order.Repository
	DoInTx(ctx context.Context, txFunc TxFunc) error
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		user:   user.New(pgConn),
		menu:   menu.New(pgConn),
		order:  order.New(pgConn),
		pgConn: pgConn,
	}
}

type impl struct {
	user   user.Repository
	menu   menu.Repository
	order  order.Repository
	txExec boil.Transactor
	pgConn *sql.DB
}

// TxFunc is a function that can be executed in a transaction
type TxFunc func(txRegistry Registry) error

// User returns user repo
func (i impl) User() user.Repository {
	return i.user
}

// Menu returns menu repo
func (i impl) Menu() menu.Repository {
	return i.menu
}

// Order returns order repo
func (i impl) Order() order.Repository {
	return i.order
}

// DoInTx handles db operations in a transaction
func (i impl) DoInTx(ctx context.Context, txFunc TxFunc) error {
	if i.txExec != nil {
		return errors.New("db tx nested in db tx")
	}

	tx, err := i.pgConn.BeginTx(ctx, nil)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	var committed bool
	defer func() {
		if committed {
			return
		}

		_ = tx.Rollback()
	}()

	newI := impl{
		user:   user.New(tx),
		menu:   menu.New(tx),
		order:  order.New(tx),
		txExec: tx,
	}

	if err = txFunc(newI); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return pkgerrors.WithStack(err)
	}

	committed = true

	return nil
}
