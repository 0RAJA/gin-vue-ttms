-- name: CreateMovie :one
insert into movie (name, director, alias_name, actors, content, avatar, duration, area, period)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning id,name,director, alias_name, actors, content, avatar, duration, area, period;

-- name: GetMovieByID :one
select movie.id,
       name,
       director,
       movie.alias_name,
       actors,
       movie.content,
       avatar,
       visit_count,
       box_office,
       duration,
       period,
       movie.score,
       area,
       (select count(*)
        from comment
        where comment.movie_id = $1)          as comment_count,
       (select count(*)
        from user_movie
        where user_movie.movie_id = movie.id) as follow_count
from movie
where id = $1
limit 1;

-- name: DeleteMovieByID :exec
delete
from movie
where id = $1;

-- name: UpdateMovie :one
update movie
set name       = $1,
    alias_name = $2,
    actors     = $3,
    content    = $4,
    avatar     = $5,
    period     = $6,
    area       = $7,
    director   = $8
where id = $9
returning id,name,alias_name,actors,content,avatar,period,area;

-- name: AddMovieVisitCount :exec
update movie
set visit_count = visit_count + $2
where id = $1;

-- name: AddMovieBoxOffice :exec
update movie
set box_office = box_office + $2
where id = $1;

-- name: GetMoviesByTagPeriodAreaOrderByVisitCount :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       count(*) over () as total
from movie,
     (select distinct movie_id from tags where tags.tag_name like $2) as tags
where period between @startTime and @endTime
  and area like $1
  and id = tags.movie_id
order by visit_count desc
limit $3 offset $4;

-- name: GetMoviesByTagPeriodAreaOrderByScore :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       count(*) over () as total
from movie,
     (select distinct movie_id from tags where tags.tag_name like $2) as tags
where period between @startTime and @endTime
  and area like $1
  and id = tags.movie_id
order by score desc
limit $3 offset $4;

-- name: GetMoviesByTagPeriodAreaOrderByPeriod :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       count(*) over () as total
from movie,
     (select distinct movie_id from tags where tags.tag_name like $2) as tags
where period between @startTime and @endTime
  and area like $1
  and id = tags.movie_id
order by period desc
limit $3 offset $4;

-- name: GetMoviesByNameOrContent :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       count(*) over () as total
from movie
where movie.name = @key
   or movie.alias_name = @key
   or movie.name_content_alias_tsv @@ plainto_tsquery(@key::varchar)
limit $1 offset $2;

-- name: GetAreas :many
select distinct area
from movie
order by area
limit $1 offset $2;

-- name: GetMoviesByIDs :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score
from movie
where id = any (@ids::bigint[])
order by array_positions(@ids::bigint[], id::bigint);

-- name: GetMoviesOrderByVisitCount :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       count(*) over () as total
from movie
order by visit_count desc
limit $1 offset $2;

-- name: GetMoviesOrderByBoxOffice :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       box_office
from movie
order by box_office desc
limit $1;

-- name: GetMoviesOrderByUserMovieCount :many
select id,
       actors,
       name,
       alias_name,
       avatar,
       period,
       score,
       movieIDs.fallow_count
from movie,
     (select movie_id, count(*) as fallow_count
      from user_movie
      group by movie_id
      order by fallow_count desc
      limit $1) as movieIDs
where movieIDs.movie_id = movie.id
order by fallow_count desc;

-- name: GetMovies :many
select id,
       name,
       director,
       alias_name,
       actors,
       content,
       avatar,
       duration,
       area,
       period,
       count(*) over () as total
from movie
order by id desc
limit $1 offset $2;
