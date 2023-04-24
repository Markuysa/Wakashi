-- +goose Up
-- +goose StatementBegin
-- auto-generated definition
create table relation
(
    master_id bigint
        references users,
    slave_id  bigint
        constraint relation_users_id_fk
            references users
);

alter table relation
    owner to postgres;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE relation cascade;
-- +goose StatementEnd
