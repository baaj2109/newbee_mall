package update_request

type UpdateAdminNameParam struct {
	LoginUserName string `json:"loginUserName"`
	NickName      string `json:"nickName"`
}

type UpdateAdminPasswordParam struct {
	OriginalPassword string `json:"originalPassword"`
	NewPassword      string `json:"newPassword"`
}
