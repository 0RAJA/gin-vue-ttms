package db

import (
	"context"
	"fmt"

	uuid2 "github.com/google/uuid"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	Querier
	TXer
}

type TXer interface {
	CreatePlanWithTx(ctx context.Context, arg *CreatePlanWithTxParams) (*CreatePlanWithTxRely, error)
	DeletePlanWithTx(ctx context.Context, id int64) error
	CreateMovieWithTx(ctx context.Context, params *CreateMovieWithTxParams) (*CreateMovieWithTx, error)
	GetMovieByIDWithTx(ctx context.Context, userID, movieID int64) (*GetMovieByIDWithTxRow, error)
	DeleteMovieByIDWithTx(ctx context.Context, cinemaID int64) error
	CreateCinemaWithTx(ctx context.Context, arg *CreateCinemaParams) (*Cinema, error)
	DeleteCinemaByIDWithTx(ctx context.Context, cinemaID int64) error
	LockTicketsTx(c context.Context, params *LockTicketsTxParams) (*LockTicketsTx, error)
	PayTicketTx(ctx context.Context, params *PayTicketTXParams) error
	DeleteOutTimeTicket(planId, userId, seatId int64) error
	DeleteOutTimeOrder(uuid uuid2.UUID) error
}

type SqlStore struct {
	*Queries
	DB *pgxpool.Pool
}

// 通过事务执行回调函数
func (store *SqlStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.Deferrable,
	})
	if err != nil {
		return err
	}
	q := store.WithTx(tx) // 使用开启的事务创建一个查询
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
