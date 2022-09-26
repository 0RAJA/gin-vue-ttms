-- name: CreateCinema :one
insert into cinema (name, avatar, rows, cols)
values ($1, $2, $3, $4)
returning *;

-- name: DeleteCinemaByID :exec
delete
from cinema
where id = $1;

-- name: GetCinemaByID :one
select *
from cinema
where id = $1
limit 1;

-- name: GetCinemas :many
select *, count(*) over () as total
from cinema
order by id desc
limit $1 offset $2;

-- name: CheckCinemaByName :one
select exists(
               select 1
               from cinema
               where name = $1
           );

-- name: UpdateCinema :one
update cinema
set name   = $2,
    avatar = $3
where id = $1
returning *;

-- name: GetCinemaByPlanID :one
select cinema.*
from plan,
     cinema
where plan.id = $1
  and plan.cinema_id = cinema.id
  and plan.end_at > now()
limit 1;

-- name: GetOrderInfoByCinemaId :one
select  m.avatar,m.name,c.name
from ttms.public.movie m ,
     ttms.public.cinema c
where m.id = @movie_id
and c.id = @cinema_id;
