-- name: GetSeatsByCinemas :many
select *
from ttms.public."seats"
where cinema_id = @cinema_id;

-- name: GetSeatsById :one
select * from
ttms.public."seats"
where id = @id;

-- name: GetOne :one
select *
from ttms.public."seats"
where row = @row
  and col = @col;

-- name: CreateSeats :copyfrom
insert into ttms.public."seats"
    (cinema_id, row, col)
values ($1, $2, $3);

-- name: CreateSeat :exec
insert into ttms.public."seats"
    (cinema_id, row, col)
values (@cinemaId, @row, @col);

-- name: UpdateSeats :exec
update ttms.public."seats"
set status = @status
where cinema_id = @cinemaId
  and row = @row
  and col = @col;

-- name: UpdateSeatsById :exec
update ttms.public."seats"
set status = @status
where id = @id;

-- name: DeleteSeatsByPlan :exec
delete
from ttms.public."seats"
where cinema_id = @cinemaId;

-- name: DeleteSeat :exec
delete
from ttms.public."seats"
where cinema_id = @cinemaId
  and col = @col
  and row = @row;

