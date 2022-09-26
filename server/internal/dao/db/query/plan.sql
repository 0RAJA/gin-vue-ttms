-- name: CreatePlan :one
insert into plan(movie_id, version, cinema_id, start_at, end_at, price)
VALUES ($1, $2, $3, $4, $5, $6)
returning id;

-- name: GetAllPlanIds :many
select id, end_at
from ttms.public.plan
where plan.end_at > now();

-- name: DeletePlan :exec
delete
from plan
where id = $1;

-- name: GetPlanByID :one
select plan.id,
       plan.version,
       movie_id,
       cinema_id,
       cinema.name as cinemaName,
       start_at,
       end_at,
       price
from plan,
     cinema
where plan.id = $1
  and cinema.id = plan.cinema_id
limit 1;

-- name: GetPlansByMovieAndStartTimeOrderByPrice :many
select plan.id,
       plan.version,
       movie_id,
       cinema_id,
       cinema.name      as cinemaName,
       start_at,
       end_at,
       price,
       count(*) over () as total
from plan,
     cinema
where movie_id = $1
  and cinema.id = plan.cinema_id
  and start_at between @startTime and @endTime
  and plan.end_at > now()
order by price
limit $2 offset $3;

-- name: GetPlansCountByTimeWithLock :one
select exists(select 1
              from plan pl
              where pl.cinema_id = $3
                and (pl.start_at <= $1
                         and pl.end_at >= $2
                  or pl.start_at >= $1
                         and pl.end_at <= $2
                  or pl.start_at between $1 and $2
                  or pl.end_at between $1 and $2)
                  for update
           );

-- name: DeleteOutDatePlans :many
delete
from plan
where end_at < now()
returning id;


-- name: ExistPlansByMovieID :one
select exists(select 1
              from plan
              where movie_id = $1
                and plan.end_at > now());

-- name: ExistPlansByCinemaID :one
select exists(select 1
              from plan
              where cinema_id = $1
                and plan.end_at > now());

-- name: GetPlans :many
select plan.id,
       plan.version,
       movie_id,
       movie.name,
       cinema_id,
       cinema.name      as cinemaName,
       start_at,
       end_at,
       price,
       count(*) over () as total
from plan,
     movie,
     cinema
where plan.movie_id = movie.id
  and plan.cinema_id = cinema.id
  and plan.end_at > now()
order by id desc
limit $1 offset $2;

-- name: GetPlansByMovie :many
select plan.id,
       plan.version,
       movie_id,
       movie.name,
       cinema_id,
       cinema.name      as cinemaName,
       start_at,
       end_at,
       price,
       count(*) over () as total
from plan,
     movie,
     cinema
where plan.movie_id = $1
  and plan.movie_id = movie.id
  and plan.cinema_id = cinema.id
  and plan.end_at > now()
order by id desc
limit $1 offset $2;
