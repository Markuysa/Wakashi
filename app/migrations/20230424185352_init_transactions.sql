-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
-- auto-generated definition
create table transactions
(
    id               bigserial
        primary key,
    card_id          bigint                not null
        references card,
    owner_id         bigint
        references users,
    operation_value  double precision,
    transaction_date timestamp,
    status           boolean default false not null,
    request_from     integer
        constraint transactions_users__fk
            references users
);


alter table transactions
    owner to postgres;




-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table transactions;
-- +goose StatementEnd
