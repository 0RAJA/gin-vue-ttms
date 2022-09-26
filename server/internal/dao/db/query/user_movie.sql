-- name: CreateUserMovie :exec
insert into user_movie (user_id, movie_id)
VALUES ($1, $2);

-- name: DeleteUserMovie :exec
delete
from user_movie
where user_id = $1
  and movie_id = $2;

-- name: ExistUserMovie :one
select exists(
               select 1
               from user_movie
               where user_id = $1
                 and movie_id = $2
           );
