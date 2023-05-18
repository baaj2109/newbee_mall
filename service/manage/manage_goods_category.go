package manage

import (
	"errors"
	"strconv"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/request"
	"github.com/baaj2109/newbee_mall/model/request/goods_category_request"
	"github.com/baaj2109/newbee_mall/utils"
	"gorm.io/gorm"
)

type ManageGoodsCategoryService struct {
}

// add category by goods category request
func (m *ManageGoodsCategoryService) CreateCategory(request goods_category_request.GoodsCategoryReq) (err error) {
	if !errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		request.CategoryLevel, request.CategoryName).First(&model.GoodsCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("category already exists")
	}

	rank, _ := strconv.Atoi(request.CategoryRank)
	category := model.GoodsCategory{
		CategoryLevel: request.CategoryLevel,
		CategoryName:  request.CategoryName,
		CategoryRank:  rank,
		IsDeleted:     0,
		CreateTime:    utils.JSONTime{Time: time.Now()},
		UpdateTime:    utils.JSONTime{Time: time.Now()},
	}
	/// vertify should be in api layer, but front end pass string to here
	/// and my verify policy is integer
	if err = utils.Verify(category, utils.GoodsCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Create(&category).Error
}

// / update category by goods category request
func (m *ManageGoodsCategoryService) UpdateCategory(request goods_category_request.GoodsCategoryReq) (err error) {
	if !errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		request.CategoryLevel, request.CategoryName).First(&model.GoodsCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("category already exists")
	}
	rank, _ := strconv.Atoi(request.CategoryRank)
	category := model.GoodsCategory{
		CategoryName: request.CategoryName,
		CategoryRank: rank,
		UpdateTime:   utils.JSONTime{Time: time.Now()},
	}
	if err := utils.Verify(category, utils.GoodsCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Where("category_id =?", request.CategoryId).Updates(&category).Error
}

func (m *ManageGoodsCategoryService) SelectCategoryById(categoryId int) (goodsCategory model.GoodsCategory, err error) {
	err = global.GVA_DB.Where("category_id=?", categoryId).First(&goodsCategory).Error
	return goodsCategory, err
}

func (m *ManageGoodsCategoryService) DeleteGoodsCategoriesByIds(ids request.IdsReq) (goodsCategory model.GoodsCategory, err error) {
	err = global.GVA_DB.Where("category_id in ?", ids.Ids).UpdateColumns(model.GoodsCategory{IsDeleted: 1}).Error
	return goodsCategory, err
}

func (m *ManageGoodsCategoryService) SelectByLevelAndParentIdsAndNumber(parentId int, categoryLevel int) (goodsCategory model.GoodsCategory, err error) {
	err = global.GVA_DB.Where("category_id in ?", parentId).Where("category_level=?", categoryLevel).Where("is_deleted=0").Error
	return goodsCategory, err

}
