package controller

import (
	"hxsg/base"
	"net/http"
)

func Login(c *base.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *base.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")

	c.String(http.StatusOK, "ok login : %v, pass:%v", userName, password)
}

func Regist(c *base.Context) {
	c.HTML(http.StatusOK, "regist.html", nil)
}

func RegistPost(c *base.Context) {
	data := map[string]interface{}{
		"message": "哈哈哈，你错了",
	}

	c.HTML(http.StatusOK, "regist.html", data)
}
