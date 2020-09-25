// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestMTraceRoute(t *testing.T) {
	ctx:=gin.Context{Request: &http.Request{Header:map[string][]string{"user-agent":{"Mobile Phone"}}}}
	ctx.BindUri("/api/test")
	MTraceRoute(&ctx)
	t.Log("test mtraceroute")
}