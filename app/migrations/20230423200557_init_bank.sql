-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table bank
(
    uin       bigint not null
        primary key,
    bank_name text
);

alter table bank
    owner to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table bank cascade;
-- +goose StatementEnd
