// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	/**beego.SetStaticPath("/static",stub.MallStaticPath)
	beego.SetViewsPath(stub.MallViewsPath)
    app:=beego.Router("/",&application.Application{});
    application.RegistRouters(app)
	beego.Run(":80")*/
	db,err:=gorm.Open("mysql","root:mysql@/test?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(db.Exec("drop table doc;").Error)
}