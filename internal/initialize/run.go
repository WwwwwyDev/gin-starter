package initialize

import (
	"fmt"
	"gin-starter/pkg/global"
	"github.com/gin-gonic/gin"
)


func Run(){
	serverCfg := global.CONFIG.Server
	gin.SetMode(serverCfg.RunMode)
	endPoint := fmt.Sprintf(":%d", serverCfg.HttpPort)
	app := RegisterRouter()
	app.Run(endPoint)
}
