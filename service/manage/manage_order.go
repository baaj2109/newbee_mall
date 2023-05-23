package manage

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/enum"
	"github.com/baaj2109/newbee_mall/model/common/request"
	"github.com/baaj2109/newbee_mall/utils"
)

type ManageOrderService struct {
}

// check order done
func (m *ManageOrderService) CheckDone(ids request.IdsReq) (err error) {
	var orders []model.Order
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(
				model.Order{
					OrderStatus: 2,
					UpdateTime:  utils.JSONTime{Time: time.Now()},
				}).Error; err != nil {
				return err
			}
		} else {
			err = errors.New("order status is not paid, so can't check done")
		}
	}
	return
}

// close order by id request
func (m *ManageOrderService) CloseOrder(ids request.IdsReq) (err error) {
	var orders []model.Order
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus == enum.ORDER_SUCCESS.Code() || order.OrderStatus < 0 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(
				model.Order{
					OrderStatus: enum.ORDER_CLOSED_BY_JUDGE.Code(),
					UpdateTime:  utils.JSONTime{Time: time.Now()},
				}).Error; err != nil {
				return err
			}
		} else {
			err = errors.New("order status is not success, so can't close order")
		}
	}
	return
}

func (m *ManageOrderService) CheckOut(ids request.IdsReq) (err error) {
	var orders []model.Order
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() && order.OrderStatus != enum.ORDER_PACKAGED.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(
				model.Order{
					OrderStatus: 3,
					UpdateTime:  utils.JSONTime{Time: time.Now()},
				}).Error; err != nil {
				return err
			}
		} else {
			err = errors.New("order status is not paid, so can't check out")
		}
	}
	return
}
