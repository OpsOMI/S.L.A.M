-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_role_permissions (
    role_id UUID NOT NULL REFERENCES t_roles(id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES t_permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_role_permissions;
-- +goose StatementEnd
