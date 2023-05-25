package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
)

type MallIndexInfoService struct {
}

/// GetConfigGoodsForIndex return相關IndexConfig

func (m MallIndexInfoService) GetConfigGoodsForIndex(configType int, num int) (list interface{}, err error) {
	var indexConfigs []model.IndexConfig
	err = global.GVA_DB.Where("config_type = ?", configType).Where("is_deleted = 0").Order("config_rank desc").Limit(num).Find(&indexConfigs).Error
	if err != nil {
		return
	}
	/// get goods id
	var ids []int
	for _, indexConfig := range indexConfigs {
		ids = append(ids, indexConfig.GoodsId)
	}
	/// get goods info
	var goodsList []model.GoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", ids).Find(&goodsList).Error
	var indexGoodsList []response.MallIndexConfigGoodsResponse

	for _, indexGoods := range goodsList {
		res := response.MallIndexConfigGoodsResponse{
			GoodsId:       indexGoods.GoodsId,
			GoodsName:     utils.SubStrLen(indexGoods.GoodsName, 30),
			GoodsIntro:    utils.SubStrLen(indexGoods.GoodsIntro, 30),
			GoodsCoverImg: indexGoods.GoodsCoverImg,
			SellingPrice:  indexGoods.SellingPrice,
			Tag:           indexGoods.Tag,
		}
		indexGoodsList = append(indexGoodsList, res)
	}
	return indexGoodsList, err
}
