-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users_auth (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(40) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    email_token VARCHAR(120),
    password_reset_token VARCHAR(120) DEFAULT NULL,
    password_reset_expiry TIMESTAMPTZ DEFAULT NULL,
    first_login BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users_auth;
-- +goose StatementEnd
