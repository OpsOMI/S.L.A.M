package rooms

type JoinReq struct {
	RoomCode string `json:"roomCode"`
	Password string `json:"password"`
}

type CreateReq struct {
	Password string `json:"password"`
}

type MyRoomReq struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}
