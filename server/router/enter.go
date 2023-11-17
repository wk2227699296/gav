package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/wkTest"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	WkTest  wkTest.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
