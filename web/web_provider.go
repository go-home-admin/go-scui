package web

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"io/fs"
	"net/http"
)

//go:embed dist
var static embed.FS

// Web @Bean
type Web struct {
	http  *servers.Http            `inject:""`
	route *providers.RouteProvider `inject:""`
}

func (s *Web) Boot() {
	fSys, err := fs.Sub(static, "dist")
	if err != nil {
		return
	}
	s.http.StaticFS("/web", http.FS(fSys))
	s.http.GET("/", func(context *gin.Context) {
		context.Redirect(302, "/web")
	})
}
