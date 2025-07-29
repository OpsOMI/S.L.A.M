-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_psychologist_profile (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_auth_id UUID NOT NULL UNIQUE,
    title VARCHAR(30),
    about_me TEXT,
    session_fee NUMERIC,
    years_of_experience INTEGER,
    profile_picture_path TEXT,
    rating INTEGER DEFAULT 0,
    session_duration_minutes INTEGER NOT NULL DEFAULT 50,
    notes_secret_key VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_psychologist_profile;
-- +goose StatementEnd
