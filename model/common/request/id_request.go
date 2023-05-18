package request

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
