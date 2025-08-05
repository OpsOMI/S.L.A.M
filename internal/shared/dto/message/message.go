package message

import "github.com/google/uuid"

type MessageReq struct {
	ClientID uuid.UUID
	Content  string
	Room     string
}
