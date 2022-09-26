-- name: CreateCommentStar :exec
insert into comment_star (user_id, comment_id)
values ($1, $2);

-- name: GetCommentStar :one
select *
from comment_star
where user_id = $1
  and comment_id = $2
limit 1;

-- name: DeleteCommentStar :exec
delete
from comment_star
where user_id = $1
  and comment_id = $2;
