
-- +migrate Up
CREATE TABLE if not exists users_link(
  user_link_id serial primary key not null,
  user_id bigint,
  email varchar(50) not null
);

-- +migrate Down
DROP TABLE IF EXISTS users_link;