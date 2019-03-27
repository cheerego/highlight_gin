package app

import (
	"git.catchme.cn/placeless/highlight_gin/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	CodeStatus int         `json:"code_status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		CodeStatus: httpCode,
		Message:    e.GetMsg(errCode),
		Data:       data,
	})
	return
}
