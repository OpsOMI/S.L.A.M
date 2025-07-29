-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_disorders (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID NOT NULL,
    category_id UUID NOT NULL,
    name VARCHAR(30) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_language FOREIGN KEY (language_id) REFERENCES t_languages(id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES t_disorder_categories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_disorders;
-- +goose StatementEnd