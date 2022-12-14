// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	api_demo "github.com/go-home-admin/go-admin/app/http/api/api_demo"
	public "github.com/go-home-admin/go-admin/app/http/api/public"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type ApiRoutes struct {
	api_demo *api_demo.Controller `inject:""`
	public   *public.Controller   `inject:""`
}

func (c *ApiRoutes) GetGroup() string {
	return "api"
}
func (c *ApiRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Get("/api/demo"): c.api_demo.GinHandleHome,
		api.Get("/version"):  c.public.GinHandleHome,
	}
}
