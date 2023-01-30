
-- +migrate Up
alter table if exists users
add column if not EXISTS is_owner boolean default false;

-- +migrate Down
alter table if exists users
drop column is EXISTS is_owner;