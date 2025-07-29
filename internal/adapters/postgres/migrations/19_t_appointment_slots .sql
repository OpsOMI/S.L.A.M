-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_appointment_slots (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    psychologist_id UUID NOT NULL,
    day_of_week INT NOT NULL CHECK (day_of_week >= 0 AND day_of_week <= 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_psychologist FOREIGN KEY (psychologist_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT uq_psychologist_day_time UNIQUE (psychologist_id, day_of_week, start_time, end_time)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_appointment_slots;
-- +goose StatementEnd
