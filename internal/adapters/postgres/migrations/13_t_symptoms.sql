-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_symptoms (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    severity_level INTEGER,
    common_age_group VARCHAR(50),
    is_physical BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    CONSTRAINT fk_language FOREIGN KEY (language_id) REFERENCES t_languages(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_symptoms;
-- +goose StatementEnd
