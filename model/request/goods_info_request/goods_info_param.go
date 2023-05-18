package goods_info_request

type GoodsInfoAddParam struct {
	GoodsName          string `json:"goodsName"`
	GoodsIntro         string `json:"goodsIntro"`
	GoodsCategoryId    int    `json:"goodsCategoryId"`
	GoodsCoverImg      string `json:"goodsCoverImg"`
	GoodsCarousel      string `json:"goodsCarousel"`
	GoodsDetailContent string `json:"goodsDetailContent"`
	OriginalPrice      string `json:"originalPrice"`
	SellingPrice       string `json:"sellingPrice"`
	StockNum           string `json:"stockNum"`
	Tag                string `json:"tag"`
	GoodsSellStatus    string `json:"goodsSellStatus"`
}
