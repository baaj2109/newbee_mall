package service

import (
	"github.com/baaj2109/newbee_mall/service/mall"
	"github.com/baaj2109/newbee_mall/service/manage"
)

type ServiceGroup struct {
	// ExampleServiceGroup example.ServiceGroup
	ManageServiceGroup manage.ManageServiceGroup
	MallServiceGroup   mall.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
