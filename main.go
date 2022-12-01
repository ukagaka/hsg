package main

import (
	"hxsg/base"
	"hxsg/internal/middleware"
	"hxsg/internal/router"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	container := base.GetContainer()
	container.InitData()

	r := base.Default()
	r.Use(middleware.LoggerMiddleware(), middleware.RecoveryMiddleware())
	router.InitRouter(r)

	r.Run(":9999")
}
