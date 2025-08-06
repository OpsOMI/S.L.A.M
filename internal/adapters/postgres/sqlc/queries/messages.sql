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
    room_id,
    content_enc
) VALUES (
    @sender_id,
    @room_id,
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