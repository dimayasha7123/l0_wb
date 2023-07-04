-- +goose Up
-- +goose StatementBegin
CREATE TABLE models(
    order_uid varchar primary key,
    track_number varchar,
    entry varchar,
    delivery jsonb,
    payment jsonb,
    items jsonb,
    locale varchar,
    internal_signature varchar,
    customer_id varchar,
    delivery_service varchar,
    shardkey varchar,
    sm_id bigint,
    date_created timestamp,
    oof_shard varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE models;
-- +goose StatementEnd
