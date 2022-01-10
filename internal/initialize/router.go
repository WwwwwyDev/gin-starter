package initialize

import (
	"gin-starter/internal/middleware"
	"github.com/gin-gonic/gin"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1")
}

//认证
func sysCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1/c")

}

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	g := r.Group("/api")
	sysNoCheckRoleRouter(g)
	sysCheckRoleRouter(g)
	return r
}
