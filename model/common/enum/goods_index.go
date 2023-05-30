package enum

type IndexConfigEnum int8

// 首頁配置項 1-搜索框熱搜 2-搜索下拉框熱搜 3-(首頁)熱銷商品 4-(首頁)新品上線 5-(首頁)為你推薦
const (
	IndexSearchHots     IndexConfigEnum = 1
	IndexSearchDownHots IndexConfigEnum = 2
	IndexGoodsHot       IndexConfigEnum = 3
	IndexGoodsNew       IndexConfigEnum = 4
	IndexGoodsRecommond IndexConfigEnum = 5
)

func (i IndexConfigEnum) Info() (int, string) {
	switch i {
	case IndexSearchHots:
		return 1, "INDEX_SEARCH_HOTS"
	case IndexSearchDownHots:
		return 2, "二級分類"
	case IndexGoodsHot:
		return 3, "三級分類"
	case IndexGoodsNew:
		return 4, "三級分類"
	case IndexGoodsRecommond:
		return 5, "三級分類"
	default:
		return 0, "DEFAULT"
	}
}

func (i IndexConfigEnum) Code() int {
	switch i {
	case IndexSearchHots:
		return 1
	case IndexSearchDownHots:
		return 2
	case IndexGoodsHot:
		return 3
	case IndexGoodsNew:
		return 4
	case IndexGoodsRecommond:
		return 5
	default:
		return 0
	}
}
