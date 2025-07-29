-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_disorder_symptoms (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    disorder_id UUID NOT NULL,
    symptom_id UUID NOT NULL,
    CONSTRAINT fk_disorder FOREIGN KEY (disorder_id) REFERENCES t_disorders(id) ON DELETE CASCADE,
    CONSTRAINT fk_symptom FOREIGN KEY (symptom_id) REFERENCES t_symptoms(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_disorder_symptoms;
-- +goose StatementEnd
