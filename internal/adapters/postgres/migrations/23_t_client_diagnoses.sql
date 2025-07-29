-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_client_diagnoses (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    client_id UUID NOT NULL,
    session_id UUID NOT NULL,
    disorder_id UUID NOT NULL,
    diagnosis_date DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES t_sessions(id) ON DELETE CASCADE,
    CONSTRAINT fk_disorder FOREIGN KEY (disorder_id) REFERENCES t_disorders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_client_diagnoses;
-- +goose StatementEnd
