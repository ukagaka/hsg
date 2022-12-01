package router

import (
	"hxsg/base"
	"net/http"
)

func Web(r *base.Engine) {

	r.GET("/", func(c *base.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	r.GET("/panic", func(c *base.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

}
