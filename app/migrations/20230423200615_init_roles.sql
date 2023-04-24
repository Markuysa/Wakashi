-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table roles
(
    role_id bigserial
        constraint unique_id
            unique,
    role    text not null
);

alter table roles
    owner to postgres;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles cascade;
-- +goose StatementEnd
