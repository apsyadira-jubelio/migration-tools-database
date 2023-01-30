
-- +migrate Up
ALTER TABLE users
ADD COLUMN role_id bigint NOT NULL DEFAULT 1,
ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(role_id);


-- +migrate Down
ALTER TABLE users
DROP COLUMN role_id,
DROP CONSTRAINT fk_role_id;