package main

import (
	"github.com/EikoNakashima/sample-gin/handler"
	"github.com/EikoNakashima/sample-gin/infra"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	g := gin.Default()

	infra.DBInit()

	routes := g.Group("/users")
	{
		routes.POST("/register", handler.Create)
		routes.GET("/userlist", handler.GetAll)
		routes.GET("/user", handler.GetOne)
		routes.PUT("/update", handler.Update)
		routes.DELETE("/delete", handler.Delete)
	}
	g.Run(":3000")
}
