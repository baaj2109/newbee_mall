package manage

import (
	"errors"
	"strconv"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/request"
	"github.com/baaj2109/newbee_mall/model/request/carousel_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/utils"
	"gorm.io/gorm"
)

type ManageCarouselService struct {
}

func (m *ManageCarouselService) CreateCarousel(request carousel_request.MallCarouselAddParam) (err error) {
	carouseRank, _ := strconv.Atoi(request.CarouselRank)
	carousel := model.Carousel{
		CarouselUrl:  request.CarouselUrl,
		RedirectUrl:  request.RedirectUrl,
		CarouselRank: carouseRank,
		CreateTime:   utils.JSONTime{Time: time.Now()},
		UpdateTime:   utils.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(carousel, utils.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Create(&carousel).Error
	return err
}

func (m *ManageCarouselService) DeleteCarousel(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&model.Carousel{}, "carousel_id in ?", ids.Ids).Error
	return err
}

// update carousel
func (m *ManageCarouselService) UpdateCarousel(request update_request.CarouselUpdateParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("carousel_id = ?", request.CarouselId).First(&model.Carousel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("carousel not found")
	}
	carouseRank, _ := strconv.Atoi(request.CarouselRank)
	mallCarousel := model.Carousel{
		CarouselUrl:  request.CarouselUrl,
		RedirectUrl:  request.RedirectUrl,
		CarouselRank: carouseRank,
		UpdateTime:   utils.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(mallCarousel, utils.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Where("carousel_id = ?", request.CarouselId).UpdateColumns(&mallCarousel).Error
	return err
}

// get carousel
func (m *ManageCarouselService) GetCarousel(id int) (mallCarousel model.Carousel, err error) {
	err = global.GVA_DB.Where("carousel_id = ?", id).First(&mallCarousel).Error
	return mallCarousel, err
}
