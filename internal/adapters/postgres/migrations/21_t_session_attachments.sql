-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_session_attachments (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL,
    file_url VARCHAR(255),
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(255),

    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES t_sessions(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_session_attachments;
-- +goose StatementEnd
