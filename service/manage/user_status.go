package manage

import (
	"errors"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request"
)

type UserStatusService struct {
}

func (u *UserStatusService) LockUser(idReq request.IdsReq, lockStatus int) error {
	if lockStatus != 0 && lockStatus != 1 {
		return errors.New("invalid lock status")
	}
	return global.GVA_DB.Model(&model.User{}).Where("user_id in ?", idReq.Ids).Update("locked_flag", lockStatus).Error
}
