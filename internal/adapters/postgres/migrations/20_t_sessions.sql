-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_sessions (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    client_id UUID NOT NULL,
    psychologist_id UUID NOT NULL,
    session_date DATE NOT NULL,
    session_start_time TIME NOT NULL,
    session_end_time TIME NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'scheduled',
    notes TEXT,
    future_notes TEXT,
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES t_users_auth(id),
    CONSTRAINT fk_psychologist FOREIGN KEY (psychologist_id) REFERENCES t_users_auth(id),

    CONSTRAINT chk_status CHECK (status IN ('scheduled', 'completed', 'cancelled', 'no_show')),
    CONSTRAINT uq_psychologist_session UNIQUE (psychologist_id, session_date, session_start_time)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_sessions;
-- +goose StatementEnd