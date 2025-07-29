-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_client_feedback (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL,
    client_id UUID NOT NULL,
    rating INT NOT NULL,
    comment TEXT,
    is_anonymous BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES t_sessions(id) ON DELETE CASCADE,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES t_users_auth(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_client_feedback;
-- +goose StatementEnd
