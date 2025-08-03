-- name: GetClientByID :one
SELECT * FROM clients WHERE id = @id;

-- name: GetClientByClientKey :one
SELECT * FROM clients WHERE client_key = @client_key;

-- name: GetClientsByUserID :many
SELECT * FROM clients WHERE user_id = @user_id;

-- name: CreateClient :one
INSERT INTO clients (
    user_id,
    client_key
) VALUES (
    @user_id,
    @client_key
)
RETURNING id;

-- name: RevokeClient :exec
UPDATE clients
SET revoked_at = CURRENT_TIMESTAMP
WHERE id = @id;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = @id;

-- name: CountClientsByUserID :one
SELECT COUNT(*) FROM clients WHERE user_id = @user_id;

-- name: IsClientExistByID :one
SELECT EXISTS (
    SELECT 1 FROM clients WHERE id = @id
);

-- name: IsClientExistByClientKey :one
SELECT EXISTS (
    SELECT 1 FROM clients WHERE client_key = @client_key
);

-- name: IsClientRevoked :one
SELECT EXISTS (
    SELECT 1 FROM clients WHERE id = @id AND revoked_at IS NOT NULL
);