package logic

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/szpnygo/goclip/svc"
)

type LogicContext struct {
	*gin.Context
	*svc.ServiceContext
}

func NewLogicContext(gc *gin.Context, sc *svc.ServiceContext) *LogicContext {
	return &LogicContext{
		Context:        gc,
		ServiceContext: sc,
	}
}

func (l *LogicContext) Result(code int, data any, msg ...string) {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	l.Set("code", strconv.Itoa(code))
	l.JSON(200, gin.H{
		"code":    code,
		"data":    data,
		"message": m,
	})
}
