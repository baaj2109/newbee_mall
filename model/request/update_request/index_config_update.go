package update_request

type IndexConfigUpdateParams struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigType  int    `json:"configType"`
	GoodsId     int    `json:"goodsId"`
	ConfigRank  string `json:"configRank"`
}
