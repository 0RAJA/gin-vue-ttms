-- name: CreateOrder :exec
insert into ttms.public."order"(user_id, order_id, movie_name, movie_avatar, cinema_name, create_at, seats, price, status,plan_id,seats_id)
values (@user_id,@order_id,@movie_name,@movie_avatar,@cinema_name,@create_at,@seats,@price,@status,@plan_id,@seats_id);

-- name: DeleteOrderByUUID :exec
delete from ttms.public."order"
where order_id = @order_id;

-- name: GetOrderByUserId :many
select * from
ttms.public."order"
where user_id = @user_id;

-- name: UpdateOrderStatus :exec
update ttms.public."order"
set
status = '已支付'
where order_id = @order_id;

-- name: DeleteOrderByTicket :exec
delete
from ttms.public."order"
where plan_id = @plan_id
and seats_id = @seats_id;

-- name: GetWaitPayOrder :many
select * from
ttms.public."order"
where status = '待支付';

-- name: SearchAllOrder :many
select * from
ttms.public."order"
order by user_id desc
limit $1 offset $2;

-- name: GetNumsAll :one
select count(*) from
ttms.public."order";

-- name: SearchOrderByCondition :many
select order_id,user_id,movie_name, movie_avatar, cinema_name, create_at, seats, price, status,plan_id,seats_id,count(*) over () as total
from ttms.public."order"
where movie_name like $1
   or cinema_name like $1
order by user_id desc
limit $2 offset $3;