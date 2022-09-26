drop table if exists tickets cascade;

drop table if exists seats cascade;

drop table if exists plan cascade;

drop table if exists cinema cascade;

drop table if exists tags cascade;

drop table if exists comment_star cascade;

drop table if exists comment cascade;

drop table if exists "user" cascade;

drop table if exists movie cascade;

drop table if exists user_movie cascade;

drop type if exists gender cascade;

drop type if exists seatsstatus cascade;

drop type if exists ticketstatus cascade;

drop type if exists orderby cascade;

drop type if exists OrderStatus cascade;

drop type if exists method cascade;

drop type if exists privilege cascade;

drop type if exists lifestate cascade;

drop function if exists update_ticket_lock_time() cascade;

drop function if exists update_movie_score_up() cascade;
drop function if exists update_movie_score_down() cascade;

drop function if exists box_office_update() cascade;

drop text search configuration chinese;
