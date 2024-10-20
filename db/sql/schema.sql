create table tasks (
    id          serial       primary key,
    title       varchar(255) not null,
    description text,
    status      varchar(50)  not null     default 'pending',
    created_at  timestamp                 default CURRENT_TIMESTAMP,
    updated_at  timestamp                 default CURRENT_TIMESTAMP
);