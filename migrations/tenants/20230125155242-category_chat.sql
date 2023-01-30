
-- +migrate Up
CREATE TABLE IF NOT EXISTS category_chat(
  category_id serial primary key not null,
  category_name varchar(50) not null,
  is_active boolean not null default false,
  created_at timestamptz default now(),
  updated_at timestamptz
);

-- +migrate Down
DROP TABLE IF EXISTS category_chat;