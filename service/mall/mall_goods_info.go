package mall

import (
	"errors"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/jinzhu/copier"
)

type MallGoodsInfoService struct {
}

func (m *MallGoodsInfoService) MallGoodsListBySearch(pageNumber int, goodsCategoryId int, keyword string, orderBy string) (searchGoodsList []response.GoodsSearchResponse, total int64, err error) {
	/// search accounding query
	var goodsList []model.GoodsInfo
	db := global.GVA_DB.Model(&model.GoodsInfo{})
	if keyword != "" {
		db.Where("goods_name like ? or goods_intro like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if goodsCategoryId >= 0 {
		db.Where("goods_category_id= ?", goodsCategoryId)
	}
	err = db.Count(&total).Error
	switch orderBy {
	case "new":
		db.Order("goods_id desc")
	case "price":
		db.Order("selling_price asc")
	default:
		db.Order("stock_num desc")
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	err = db.Limit(limit).Offset(offset).Find(&goodsList).Error
	// return
	for _, goods := range goodsList {
		searchGoods := response.GoodsSearchResponse{
			GoodsId:       goods.GoodsId,
			GoodsName:     utils.SubStrLen(goods.GoodsName, 28),
			GoodsIntro:    utils.SubStrLen(goods.GoodsIntro, 28),
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
		}
		searchGoodsList = append(searchGoodsList, searchGoods)
	}
	return
}

func (m *MallGoodsInfoService) GetMallGoodsInfo(id int) (res response.GoodsInfoDetailResponse, err error) {
	var mallGoodsInfo model.GoodsInfo
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	if mallGoodsInfo.GoodsSellStatus != 0 {
		return response.GoodsInfoDetailResponse{}, errors.New("goods is sold out")
	}
	err = copier.Copy(&res, &mallGoodsInfo)
	if err != nil {
		return response.GoodsInfoDetailResponse{}, err
	}
	var list []string
	list = append(list, mallGoodsInfo.GoodsCarousel)
	res.GoodsCarouselList = list

	return
}
