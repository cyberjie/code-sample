// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Default(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
func ErrorPageNotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "error.html", gin.H{"code": 404})
}
func Echo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": ctx.Query("data")})
}
func FileUpLoad(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{})
}
func UpLoadFile(ctx *gin.Context) {
	upload,err:=ctx.FormFile("file");
	if err==nil{
		ctx.SaveUploadedFile(upload,"upload/"+upload.Filename)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
