package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/response"
)

type MallCarouselService struct {
}

// GetCarouselsForIndex 回傳固定數量的輪播圖對象（首頁使用）
func (m *MallCarouselService) GetCarouselsForIndex(num int) (mallCarousels []model.Carousel, list interface{}, err error) {
	var carouselIndexs []response.MallCarouselIndexResponse
	err = global.GVA_DB.Where("is_deleted = 0").Order("carousel_rank desc").Limit(num).Find(&mallCarousels).Error
	for _, carousel := range mallCarousels {
		carouselIndex := response.MallCarouselIndexResponse{
			CarouselUrl: carousel.CarouselUrl,
			RedirectUrl: carousel.RedirectUrl,
		}
		carouselIndexs = append(carouselIndexs, carouselIndex)
	}
	return mallCarousels, carouselIndexs, err
}
