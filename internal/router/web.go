package router

import (
	"hxsg/base"
	"hxsg/internal/controller"
	"net/http"
)

func InitRouter(r *base.Engine) {

	r.GET("/login", controller.Login)
	r.POST("/login", controller.LoginPost)

	r.GET("/regist", controller.Regist)

	r.POST("/regist", controller.RegistPost)

	r.GET("/", func(c *base.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})

}
