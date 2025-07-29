-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_certifications (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    psychologist_profile_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    issued_by VARCHAR(255),
    issue_date DATE,
    description TEXT,
    admin_approved BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    CONSTRAINT fk_psychologist_profile FOREIGN KEY (psychologist_profile_id) REFERENCES t_psychologist_profile(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_certifications;
-- +goose StatementEnd
