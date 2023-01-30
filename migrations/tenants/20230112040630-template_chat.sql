
-- +migrate Up

-- TEMPLATE_CHAT TABLE
CREATE TABLE IF NOT EXISTS templates_chat (
  template_id serial NOT NULL primary key,
  template_name varchar(50) not null,
  is_active boolean default false,
  created_at timestamptz default now(),
  updated_at timestamptz,
  created_by varchar(50) NOT NULL,
  updated_by varchar(50)
);

-- +migrate Down
DROP TABLE IF EXISTS templates_chat;
