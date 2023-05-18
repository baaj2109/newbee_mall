package user_request

type AdminLoginParam struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"passwordMd5"`
}
type AdminParam struct {
	LoginUserName string `json:"loginUserName"`
	LoginPassword string `json:"loginPassword"`
	NickName      string `json:"nickName"`
}
