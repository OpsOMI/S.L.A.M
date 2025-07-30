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
