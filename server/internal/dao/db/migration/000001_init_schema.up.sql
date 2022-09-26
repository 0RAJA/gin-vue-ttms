-- 使用中文分词
--连接到目标数据库，创建zhparser解析器
-- CREATE EXTENSION zhparser;
-- 将zhparser解析器作为全文检索配置项
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);
--普遍情况下，我们只需要按照名词(n)，动词(v)，形容词(a)，成语(i),叹词(e)和习用语(l)6种方式对句子进行划分就可以了，词典使用的是内置的simple词典，即仅做小写转换
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;

create type Gender As ENUM ('男','女','未知');
create type SeatsStatus As ENUM ('正常','损坏','走廊');
create type TicketStatus As ENUM ('正常','锁定','已售');
create type Method As ENUM ('POST','GET','PUT','DELETE');
create type OrderBy As ENUM ('visit_count','score','period');
create type Privilege As ENUM ('BAN','管理员','用户');
create type LifeState As ENUM ('单身','热恋','已婚','为人父母','未知');

create table "user"
(
    id        bigserial primary key,
    username  varchar(255) not null unique,
    password  varchar(255) not null,
    avatar    varchar(255) not null,
    lifeState LifeState             default '未知',
    hobby     varchar(255)          default null,
    email     varchar(255) not null unique,
    birthday  timestamptz  not null,
    gender    Gender       not null default '未知',
    signature text         not null default '日常摆烂',
    privilege Privilege    not null default '用户'
);

create table "movie"
(
    id                     bigserial primary key,
    actors                 varchar(128)[] not null,
    name                   varchar(255)   not null,
    alias_name             varchar(255)   not null,
    content                text           not null,
    avatar                 varchar(255)   not null,
    visit_count            bigint         not null default 0,
    box_office             real           not null default 0 check (box_office >= 0 ),
    duration               smallint       not null check (duration > 0),
    period                 timestamptz    not null,
    score                  real           not null default 0 check (score >= 0 and score <= 10),
    area                   varchar(255)   not null,
    name_content_alias_tsv tsvector
);
create index movie_period_index on movie (period desc);
create index movie_visit_count_index on movie (visit_count desc);
create index movie_name_index on movie (name);
create index movie_alias_name on movie (alias_name);
-- 分词查询索引
create index name_content_alias_idx on movie using gin (to_tsvector('chinese', name || alias_name || content));

-- 触发器更新 name_content_alias_tsv
CREATE TRIGGER trigger_name_content_alias_tsv
    BEFORE INSERT OR UPDATE
    ON movie
    FOR EACH ROW
EXECUTE PROCEDURE
    tsvector_update_trigger(name_content_alias_tsv, 'public.chinese', name, alias_name, content);

create table "user_movie"
(
    user_id  bigint not null references "user" (id) on DELETE CASCADE,
    movie_id bigint not null references "movie" (id) on DELETE CASCADE,
    primary key (user_id, movie_id)
);

create table "cinema"
(
    id     bigserial primary key,
    name   varchar(255) not null unique,
    avatar varchar(255) not null,
    rows   smallint     not null check (rows > 0),
    cols   smallint     not null check (cols > 0)
);

create table "comment"
(
    id         bigserial primary key,
    content    text         not null,
    movie_id   bigint       not null references "movie" (id) on DELETE CASCADE,
    user_id    bigint       not null references "user" (id) on DELETE CASCADE,
    score      real         not null check (score >= 0 and score <= 10),
    ip_address varchar(255) not null,
    created_at timestamptz  not null default now(),
    unique (user_id, movie_id)
);

create table "comment_star"
(
    user_id    bigint not null references "user" (id) on DELETE CASCADE,
    comment_id bigint not null references "comment" (id) on DELETE CASCADE,
    primary key (user_id, comment_id)
);

create table "seats"
(
    id        bigserial primary key,
    cinema_id bigint      not null references "cinema" (id) on DELETE CASCADE,
    row       smallint    not null check (row > 0),
    col       smallint    not null check (col > 0),
    status    SeatsStatus not null default '正常'
);

create table "plan"
(
    id        bigserial primary key,
    movie_id  bigint       not null references "movie" (id) on DELETE CASCADE,
    cinema_id bigint       not null references "cinema" (id) on DELETE CASCADE,
    version   varchar(255) not null,
    start_at  timestamptz  not null,
    end_at    timestamptz  not null,
    price     real         not null check (price >= 0)
);

create table "tickets"
(
    user_id   bigint       not null default 0,
    plan_id   bigint       not null references "plan" (id) on DELETE CASCADE,
    seats_id  bigint       not null references "seats" (id) on DELETE CASCADE,
    price     real         not null check (price >= 0),
    status    TicketStatus not null default '正常',
    lock_time timestamptz  not null default now(),
    constraint tickets_primary unique (plan_id, seats_id)
);

create table tags
(
    movie_id bigint       not null references "movie" (id) on DELETE CASCADE,
    tag_name varchar(255) not null,
    constraint tags_primary unique (movie_id, tag_name)
);

-- 更新票房信息
create or replace function box_office_update() returns trigger as
$$
begin
    if new.status = '已售' then
        update movie
        set box_office = box_office + new.price
        where movie.id = (select plan.movie_id from plan where plan.id = new.plan_id limit 1);
    end if;
    if new.status = '正常' then
        update movie
        set box_office = box_office - new.price
        where movie.id = (select plan.movie_id from plan where plan.id = new.plan_id limit 1);
    end if;
    return new;
end;
$$ language 'plpgsql';

create trigger "box_office_update"
    after update
    on "tickets"
    for each row
execute procedure box_office_update();

-- 更新电影评分
create or replace function update_movie_score_up() returns trigger as
$$
begin
    update "movie"
    set score = (select avg(comment.score) from "comment" where "comment"."movie_id" = new."movie_id")
    where "movie"."id" = new.movie_id;
    return new;
end;
$$ language 'plpgsql';

create trigger "update_movie_score_up"
    after insert or update
    on "comment"
    for each row
execute procedure update_movie_score_up();

create or replace function update_movie_score_down() returns trigger as
$$
begin
    if not exists(select id from "comment" where "comment"."movie_id" = old."movie_id") then
        update "movie"
        set "score" = 0
        where "movie"."id" = old."movie_id";
        return old;
    end if;
    update "movie"
    set score = (select avg(comment.score) from "comment" where "comment"."movie_id" = old."movie_id")
    where "movie"."id" = old.movie_id;
    return old;
end;
$$ language 'plpgsql';

create trigger "update_movie_score_down"
    after delete
    on "comment"
    for each row
execute procedure update_movie_score_down();

-- 更新票锁定时间
CREATE OR REPLACE FUNCTION update_ticket_lock_time()
    RETURNS TRIGGER AS
$$
BEGIN
    if NEW.status = '锁定' then
        NEW.lock_time = now();
    end if;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_ticket_lock_time
    BEFORE UPDATE of status
    ON "tickets"
    FOR EACH ROW
EXECUTE PROCEDURE update_ticket_lock_time();
