-- name: CreateComment :one
insert into comment (content, movie_id, user_id, score, ip_address)
values ($1, $2, $3, $4, $5)
returning *;

-- name: DeleteCommentByID :exec
delete
from comment
where id = $1;

-- name: GetCommentsByUserID :many
select comment.id                                                                     as commentID,
       comment.content,
       comment.score,
       (select count(*) from comment_star where comment_star.comment_id = comment.id) as star_num,
       ip_address,
       created_at,
       movie_id,
       movie.name                                                                     as movieName,
       movie.avatar                                                                   as movieAvatar,
       count(*) over ()                                                               as total
from comment,
     movie
where comment.user_id = $1
  and movie_id = movie.id
order by created_at desc
limit $2 offset $3;

-- name: GetCommentByID :one
select comment.id                                                                     as commentID,
       content,
       score,
       (select count(*) from comment_star where comment_star.comment_id = comment.id) as star_num,
       ip_address,
       created_at,
       "user".id                                                                      as userID,
       "user".username,
       "user".avatar
from comment,
     "user"
where comment.id = $1
  and "user".id = comment.user_id
limit 1;

-- name: GetCommentsByMovieID :many
select comment.id                                                                     as commentID,
       content,
       score,
       (select count(*) from comment_star where comment_star.comment_id = comment.id) as star_num,
       ip_address,
       created_at,
       "user".id                                                                      as userID,
       "user".username,
       "user".avatar,
       count(*) over ()                                                               as total
from comment,
     "user"
where movie_id = $1
  and "user".id = comment.user_id
order by star_num desc, comment.id desc
limit $2 offset $3;

-- name: ExistComment :one
select exists(
               select 1
               from comment
               where user_id = $1
                 and movie_id = $2
           );
