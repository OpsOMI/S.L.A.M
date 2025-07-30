-- name: AddUserToRoom :exec
INSERT INTO room_members (
    room_id,
    user_id
) VALUES (
    @room_id,
    @user_id
);
-- name: RemoveUserFromRoom :exec
DELETE FROM room_members
WHERE room_id = @room_id AND user_id = @user_id;

-- name: IsUserInRoom :one
SELECT EXISTS (
    SELECT 1 FROM room_members
    WHERE room_id = @room_id AND user_id = @user_id
);
