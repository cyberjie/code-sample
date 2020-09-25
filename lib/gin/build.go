// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"app/internal/middleware"
	"app/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
    middleware.UseAllMiddleWares(engine)
	router.RouteAllRouters(engine)
	engine.Run(":80")
}
