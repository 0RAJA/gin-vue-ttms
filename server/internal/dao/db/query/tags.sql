-- name: GetTags :many
select distinct tags.tag_name from ttms.public."tags";


-- name: GetTagsInMovie :many
select tag_name from ttms.public."tags"
where movie_id = @MovieId::bigint;

-- name: GetMovieInTag :many
select * from ttms.public."tags"
where tag_name = @tag_name::varchar;

-- name: CreateTag :copyfrom
insert into ttms.public."tags"
(movie_id,tag_name)
values ($1, $2);

-- name: DeleteByMovieId :exec
delete from ttms.public."tags"
where movie_id = @movie_id::bigint;

-- name: DeleteOneByMovieAndTag :exec
delete from ttms.public."tags"
where movie_id = @movie_id::bigint
and tag_name = @tag_name::varchar;

-- name: DeleteByTagName :exec
delete from ttms.public."tags"
where tag_name = @tag_name::varchar;

-- name: UpdateMovieTag :exec
update ttms.public."tags"
set
tag_name = @newTagName::varchar
where movie_id = @movie_id::bigint
and @oldTagName::varchar = tag_name;



