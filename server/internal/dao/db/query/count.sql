-- name: CreateVisitCount :exec
insert into visit_count (visit_count)
values ($1);

-- name: GetVisitCountsByCreateDate :one
select sum(visit_count.visit_count)
from visit_count
where create_date >= @startTime
  and create_date <= @endTime;
