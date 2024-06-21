package mytest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/mytest"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pkgTest1"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	MyTestApiGroup   mytest.ApiGroup
	PkgTest1ApiGroup pkgTest1.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
