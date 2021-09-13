package v1

import (
	"freado/api/v1/userApp"
)

type ApiGroup struct {
	AuthApiGroup userApp.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
