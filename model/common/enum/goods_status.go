package enum

type GoodsStatusEnum int

const (
	GOODS_DEFAULT GoodsStatusEnum = -9
	GOODS_UNDER   GoodsStatusEnum = 0
)

func (g GoodsStatusEnum) Code() int {
	switch g {
	case GOODS_UNDER:
		return 0
	default:
		return -9
	}
}
