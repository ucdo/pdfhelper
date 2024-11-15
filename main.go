package main

import (
	"Octopus/PdfHelper/middleware"
	"Octopus/PdfHelper/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// CORS 跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 注册路由
	route.Register(r)
	log.Fatal(r.Run(":12345"))
}
