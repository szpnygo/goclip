package main

import (
	"github.com/gin-gonic/gin"
	"github.com/szpnygo/goclip/routes"
	"github.com/szpnygo/goclip/svc"
)

func main() {
	engine := gin.Default()
	svcCtx := svc.NewServiceContext()
	route := routes.NewRoutesManager(svcCtx, engine)
	route.InitRoutes()
	_ = engine.Run()
}
