-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_expertise (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    psychologist_profile_id UUID NOT NULL,
    disorder_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    UNIQUE (psychologist_profile_id, disorder_id),
    CONSTRAINT fk_psychologist_profile FOREIGN KEY (psychologist_profile_id) REFERENCES t_psychologist_profile(id) ON DELETE CASCADE,
    CONSTRAINT fk_disorder FOREIGN KEY (disorder_id) REFERENCES t_disorders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_psychologist_expertise;
-- +goose StatementEnd
