
-- +migrate Up
CREATE TABLE if not exists agent_groups (
  user_group_id serial primary key not null,
  user_id bigint,
  group_id bigint,
  created_at timestamptz default now(),
    CONSTRAINT fk_ug_group_id FOREIGN KEY (group_id) REFERENCES groups (group_id) on delete CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS agent_groups;