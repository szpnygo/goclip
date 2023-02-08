package routes

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/szpnygo/goclip/controller"
	"github.com/szpnygo/goclip/logic"
	"github.com/szpnygo/goclip/svc"
)

type RoutesManager struct {
	svcCtx *svc.ServiceContext
	*gin.Engine
}

func NewRoutesManager(svcCtx *svc.ServiceContext, r *gin.Engine) *RoutesManager {
	return &RoutesManager{
		svcCtx: svcCtx,
		Engine: r,
	}
}

func (r *RoutesManager) InitRoutes() {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Push-Id", "App", "App-Version", "X-Device-Id", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(origin, "localhost")
		},
	}))

	r.POST("/image", r.ctx(controller.UploadImage))
	r.POST("/images", r.ctx(controller.ListImage))
	r.GET("/search", r.ctx(controller.Search))
}

func (r *RoutesManager) ctx(f func(ctx *logic.LogicContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f(logic.NewLogicContext(ctx, r.svcCtx))
	}
}
