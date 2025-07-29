-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users_2fa (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_auth_id UUID NOT NULL,
    method TEXT NOT NULL DEFAULT 'email', -- 'totp', 'sms', 'email' etc.
    secret TEXT,          
    enabled BOOLEAN NOT NULL DEFAULT FALSE,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users_2fa;
-- +goose StatementEnd
