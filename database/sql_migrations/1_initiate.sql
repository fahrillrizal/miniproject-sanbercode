-- +migrate Up
-- +migrate StatementBegin

create table person (
    id          SERIAL PRIMARY KEY,
    first_name  varchar(256),
    last_name   varchar(256)
)

-- +migrate StatementEnd
