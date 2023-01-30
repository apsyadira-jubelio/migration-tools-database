
-- +migrate Up
alter table if exists users
add column is_active boolean default false;

-- +migrate Down
alter table if exists users
drop column is_active;
