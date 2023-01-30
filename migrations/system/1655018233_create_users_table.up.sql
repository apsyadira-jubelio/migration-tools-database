-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE TABLE IF NOT EXISTS users (
    user_id serial NOT NULL primary key,
    name varchar(50) NOT NULL,
    email varchar(50) NOT NULL ,
    password varchar(255),
    schema_name varchar(255) not null,
    picture varchar(255) NOT NULL,
    company_id BIGINT NOT NULL,
    tenant_id varchar(255) NOT NULL,
    created_date timestamptz default now(),
    updated_date timestamptz
);

create unique index users_user_id_uindex
    on users (user_id);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS users;