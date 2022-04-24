-- +goose Up

create table users (
    id bigserial primary key,
    name text not null,
    age bigint not null,
    email text not null
);

-- +goose Down

drop table users;