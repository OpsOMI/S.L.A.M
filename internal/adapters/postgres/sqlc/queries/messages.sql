-- name: GetMessagesByReceiver :many
SELECT
    s.nickname AS sender_nickname,
    r.code AS room_code,
    m.content_enc
FROM
    messages AS m
JOIN users AS s ON m.sender_id = s.id
LEFT JOIN rooms AS r ON m.room_id = r.id
WHERE
    m.receiver_id = @receiver_id
ORDER BY
    m.created_at ASC;

-- name: GetMessagesByRoomCode :many
SELECT
    s.nickname AS sender_nickname,
    m.content_enc
FROM
    messages AS m
JOIN users AS s ON m.sender_id = s.id
JOIN rooms AS r ON m.room_id = r.id
WHERE
    r.code = @room_code
ORDER BY
    m.created_at ASC;

-- name: CreateMessage :exec
INSERT INTO messages (
    sender_id,
    receiver_id,
    room_id,
    content_enc
) VALUES (
    @sender_id,
    sqlc.narg(receiver_id),
    sqlc.narg(room_id),
    @content
);

-- name: CountMessagesByRoomCode :one
SELECT
    COUNT(*)
FROM
    messages AS m
JOIN users AS s ON m.sender_id = s.id
JOIN rooms AS r ON m.room_id = r.id
WHERE
    r.code = @room_code;