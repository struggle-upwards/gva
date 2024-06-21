package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/mytest"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest1"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Mytest   mytest.RouterGroup
	PkgTest1 pkgTest1.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
