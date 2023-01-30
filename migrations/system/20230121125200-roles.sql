
-- +migrate Up
create table if not exists roles(
  role_id serial not null primary key,
  role_name varchar(50) not null,
  created_at timestamptz default now(),
  updated_at timestamptz
);

insert into roles (role_name) values ('Super Admin');
insert into roles (role_name) values ('Account Admin');
insert into roles (role_name) values ('Admin');
insert into roles (role_name) values ('Agent');

-- +migrate Down
drop table if exists roles;