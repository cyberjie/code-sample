// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package application

import (
	"app/application/api"
	"app/application/mall"
	"app/stub"
	"github.com/astaxie/beego"
)

type Application struct {
	beego.Controller
}
func registRoute(app *beego.App,route [][]stub.Router){
	for _,r1:=range route{
		for _,r2:=range r1{
			app.Handlers.Add(r2.Path,r2.Controller,r2.Caller)
		}
	}
}
func RegistRouters(app *beego.App){
	routers:=[][][]stub.Router{api.Routers,mall.Routers}
	for _,v:=range routers{
		registRoute(app,v)
	}
}