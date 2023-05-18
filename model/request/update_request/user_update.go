package update_request

type UpdateNameParam struct {
	LoginUserName string `json:"loginUserName"`
	NickName      string `json:"nickName"`
}

type UpdatePasswordParam struct {
	OriginalPassword string `json:"originalPassword"`
	NewPassword      string `json:"newPassword"`
}
