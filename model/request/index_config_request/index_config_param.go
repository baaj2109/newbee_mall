package index_config_request

type IndexConfigAddParams struct {
	ConfigName  string `json:"configName"`
	ConfigType  int    `json:"configType"`
	GoodsId     string `json:"goodsId"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigRank  string `json:"configRank"`
}
