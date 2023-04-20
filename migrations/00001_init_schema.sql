-- +goose Up
-- +goose StatementBegin
create table if not exists entity
(
    entity_uuid    uuid not null primary key,
    go_time        timestamp,
    go_time_tz     timestamptz,
    go_time_utc    timestamp,
    go_time_utc_tz timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists entity;
-- +goose StatementEnd
