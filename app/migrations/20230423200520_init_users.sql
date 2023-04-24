-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table users
(
    id       bigserial
        primary key,
    role     bigint not null
        references roles (role_id),
    username text   not null
        unique,
    password text   not null
);

alter table users
    owner to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users cascade ;
-- +goose StatementEnd
