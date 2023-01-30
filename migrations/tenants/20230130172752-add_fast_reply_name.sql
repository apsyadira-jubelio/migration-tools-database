
-- +migrate Up
alter table fast_reply_template
add column if not exists template_name varchar(50) not null;

-- +migrate Down
alter table fast_reply_template
drop column if exists template_name;