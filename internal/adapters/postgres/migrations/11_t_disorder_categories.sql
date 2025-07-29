-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_disorder_categories (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    CONSTRAINT fk_language FOREIGN KEY (language_id) REFERENCES t_languages(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_disorder_categories;
-- +goose StatementEnd
