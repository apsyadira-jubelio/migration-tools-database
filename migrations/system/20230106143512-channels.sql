
-- +migrate Up
create table if not exists channels(
  id serial not null primary key,
  channel_name varchar(50) not null,
  channel_id bigint not null
);

-- +migrate Down
drop table if exists channels;