// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func MTraceRoute(ctx *gin.Context) {
	fmt.Printf("url:%s,user-agent:%s\n", ctx.FullPath(), ctx.GetHeader("user-agent"))
	ctx.Next()
}
