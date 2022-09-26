-- name: GetByPlan :many
select *
from ttms.public."tickets"
where plan_id = @planId;

-- name: GetAllTickets :many
select *
from ttms.public."tickets"
order by plan_id desc
limit $1 offset $2;

-- name: GetTicketNum :one
select count(*)
from ttms.public."tickets";

-- name: GetTicketsLocked :many
select *
from ttms.public."tickets"
where plan_id = @planId
  and status = '锁定';

-- name: GetTicket :one
select *
from ttms.public."tickets"
where plan_id = @planId
  and seats_id = @seatID;

-- name: CreateTickets :copyfrom
insert into ttms.public."tickets" (plan_id, seats_id, price)
values ($1, $2, $3);

-- name: CreateTicket :exec
insert into ttms.public."tickets"
    (plan_id, seats_id, price)
values (@planId, @seatID, @price);

-- name: DeleteByPlan :exec
delete
from ttms.public."tickets"
where plan_id = @planId;

-- name: DeleteBySeats :exec
delete
from ttms.public."tickets"
where seats_id = @seatId
  and plan_id = @planId;

-- name: LockTicket :exec
update ttms.public.tickets
set status  = @status,
    user_id = @user_id
where plan_id = @plan_id
  and seats_id = @seats_id;

-- name: PayTicket :exec
update ttms.public.tickets
set status  = @status,
    user_id = @user_id
where plan_id = @plan_id
  and seats_id = @seats_id;

-- name: GetTicketsByPlan :many
select t.plan_id, t.seats_id, s.status, t.price, t.status, s.row, s.col
from ttms.public.seats s,
     ttms.public.tickets t
where t.plan_id = @plan_id
  and t.seats_id = s.id;

-- 查询演出计划是否有已经售出或者锁定的票
-- name: ExistSoldTicketsByPlan :one
select exists(select 1
              from tickets,
                   plan
              where tickets.plan_id = @planId
                and plan.id = tickets.plan_id
                and (tickets.status = '已售'
                  or tickets.status = '锁定')
                and plan.end_at > now());

-- name: SearchTicketByPlanId :many
select *
from ttms.public.tickets
where plan_id = $1
order by plan_id desc
limit $2 offset $3;

-- name: QueryCountTicketPlan :one
select count(*)
from ttms.public.tickets
where plan_id = @plan_id;

-- name: UnLockTicket :exec
update ttms.public.tickets
set status = '正常'
where plan_id = @plan_id
  and user_id = @user_id
  and seats_id = @seats_id;

