// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package router

import "github.com/gin-gonic/gin"

func RouteAllRouters(engine *gin.Engine){
	engine.Static("/static","internal/static")
	engine.LoadHTMLGlob("internal/template/*")
	engine.GET("/index",Default)
	engine.GET("/",Default)
	engine.GET("/404.html",ErrorPageNotFound);
	engine.GET("/upload",FileUpLoad)
	api:=engine.Group("/api");
	{
		api.Any("/echo", Echo)
		api.POST("/upload",UpLoadFile)
	}
}