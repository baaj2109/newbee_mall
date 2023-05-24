package manage

import (
	"errors"
	"strconv"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request"
	"github.com/baaj2109/newbee_mall/model/request/index_config_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/utils"
	"gorm.io/gorm"
)

type ManageIndexConfigService struct {
}

// create index config by index config request
func (m *ManageIndexConfigService) CreateIndexConfig(request index_config_request.IndexConfigAddParams) error {

	var goodsInfo model.GoodsInfo
	if errors.Is(global.GVA_DB.Where("goods_id=?", request.GoodsId).First(&goodsInfo).Error, gorm.ErrRecordNotFound) {
		return errors.New("goods not found")
	}
	if errors.Is(global.GVA_DB.Where("config_type =? and goods_id=? and is_deleted=0", request.ConfigType, request.GoodsId).First(&model.IndexConfig{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("already have goods id with same config type")
	}
	goodsId, _ := strconv.Atoi(request.GoodsId)
	configRank, _ := strconv.Atoi(request.ConfigRank)
	mallIndexConfig := model.IndexConfig{
		ConfigName:  request.ConfigName,
		ConfigType:  request.ConfigType,
		GoodsId:     goodsId,
		RedirectUrl: request.RedirectUrl,
		ConfigRank:  configRank,
		CreateTime:  utils.JSONTime{Time: time.Now()},
		UpdateTime:  utils.JSONTime{Time: time.Now()},
	}
	if err := utils.Verify(mallIndexConfig, utils.IndexConfigAddParamVerify); err != nil {
		return errors.New(err.Error())
	}

	err := global.GVA_DB.Create(&mallIndexConfig).Error
	return err

}

// / delet index config by id request
func (m *ManageIndexConfigService) DeleteIndexConfig(idReq request.IdsReq) error {
	return global.GVA_DB.Where("config_id in ?", idReq.Ids).Delete(&model.IndexConfig{}).Error
}

// / update index config by update index config params request
func (m *ManageIndexConfigService) UpdateIndexConfig(request update_request.IndexConfigUpdateParams) (err error) {
	if errors.Is(global.GVA_DB.Where("goods_id = ?", request.GoodsId).First(&model.GoodsInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("goods not found")
	}
	if errors.Is(global.GVA_DB.Where("config_id=?", request.ConfigId).First(&model.IndexConfig{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("config not found")
	}
	configRank, _ := strconv.Atoi(request.ConfigRank)
	mallIndexConfig := model.IndexConfig{
		ConfigId:    request.ConfigId,
		ConfigType:  request.ConfigType,
		ConfigName:  request.ConfigName,
		RedirectUrl: request.RedirectUrl,
		GoodsId:     request.GoodsId,
		ConfigRank:  configRank,
		UpdateTime:  utils.JSONTime{Time: time.Now()},
	}
	if err = utils.Verify(mallIndexConfig, utils.IndexConfigUpdateParamVerify); err != nil {
		return errors.New(err.Error())
	}
	var newIndexConfig model.IndexConfig
	err = global.GVA_DB.Where("config_type=? and goods_id=?", mallIndexConfig.ConfigType, mallIndexConfig.GoodsId).First(&newIndexConfig).Error
	if err != nil && newIndexConfig.ConfigId == mallIndexConfig.ConfigId {
		return errors.New("already have goods id with same config type")
	}
	err = global.GVA_DB.Where("config_id=?", mallIndexConfig.ConfigId).Updates(&mallIndexConfig).Error
	return err
}

func (m *ManageIndexConfigService) GetMallIndexConfig(id uint) (mallIndexConfig model.IndexConfig, err error) {
	err = global.GVA_DB.Where("config_id = ?", id).First(&mallIndexConfig).Error
	return
}
