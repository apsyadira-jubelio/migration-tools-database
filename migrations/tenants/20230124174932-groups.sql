
-- +migrate Up
CREATE TABLE if not exists groups (
  group_id serial primary key not null,
  group_name varchar(50) not null,
  description text,
  created_at timestamptz default now(),
  updated_at timestamptz
);

-- +migrate Down
DROP TABLE IF EXISTS groups;