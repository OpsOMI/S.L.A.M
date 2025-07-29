-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_session_payments (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    paid_at TIMESTAMPTZ,
    paid_amount NUMERIC,
    payment_method VARCHAR(100) NOT NULL CHECK (payment_method IN ('cash', 'credit_card', 'paypal', 'bank_transfer')),
    payment_note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES t_sessions(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_session_payments;
-- +goose StatementEnd
