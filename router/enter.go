package router

import "github.com/baaj2109/newbee_mall/router/manage"

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	// Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
