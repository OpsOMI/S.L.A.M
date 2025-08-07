-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = @id;

-- name: GetRoomByCode :one
SELECT * FROM rooms WHERE code = @code;

-- name: GetRoomsByOwnerID :many
SELECT * FROM rooms WHERE owner_id = @owner_id ORDER BY created_at DESC LIMIT @lim OFFSET @off;

-- name: CreateRoom :one
INSERT INTO rooms (
    owner_id,
    code,
    password
) VALUES (
    @owner_id,
    @code,
    @password
)
RETURNING id;

-- name: DeleteRoomByID :exec
DELETE FROM rooms WHERE id = @id;

-- name: CountRoomsByOwnerID :one
SELECT COUNT(*) FROM rooms WHERE owner_id = @owner_id;

-- name: IsRoomExistByID :one
SELECT EXISTS (
    SELECT 1 FROM rooms WHERE id = @id
);

-- name: IsRoomExistByCode :one
SELECT EXISTS (
    SELECT 1 FROM rooms WHERE code = @code
);

-- name: IsRoomExistByOwnerID :one
SELECT EXISTS (
    SELECT 1 FROM rooms WHERE owner_id = @owner_id
);
