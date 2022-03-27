package main

import (
	"log"

	"github.com/EikoNakashima/sample-gin/handler"
	"github.com/EikoNakashima/sample-gin/infra"
	"github.com/EikoNakashima/sample-gin/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// engineを作成します。
	engine := infra.DBInit()

	// サービスを作成します。
	factory := service.NewService(engine)

	// 最後にengineを閉めます。
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	// Ginを作成します。
	g := gin.Default()

	// サービス層のMiddlewareを作成します。
	g.Use(service.ServiceFactoryMiddleware(factory))

	// v1というグループを作成します。
	routes := g.Group("/v1")

	// ルーティングです。
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}
	// デフォルトは8080です。
	g.Run(":3000")
}
