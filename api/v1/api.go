package v1

import "github.com/baaj2109/newbee_mall/api/v1/manage"

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
}

var ApiGroupApp = new(ApiGroup)
