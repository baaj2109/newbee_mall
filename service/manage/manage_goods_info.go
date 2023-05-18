package manage

import (
	"errors"
	"strconv"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/enum"
	"github.com/baaj2109/newbee_mall/model/request/goods_info_request"
	"github.com/baaj2109/newbee_mall/utils"
	"gorm.io/gorm"
)

type ManageGoodsInfoService struct {
}

// create goods info
func (m *ManageGoodsInfoService) CreateGoodsInfo(request goods_info_request.GoodsInfoAddParam) (err error) {
	var goodsCategory model.GoodsCategory
	err = global.GVA_DB.Where("category_id = ? AND is_deleted=0", request.GoodsCategoryId).First(&goodsCategory).Error
	if err != nil {
		return errors.New("goods category not found")
	}
	if goodsCategory.CategoryLevel != enum.LevelThree.Code() {
		return errors.New("invalid goods category level")
	}

	if errors.Is(global.GVA_DB.Where("goods_name = ? AND goods_category_id = ?", request.GoodsName, request.GoodsCategoryId).First((&model.GoodsInfo{})).Error, gorm.ErrRecordNotFound) {
		return errors.New("goods name already exists")
	}
	originalPrice, _ := strconv.Atoi(request.OriginalPrice)
	sellingPrice, _ := strconv.Atoi(request.SellingPrice)
	stockNum, _ := strconv.Atoi(request.StockNum)
	goodsSellStatus, _ := strconv.Atoi(request.GoodsSellStatus)

	goodsInfo := model.GoodsInfo{
		GoodsName:          request.GoodsName,
		GoodsIntro:         request.GoodsIntro,
		GoodsCategoryId:    request.GoodsCategoryId,
		GoodsCoverImg:      request.GoodsCoverImg,
		GoodsDetailContent: request.GoodsDetailContent,
		OriginalPrice:      originalPrice,
		SellingPrice:       sellingPrice,
		StockNum:           stockNum,
		Tag:                request.Tag,
		GoodsSellStatus:    goodsSellStatus,
		CreateTime:         utils.JSONTime{Time: time.Now()},
		UpdateTime:         utils.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(goodsInfo, utils.GoodsAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Create(&goodsInfo).Error
	return err
}
