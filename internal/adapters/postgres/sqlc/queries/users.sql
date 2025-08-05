-- name: GetUserByID :one
SELECT * FROM users WHERE id = @id;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = @username;

-- name: GetUserByNickname :one
SELECT * FROM users WHERE nickname = @nickname;

-- name: GetUserByPrivateCode :one
SELECT * FROM users WHERE private_code = @private_code;

-- name: UserLogin :one
SELECT
    sqlc.embed(u), sqlc.embed(c)
FROM
    users AS u
INNER JOIN clients AS c ON c.user_id = u.id
WHERE
    u.username = @username;

-- name: GetUserFullInfo :one
SELECT
    sqlc.embed(u), sqlc.embed(c)
FROM
    users AS u
INNER JOIN clients AS c ON c.user_id = u.id
WHERE
    u.private_code = @private_code;

-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    nickname,
    private_code,
    role
) VALUES (
    @username,
    @password,
    @nickname,
    @private_code,
    @role
)
RETURNING id;

-- name: ChangeNickname :exec
UPDATE users
SET
    nickname = COALESCE(@nickname, nickname)
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

-- name: IsUserExistByPrivateCode :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE private_code = @private_code
);
