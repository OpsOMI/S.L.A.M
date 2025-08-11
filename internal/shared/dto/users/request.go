package users

type LoginReq struct {
	ClientKey string `json:"clientKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RegisterReq struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
