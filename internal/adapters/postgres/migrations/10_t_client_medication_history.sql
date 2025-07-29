-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_client_medication_history (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    client_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    dosage VARCHAR(100),        
    note TEXT,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES t_users_auth(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_client_medication_history;
-- +goose StatementEnd
