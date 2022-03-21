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
	engine := infra.DBInit()
	factory := service.NewService(engine)
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()
	g := gin.Default()
	g.Use(service.ServiceFactoryMiddleware(factory))

	routes := g.Group("/v1")
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}
	g.Run(":3000")
}
