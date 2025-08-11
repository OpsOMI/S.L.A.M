package client

type ClientReq struct {
	ClientKey string `json:"clientKey"`
}

type ClientResp struct {
	IsExists bool `json:"isExists"`
}

func ToClientResp(
	isExists bool,
) ClientResp {
	return ClientResp{
		IsExists: isExists,
	}
}
