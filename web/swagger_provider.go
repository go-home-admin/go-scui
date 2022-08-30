package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"strings"
)

//go:embed swagger
var web embed.FS

// Swaager @Bean
type Swaager struct {
	http  *servers.Http            `inject:""`
	route *providers.RouteProvider `inject:""`
}

func (s *Swaager) Boot() {
	s.http.GET("/swagger", func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		by, err := web.ReadFile("web/index.html")
		if err != nil {
			context.String(200, "未生成文档, 请执行 toolset make:swagger")
		} else {
			context.String(200, string(by))
		}
	})

	s.http.GET("/swagger/doc.json", func(context *gin.Context) {
		by, err := web.ReadFile("web/doc.json")
		if err != nil {
			context.String(200, "未生成文档, 请执行 toolset make:swagger")
		} else {
			str := string(by)
			for group, v := range s.route.RouteGroupConfig {
				str = strings.Replace(str, "$["+group+"]", v.GetPrefix(), -1)
			}
			context.String(200, str)
		}
	})
}
