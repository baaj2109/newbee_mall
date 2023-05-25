package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/enum"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/jinzhu/copier"
)

type MallGoodsCategoryService struct {
}

func (m *MallGoodsCategoryService) GetCategoriesForIndex() (newBeeMallIndexCategoryVOS []response.NewBeeMallIndexCategoryVO, err error) {
	// get first category data
	firstLevelCategories, _ := selectByLevelAndParentIdsAndNumber([]int{0}, enum.LevelOne.Code(), 10)
	if firstLevelCategories == nil {
		return
	}
	var firstLevelCategoryIds []int
	for _, firstLevelCategory := range firstLevelCategories {
		firstLevelCategoryIds = append(firstLevelCategoryIds, firstLevelCategory.CategoryId)
	}
	/// get second category data
	secondLevelCategories, _ := selectByLevelAndParentIdsAndNumber(firstLevelCategoryIds, enum.LevelTwo.Code(), 0)
	if secondLevelCategories == nil {
		return
	}
	var secondLevelCategoryIds []int
	for _, secondLevelCategory := range secondLevelCategories {
		secondLevelCategoryIds = append(secondLevelCategoryIds, secondLevelCategory.CategoryId)
	}
	/// get third category data
	thirdLevelCategories, _ := selectByLevelAndParentIdsAndNumber(secondLevelCategoryIds, enum.LevelThree.Code(), 0)
	if thirdLevelCategories == nil {
		return
	}
	/// cluster thirdLevelCategories by parent id
	thirdLevelCategoryMap := make(map[int][]model.GoodsCategory)
	for _, thirdLevelCategory := range thirdLevelCategories {
		thirdLevelCategoryMap[thirdLevelCategory.ParentId] = []model.GoodsCategory{}
	}
	for k, v := range thirdLevelCategoryMap {
		for _, third := range thirdLevelCategories {
			if k == third.ParentId {
				v = append(v, third)
			}
			thirdLevelCategoryMap[k] = v
		}
	}

	var secondLevelCategoryVOS []response.SecondLevelCategoryVO
	/// handle second category
	for _, secondLevelCategory := range secondLevelCategories {
		var secondLevelCategoryVO response.SecondLevelCategoryVO
		err = copier.Copy(&secondLevelCategoryVO, &secondLevelCategory)
		// if there is data under second level category then put into secondLevelCategoryVOS
		if _, ok := thirdLevelCategoryMap[secondLevelCategory.CategoryId]; ok {
			/// get list from third level category map by second level category id
			tempGoodsCategories := thirdLevelCategoryMap[secondLevelCategory.CategoryId]
			var thirdLevelCategoryRes []response.ThirdLevelCategoryVO
			err = copier.Copy(&thirdLevelCategoryRes, &tempGoodsCategories)
			secondLevelCategoryVO.ThirdLevelCategoryVOS = thirdLevelCategoryRes
			secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
		}
	}

	if secondLevelCategoryVOS == nil {
		return
	}
	/// handle first category
	secondLevelCategoryVOMap := make(map[int][]response.SecondLevelCategoryVO)
	for _, secondLevelCategory := range secondLevelCategoryVOS {
		secondLevelCategoryVOMap[secondLevelCategory.ParentId] = []response.SecondLevelCategoryVO{}
	}
	for k, v := range secondLevelCategoryVOMap {
		for _, second := range secondLevelCategoryVOS {
			if k == second.ParentId {
				var secondLevelCategory response.SecondLevelCategoryVO
				copier.Copy(&secondLevelCategory, &second)
				v = append(v, secondLevelCategory)
			}
			secondLevelCategoryVOMap[k] = v
		}
	}

	for _, firstCategory := range firstLevelCategories {
		var newBeeMallIndexCategoryVO response.NewBeeMallIndexCategoryVO
		err = copier.Copy(&newBeeMallIndexCategoryVO, &firstCategory)
		/// if there is data under first category, then put into newBeeMallIndexCategoryVOS
		if _, ok := secondLevelCategoryVOMap[firstCategory.CategoryId]; ok {
			/// get list from second level category map by first level category id
			tempGoodsCategories := secondLevelCategoryVOMap[firstCategory.CategoryId]
			newBeeMallIndexCategoryVO.SecondLevelCategoryVOS = tempGoodsCategories
			newBeeMallIndexCategoryVOS = append(newBeeMallIndexCategoryVOS, newBeeMallIndexCategoryVO)
		}
	}

	return
}

// / get category data
func selectByLevelAndParentIdsAndNumber(ids []int, level int, limit int) (categories []model.GoodsCategory, err error) {

	global.GVA_DB.Where("parent_id in ? and category_level =? and is_deleted = 0", ids, level).
		Order("category_rank desc").Limit(limit).Find(&categories)
	return
}
