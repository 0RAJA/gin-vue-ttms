create type OrderStatus As ENUM ('待支付','已支付');
create table "order"(
                        user_id bigint not null references "user" (id) on DELETE cascade,
                        plan_id bigint not null,
                        seats_id varchar(255) not null,
                        order_id uuid   not null,
                        movie_name varchar(255) not null,
                        movie_avatar varchar(255) not null,
                        cinema_name varchar(255) not null,
                        create_at timestamptz not null default now(),
                        seats     varchar(255) not null,
                        price     real      not null,
                        status    OrderStatus not null default '待支付'
);

