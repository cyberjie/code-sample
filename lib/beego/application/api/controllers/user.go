// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
beego.Controller
}
func (u *UserController)Echo(){
	u.Ctx.WriteString("echo")
}