package db

import (
	"context"

	uuid2 "github.com/google/uuid"
)

func (store *SqlStore) DeleteOutTimeOrder(uuid uuid2.UUID) error {
	return store.execTx(context.Background(), func(queries *Queries) error {
		err := queries.DeleteOrderByUUID(context.Background(), uuid)
		if err != nil {
			return err
		}
		return nil
	})
}
