
-- +migrate Up

-- TENANT_CHANNEL TABLE
CREATE TABLE IF NOT EXISTS tenant_channel (
  tenant_channel_id serial NOT NULL primary key,
  channel_id bigint not null,
  app_id varchar(255),
  channel_user_id varchar(255),
  channel_user_secret varchar(255) NOT NULL,
  store_id varchar(50),
  channel_type varchar(50),
  is_active boolean default true,
  extra_info jsonb,
  channel_tag varchar(25),
  created_at timestamptz default now(),
  updated_at timestamptz,
  created_by bigint NOT NULL,
  updated_by bigint
);

-- +migrate Down
DROP TABLE IF EXISTS tenant_channel;