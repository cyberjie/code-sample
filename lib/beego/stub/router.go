// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package stub

import "github.com/astaxie/beego"

type Router struct {
	Path string
	Controller beego.ControllerInterface
	Caller string
}
