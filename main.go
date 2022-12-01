package main

import (
	"hxsg/base"
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
	router.Web(r)
	r.Run(":9999")
}
