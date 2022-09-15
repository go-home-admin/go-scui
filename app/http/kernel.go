package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// Kernel @Bean
type Kernel struct {
	*servers.Http `inject:""`
}

func (k *Kernel) Init() {
	// 全局中间件
	k.Middleware = []gin.HandlerFunc{
		gin.Recovery(),
		Cors(),
	}

	// 分组中间件, 在路由提供者中自行设置
	k.MiddlewareGroup = map[string][]gin.HandlerFunc{
		"admin": {},
		"api":   {},
	}

	k.SetTrustedProxies(nil)
}

func (k Kernel) Boot() {
	fmt.Printf("Please check http://127.0.0.1:%v\n", k.Port)
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return NewKernel()
}
