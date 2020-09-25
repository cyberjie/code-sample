package routers

import (
	"app/application/api/controllers"
	"app/stub"
)

var(
	UserRouters=[]stub.Router{{Path: "/api/echo",Controller: &controllers.UserController{},Caller: "*:Echo"}}
)