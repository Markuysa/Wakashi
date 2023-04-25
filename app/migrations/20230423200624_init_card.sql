-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table card
(
    card_number bigint not null
        primary key,
    bank_id     bigint not null
        references bank,
    card_limit  bigint,
    owner_id    bigint
        references users,
    cvv_code    text   not null,
    total       double precision default 0
);
alter table card
    owner to postgres;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE card cascade;
-- +goose StatementEnd
