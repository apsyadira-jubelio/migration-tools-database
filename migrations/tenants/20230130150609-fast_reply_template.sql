
-- +migrate Up
DROP TYPE IF EXISTS share_template_to;
CREATE TYPE share_template_to AS ENUM ('me', 'group');

CREATE TABLE if not exists fast_reply_template (
  fr_id serial primary key not null,
  category_id bigint,
  group_id bigint,
  content text not null,
  shortcut_code varchar(50) not null,
  images text,
  share_to share_template_to null,
  created_at timestamptz default now(),
  updated_at timestamptz,
  created_by bigint not null,
  updated_by bigint,
  CONSTRAINT fk_fast_reply_group_id FOREIGN KEY (group_id) REFERENCES groups (group_id),
  CONSTRAINT fk_fast_reply_category_id FOREIGN KEY (category_id) REFERENCES category_chat (category_id)
);

CREATE UNIQUE INDEX idx_fast_reply_template_shortcut_code ON fast_reply_template (shortcut_code);

-- +migrate Down
DROP TYPE IF EXISTS share_template_to;
DROP UNIQUE INDEX IF EXISTS idx_fast_reply_template_shortcut_code;
DROP TABLE IF EXISTS fast_reply_template;