
-- +migrate Up

-- TENANT_CHANNEL TABLE
CREATE TABLE IF NOT EXISTS tenant_channel (
  tenant_channel_id serial NOT NULL primary key,
  tenant_id varchar(255) NOT NULL,
  user_id bigint not null,
  channel_id BIGINT not null,
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
  created_by varchar(50) NOT NULL,
  updated_by varchar(50),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- +migrate Down
DROP TABLE IF EXISTS tenant_channel;