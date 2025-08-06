package rooms

type JoinReq struct {
	RoomCode string `json:"roomCode"`
	Password string `json:"password"`
}
