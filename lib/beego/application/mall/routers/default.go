package routers

import (
	"app/application/mall/controllers"
	"app/stub"
)

var(
	MallRouters=[]stub.Router{{Path: "/mall",Controller: &controllers.DefaultController{},Caller: "*:Index"}}
)