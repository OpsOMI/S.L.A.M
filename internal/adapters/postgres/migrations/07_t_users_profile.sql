-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users_profile (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_auth_id UUID NOT NULL,
    role_id UUID NOT NULL,
    nickname VARCHAR(30),
    name VARCHAR(30),
    surname VARCHAR(30),
    phone VARCHAR(20),
    birth_date TIMESTAMP,
    gender VARCHAR(1),
    score INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES t_roles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users_profile;
-- +goose StatementEnd
