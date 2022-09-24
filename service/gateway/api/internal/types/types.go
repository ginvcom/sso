package types

type LoginReq struct {
	SystemCode string `json:"systemCode"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	Remember   string `json:"remember"`
}

type LoginReply struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Gender      int64  `json:"gender"`
	AccessToken string `json:"accessToken"`
	Redirect    string `json:"redirect"`
	Expire      int64  `json:"expire"`
}

type LoginResponse struct {
	Code int64      `json:"code"`
	Data LoginReply `json:"data"`
	Msg  string     `json:"msg"`
}

type UserBasic struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Gender int64  `json:"gender"`
}

type LoginInfo struct {
	SystemCode  string
	Mobile      string
	MobileMsg   string
	Password    string
	PasswordMsg string
	Msg         string
	Code        int64
	Redirect    string
}

type VerifyReq struct {
	Token       string `json:"token"`
	SystemCode  string `json:"systemCode"`
	ServiceCode string `json:"serviceCode"`
	URI         string `json:"uri"`
}

type VerifyReply struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type VerifyResponse struct {
	Code int
	Data VerifyReply
	Msg  string
}

type MenuReq struct {
	Token      string `json:"token"`
	SystemCode string `json:"systemCode"`
}
