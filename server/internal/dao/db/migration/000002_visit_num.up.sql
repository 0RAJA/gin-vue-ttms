create table visit_count
(
    create_date timestamptz not null default now(),
    visit_count bigint      not null,
    primary key (create_date)
);
