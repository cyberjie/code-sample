// Copyright 2020 CyberJie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package api

import (
	"app/application/api/routers"
	"app/stub"
)

var(
	Routers=[][]stub.Router{routers.UserRouters}
)