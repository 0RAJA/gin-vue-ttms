// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: copyfrom.go

package db

import (
	"context"
)

// iteratorForCreateSeats implements pgx.CopyFromSource.
type iteratorForCreateSeats struct {
	rows                 []*CreateSeatsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateSeats) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateSeats) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].CinemaID,
		r.rows[0].Row,
		r.rows[0].Col,
	}, nil
}

func (r iteratorForCreateSeats) Err() error {
	return nil
}

func (q *Queries) CreateSeats(ctx context.Context, arg []*CreateSeatsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"ttms", "public", "seats"}, []string{"cinema_id", "row", "col"}, &iteratorForCreateSeats{rows: arg})
}

// iteratorForCreateTag implements pgx.CopyFromSource.
type iteratorForCreateTag struct {
	rows                 []*CreateTagParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateTag) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateTag) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].MovieID,
		r.rows[0].TagName,
	}, nil
}

func (r iteratorForCreateTag) Err() error {
	return nil
}

func (q *Queries) CreateTag(ctx context.Context, arg []*CreateTagParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"ttms", "public", "tags"}, []string{"movie_id", "tag_name"}, &iteratorForCreateTag{rows: arg})
}

// iteratorForCreateTickets implements pgx.CopyFromSource.
type iteratorForCreateTickets struct {
	rows                 []*CreateTicketsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateTickets) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateTickets) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].PlanID,
		r.rows[0].SeatsID,
		r.rows[0].Price,
	}, nil
}

func (r iteratorForCreateTickets) Err() error {
	return nil
}

func (q *Queries) CreateTickets(ctx context.Context, arg []*CreateTicketsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"ttms", "public", "tickets"}, []string{"plan_id", "seats_id", "price"}, &iteratorForCreateTickets{rows: arg})
}