package manage

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request"
	"github.com/baaj2109/newbee_mall/model/request/index_config_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageIndexConfigApi struct {
}

// create index config by gin context
func (m *ManageIndexConfigApi) CreateIndexConfig(c *gin.Context) {
	var param index_config_request.IndexConfigAddParams
	_ = c.ShouldBindJSON(&param)
	if err := mallIndexConfigService.CreateIndexConfig(param); err != nil {
		global.GVA_LOG.Error("failed to create index config", zap.Error(err))
		response.FailWithMessage("failed to create index config"+err.Error(), c)
	}
	response.OkWithMessage("success to create index config", c)
}

// delete index config by gin context
func (m *ManageIndexConfigApi) DeleteIndexConfig(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallIndexConfigService.DeleteIndexConfig(ids); err != nil {
		global.GVA_LOG.Error("failed to delete index config", zap.Error(err))
		response.FailWithMessage("failed to delete index config"+err.Error(), c)
	}
	response.OkWithMessage("success to delete index config", c)
}

// update index config by gin context
func (m *ManageIndexConfigApi) UpdateIndexConfig(c *gin.Context) {
	var request update_request.IndexConfigUpdateParams
	_ = c.ShouldBindJSON(&request)
	if err := mallIndexConfigService.UpdateIndexConfig(request); err != nil {
		global.GVA_LOG.Error("failed to update index config", zap.Error(err))
		response.FailWithMessage("failed to update index config"+err.Error(), c)
	}
	response.OkWithMessage("success to update index config", c)
}

// get index config by gin context
func (m *ManageIndexConfigApi) FindIndexConfig(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	indexConfig, err := mallIndexConfigService.GetMallIndexConfig(uint(id))
	if err != nil {
		global.GVA_LOG.Error("failed to get index config", zap.Error(err))
		response.FailWithMessage("failed to get index config"+err.Error(), c)
	}
	response.OkWithData(indexConfig, c)
}
