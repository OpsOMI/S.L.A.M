-- name: GetUserByID :one
SELECT * FROM users WHERE id = @id;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = @username;

-- name: GetUserByNickname :one
SELECT * FROM users WHERE nickname = @nickname;

-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    nickname,
    role
) VALUES (
    @username,
    @password,
    @nickname,
    @role
)
RETURNING id;

-- name: ChangeNickname :exec
UPDATE users
SET
    nickname = COALESCE(@nickname, nickname)
WHERE
    id = @id;

-- name: BanUser :exec
UPDATE users
SET
    role = 'banned'
WHERE
    id = @id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = @id;

-- name: IsUserExistByID :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE id = @id
);

-- name: IsUserExistByUsername :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE username = @username
);

-- name: IsUserExistByNickname :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE nickname = @nickname
);
