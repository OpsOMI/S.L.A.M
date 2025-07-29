-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_permissions (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_permissions;
-- +goose StatementEnd
