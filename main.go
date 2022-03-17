package main

import (
	"github.com/EikoNakashima/sample-gin/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()

	// CRUD 書籍
	group := engine.Group("/users")
	{
		group.POST("/register", handler.Create)
		group.GET("/userlist", handler.GetAll)
		group.GET("/user", handler.GetOne)
		group.PUT("/update", handler.Update)
		group.DELETE("/delete", handler.Delete)
	}
	engine.Run(":3000")
}
